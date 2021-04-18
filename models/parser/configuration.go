package parser
// 实现代理配置内容的解析，以实现区分用户维度的前向代理功能

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/echo/library"
	"github.com/guoruibiao/echo/models/dao"
	"strings"
)

const (
	CUID_KEY    string  = "cuid"
	CUID_PREFIX string  = "cuid="
)

// 从原始请求中抽取出 唯一标识符信息（这里指定为CUID）
func FindCuid(ctx *gin.Context) (cuid string) {
	// 1. 从 POST 参数中查找 application/x-www-form-urlencoded、application/form 形式
	ctx.Request.ParseForm()
	for pName, pValues := range ctx.Request.PostForm {
		if len(pValues) <= 0 {
			continue
		}
		if strings.ToLower(pName) != CUID_KEY {
			continue
		}
		for _, target := range pValues {
			if target == "" {
				continue
			}
			return target
		}
	}

	// 2. 从 GET 参数中查找
	queryString := ctx.Request.URL.RawQuery
	if strings.Contains(strings.ToLower(queryString), CUID_KEY) {
		items := strings.Split(queryString, "&")
		if len(items) >= 1 {
			for _, item := range items {
				if !strings.HasPrefix(strings.ToLower(item), CUID_PREFIX) {
					continue
				}
				row := strings.Split(item, "=")
				if len(row) != 2 {
					continue
				}
				if row[1] == "" {
					continue
				}
				return row[1]
			}
		}
	}

	// 3. 从 header 中查找
	for hName, hValues := range ctx.Request.Header {
		if strings.ToLower(hName) != CUID_KEY {
			continue
		}
		for _, hValue := range hValues {
			if hValue == "" {
				continue
			}
			return hValue
		}
	}

	// 4. 从 cookie 中查找
	for _, cookie := range ctx.Request.Cookies() {
		if strings.ToLower(cookie.Name) != CUID_KEY {
			continue
		}
		if cookie.Value == "" {
			continue
		}
		return cookie.Value
	}

	return
}

func GetConfigurations(cuid string) (configurations dao.Configuration, err error) {
	// TODO mock 删除
	if cuid == "tiger" {
		configurations.Cuid = cuid
		// 对 GET 参数中含有触发词 triggerget 进行内容的正向代理
		configurations.Items = append(configurations.Items, dao.ConfigItem{
			Type: library.CONFIG_TRIGGER_POSITION_QUERYSTRING,
			Trigger: "triggerget",
			TargetProtocol: library.PROTOCOL_HTTP,
			TargetHost: "helloworld.com",
			TargetPort: 80,
			TargetPath: "/",
		})

		// 对 POST 参数中含有触发词 triggerpost 进行内容的正向代理
		configurations.Items = append(configurations.Items, dao.ConfigItem{
			Type: library.CONFIG_TRIGGER_POSITION_POSTFORM,
			Trigger: "triggerpost",
			TargetProtocol: library.PROTOCOL_HTTPS,
			TargetHost: "httpbin.org",
			TargetPort: 443,
			TargetPath: "/post",
		})
	}

	return
}