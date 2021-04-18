package parser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/echo/library"
	"github.com/guoruibiao/echo/models/dao"
	"net"
	"strings"
)

// 实现上游请求参数的解析，并配合配置信息进行最终代理的请求参数封装
func GetRequest(ctx *gin.Context, cfg *dao.Configuration) (request dao.Request) {
	// 搞一些默认参数
	request.Protocol = "http"
	request.HttpRequest = ctx.Request

	// header 透传
	headers := make(map[string]string)
	for key, value := range ctx.Request.Header {
		headers[key] = value[0]
	}
	remoteAddr, _ := library.GetExternalIP()
	headers["REMOTE_ADDR"]          = remoteAddr.String()
	headers["HTTP_VIA"]             = remoteAddr.String()
	if xForward, exist := headers["HTTP_X_FORWARDED_FOR"]; exist && xForward != ""{
		prefix := ctx.Request.Header["HTTP_X_FORWARDED_FOR"][0]
		headers["HTTP_X_FORWARDED_FOR"] = fmt.Sprintf("%s,%s", prefix, ctx.Request.RemoteAddr)
	}else{
		// 本机测试 [::1]:12345
		if strings.HasPrefix(ctx.Request.RemoteAddr, "[::1]") {
			headers["HTTP_X_FORWARDED_FOR"] = remoteAddr.String()
		}else {
			headers["HTTP_X_FORWARDED_FOR"] = strings.Split(ctx.Request.RemoteAddr, ":")[0]
		}
	}
	request.Headers = headers
	if strings.Contains(ctx.Request.Proto, "http") && !strings.Contains(ctx.Request.Proto, "HTTPS") {
		request.Protocol = "http"
		request.Port     = 80
	}else if !strings.Contains(ctx.Request.Proto, "http") && strings.Contains(ctx.Request.Proto, "HTTPS"){
		request.Protocol = "https"
		request.Port     = 443
	}
	request.Method   = ctx.Request.Method
	request.Path     = ctx.Request.RequestURI
	request.Host     = ctx.Request.Host
	request.FormData = make(map[string]string)
	if uriRow := strings.Split(ctx.Request.RequestURI, "?"); len(uriRow) == 2 {
		request.Path = uriRow[0]
		request.QueryString = uriRow[1]
	}
	if hostRow := strings.Split(ctx.Request.Host, ":"); len(hostRow) == 2 {
		request.Host = hostRow[0]
	}

	// application/x-www-form-urlencoded、application/form 形式 TODO 后续再进行支持 JSON 形式的 HTTP 请求
	ctx.Request.ParseForm()
	for k, v := range ctx.Request.PostForm {
		if len(v) <= 0 {
			continue
		}
		// TODO 可能会存在同名参数值为 list 的数据丢失问题
		request.FormData[k] = v[0]
	}

	// 拷贝原来的请求到请求对象上
	if clientIP, _, err := net.SplitHostPort(ctx.Request.RemoteAddr); err == nil {
		if prior, ok := request.Headers["X-Forwarded-For"]; ok {
			clientIP = prior + ", " + clientIP
		}
		request.Headers["X-Forwarded-For"] = clientIP
	}

	//////////////
	// 定制增加自己的额外参数
	//////////////
	// 触发词检测，以及配置内容的定制替换（此处主要是替换 HOST 的内容）
	for _, cfgItem := range cfg.Items {
		if cfgItem.Type == library.CONFIG_TRIGGER_POSITION_HOST {
			//
		}else if cfgItem.Type == library.CONFIG_TRIGGER_POSITION_URI {
			//
		}else if cfgItem.Type == library.CONFIG_TRIGGER_POSITION_QUERYSTRING {
			// 拿到 ctx 的请求 QueryString 信息，然后判断是否有触发行为，据此做定制替换
			for kName, kvalue := range GetRequestParameters(ctx) {
				if kName == cfgItem.Trigger || kvalue == cfgItem.Trigger ||
					fmt.Sprintf("%s=%s", kName, kvalue) == cfgItem.Trigger {
					// 替换
					request.Protocol = cfgItem.TargetProtocol
					request.Host = cfgItem.TargetHost
					request.Port = cfgItem.TargetPort
					break
				}
			}
		}else if cfgItem.Type == library.CONFIG_TRIGGER_POSITION_POSTFORM {
			// 拿到所有 post 参数，并进行替换处理
			for pName, pValue := range request.FormData {
				if pName == cfgItem.Trigger || pValue == cfgItem.Trigger {
					// 开始替换
					request.Protocol = cfgItem.TargetProtocol
					request.Host = cfgItem.TargetHost
					request.Port = cfgItem.TargetPort
					request.Path  = cfgItem.TargetPath
					break
				}
			}
		}
	}

    // 返回最终定制化完成的 "请求体"
	return request
}

func GetRequestParameters(ctx *gin.Context) (parameters map[string]string) {
	// 初始化字典用于存储 GET 参数
	parameters = make(map[string]string)

	queryString := ctx.Request.URL.RawQuery
	if queryString == "" || !strings.Contains(queryString, "=") {
		return
	}
	rows := strings.Split(queryString, "&")
	for _, rowString := range rows{
		if rowString == "" || !strings.Contains(rowString, "=") {
			continue
		}
		row := strings.Split(rowString, "=")
		if len(row) != 2 || row[0] == "" {
			continue
		}
		parameters[row[0]] = row[1]
	}
	return
}