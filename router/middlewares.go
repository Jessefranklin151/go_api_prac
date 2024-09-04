package router

import "github.com/gin-gonic/gin"

func createMiddlewares(router *gin.Engine) {
	router.Use(securityMiddleware)
}

func securityMiddleware(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if token != "token certo" {
		c.JSON(401, gin.H{
			"error": "Token inválido",
		})
		c.Abort()
	}

	c.Next()

}
