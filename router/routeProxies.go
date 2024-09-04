package router

import (
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/config"
	"github.com/gin-gonic/gin"
)

func setupProxies(router *gin.Engine) {
	router.SetTrustedProxies(config.GetTrustedProxies())
}
