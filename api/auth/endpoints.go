package auth

import (
	"Anvarjon-33/Nuxt_Go/app/utils"
	"Anvarjon-33/Nuxt_Go/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var user = model.AuthUser{}

func Auth(r *gin.Engine) {
	api := r.Group("/api/auth")
	{
		api.POST("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"token": "",
			})
		})
		api.POST("/logout", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/logout",
			})
		})
		api.POST("/refresh-token", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "refresh-token",
			})
		})
		api.POST("/register", func(context *gin.Context) {
			var user *model.AuthUser
			context.Bind(&user)

			err := model.DB.Where(&model.AuthUser{Email: user.Email}).First(&user).Error
			if err != nil {
				newUUID, _ := uuid.NewUUID()
				h, err := utils.HashPassword(user.Password)
				if err != nil {
					panic(err)
				}
				model.DB.Create(&model.AuthUser{
					ID: newUUID, Name: user.Name, Email: user.Email, Password: h,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": "Email already ahs been taken",
				})
			}

		})
		api.GET("/session", func(context *gin.Context) {
			sess := context.Request.Cookies()
			auth := context.Request.Header.Get("Authorization")

			fmt.Println("SEESION : ", sess)
			fmt.Println("AUTHORIZATION : ", auth)

			context.JSON(http.StatusOK, gin.H{
				"message": "Something is wrong",
			})
			//context.JSON(http.StatusOK, gin.H{
			//	"message": "/sesion String MEssage",
			//})
		})
	}
}
