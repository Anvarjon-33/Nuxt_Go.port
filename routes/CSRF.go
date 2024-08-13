package routes

import (
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Csrf() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Writer.Header().Set("X-CSRF-Token", "fdvfsdvfdvlihekjllfkjewrfhlkerwjglerhl")

		//context.JSON(http.StatusOK, gin.H{
		//	"token": "fdvfsdvfdvlihekjllfkjewrfhlkerwjglerhl",
		//})
	}
}
