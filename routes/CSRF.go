package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var r = gin.Default()

func Auth() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"try": "Example",
		})
	}
}
