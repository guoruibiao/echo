package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/echo/models"
	"github.com/guoruibiao/echo/models/parser"
	"io/ioutil"
	"net/http"
)

func Proxy( ctx *gin.Context) {
	// 获取 CUID 或其他唯一标识符
	cuid := parser.FindCuid(ctx)
	cuid = "tiger"
	// 根据 CUID 获取配置映射表
	configuration, err := parser.GetConfigurations(cuid)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	// 组装成定制的请求体
	req := parser.GetRequest(ctx, &configuration)
	// TODO models.Do()方法中，对于未匹配到的请求体需要做更细致的封装，后续可置于 parser.GetRequest 或者 models.Do中
	if err, resp := models.Do(req); err != nil {
		fmt.Println("请求失败", err)
		ctx.JSON(http.StatusOK, err)
	}else{
		fmt.Println("--------------------")
		fmt.Printf("response=%#v\n", resp)
		fmt.Println("--------------------")
		defer resp.Body.Close()
		if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
			ctx.Writer.Write(bytes)
		}else{
			fmt.Println("解析失败", err)
			ctx.JSON(http.StatusOK, err)
		}
	}

}
