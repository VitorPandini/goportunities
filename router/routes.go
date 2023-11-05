package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the opening page",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the opening POST",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the opening delete page",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the opening put page",
			})
		})
		v1.GET("/opening/id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the opening page for ID",
			})
		})
	}
}
