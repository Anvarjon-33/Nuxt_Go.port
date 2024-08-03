package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Auth(r *gin.Engine) {
	api := r.Group("/api/auth")
	{
		api.POST("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/login",
			})
		})
		api.POST("/logout", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/logout",
			})
		})
		api.POST("/register", func(context *gin.Context) {
			post, err := io.ReadAll(context.Request.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(post))
			context.JSON(http.StatusOK, gin.H{
				"message": "/register",
			})
		})
		api.POST("/session", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/sesion String MEssage",
			})
		})
	}
}
