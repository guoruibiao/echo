package dao

import "net/http"

type Request struct {
	HttpRequest *http.Request
	Protocol string
	Method string
	Path   string
	Host   string
	Port   int
	QueryString string
	/*
	 * IP 相关处理
	 * 1. 没有使用代理服务器
	 *    REMOTE_ADDR            真实 IP
	 *    HTTP_VIA               没有数值或者不显示
	 *    HTTP_X_FORWARDED_FOR   没数值或者不显示
	 *
	 * 2. 使用透明代理 √
	 *    REMOTE_ADDR            最后一个代理服务器的 IP
	 *    HTTP_VIA               代理服务器 IP
	 *    HTTP_X_FORWARDED_FOR   真实 IP，经过多个代理服务器时，尾部追加，无法"隐身"
	 *
	 * 3. 使用普通匿名代理
	 *    REMOTE_ADDR            最后一个代理服务器 IP
	 *    HTTP_VIA               代理服务器 IP
	 *    HTTP_X_FORWARDED_FOR   代理 IP，经过多个代理服务器时，尾部追加，无法"隐身"
	 *
	 * 4. 使用欺骗性代理
	 *    REMOTE_ADDR            代理服务器 IP
	 *    HTTP_VIA               代理服务器 IP
	 *    HTTP_X_FORWARDED_FOR   随机 IP，经过多个代理服务器时，尾部追加，可以"隐身"
	 *
	 * 5. 使用高匿名代理
	 *    REMOTE_ADDR            代理服务器 IP
	 *    HTTP_VIA               没数值或者不显示
	 *    HTTP_X_FORWARDED_FOR   没数值或者不显示
	 *
	 */
	Headers map[string]string
	FormData map[string]string
}
