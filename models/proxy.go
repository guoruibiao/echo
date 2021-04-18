package models

import (
	"fmt"
	"github.com/guoruibiao/echo/library"
	"github.com/guoruibiao/echo/models/dao"
	"github.com/guoruibiao/gorequests"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func Do(request dao.Request) (err error, res *http.Response) {
	// 解决回环问题 MOCK 代码
	if request.Host == "localhost" && request.Port == 0{

		if clientIp, _, err := net.SplitHostPort(request.HttpRequest.RemoteAddr); err == nil {
			if prior, ok := request.HttpRequest.Header["X-Forward-For"]; ok {
				clientIp = strings.Join(prior, ",") + "," + clientIp
			}
			request.HttpRequest.Header.Set("X-Forward-For", clientIp)
			request.Headers["X-Forward-For"] = clientIp
		}

		request.Method = request.HttpRequest.Method
		request.Host, _, _ = net.SplitHostPort(request.HttpRequest.Host)
		request.Port, _ = strconv.Atoi(request.HttpRequest.URL.Port())
		request.Protocol = strings.Split(request.HttpRequest.Proto, "/")[0]
		request.Path   = request.HttpRequest.URL.Path
		request.QueryString = request.HttpRequest.URL.RawQuery

		// TODO 默认配置都打到 httpbin.org 上，后续应该做成可配置形式
		request.Protocol = "https"
		request.Host  = "httpbin.org"
		request.Port  = 443
		request.Path  = "/" + strings.ToLower(request.Method)
	}

	url := fmt.Sprintf("%s://%s:%d%s?%s",
		request.Protocol, request.Host, request.Port, request.Path, request.QueryString)
	if (request.Protocol == library.PROTOCOL_HTTP && request.Port == 80) ||
		(request.Protocol == library.PROTOCOL_HTTPS && request.Port == 443) {
		url = fmt.Sprintf("%s://%s%s?%s",
			request.Protocol, request.Host, request.Path, request.QueryString)
	}
	fmt.Println("--------------------")
	fmt.Printf("request=%#v\n", request)
	fmt.Println("--------------------")


	response, err := gorequests.NewRequest(request.Method, url).
		Headers(request.Headers).
		Form(request.FormData).
		DoRequest()

	if err != nil {
		return
	}else {
		res = response.Response
	}
	return
}