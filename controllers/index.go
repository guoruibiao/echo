package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
