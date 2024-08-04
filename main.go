package main

import (
	"Anvarjon-33/Nuxt_Go.port/api/auth"
	"Anvarjon-33/Nuxt_Go.port/db"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


var par = make(chan map[string]string, 1)
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}
func main() {
	gin.ForceConsoleColor()

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/data", func(c *gin.Context) {
		var res = make(map[string]string)
		params := c.Request.URL.Query()
		for key, val := range params {
			res[key] = strings.Join(val, "")
		}
		par <- res
	})

	auth.Auth(r)

	// ######################################
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware

		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	// ######################################

	r.GET("/JSONP?callback=x", func(c *gin.Context) {

		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	r.POST("/data", func(c *gin.Context) {

		var res = make(map[string]string)
		params := c.Request.URL.Query()
		for key, val := range params {
			res[key] = strings.Join(val, "")
			res[key] = strings.Trim(res[key], " ")
			if res[key] != "" {
				par <- res
			}
		}
		select {
		case p := <-par:
			fmt.Println(p["name"])
			res, _ := db.GetUser(p["name"])
			c.JSON(http.StatusOK, gin.H{
				"message": res,
			})
		default:
			fmt.Println("Default Response !")
			res, _ := db.GetUser_()
			c.JSON(http.StatusOK, gin.H{
				"message": res,
			})
		}
	})

	// ######################################

	type FormColor struct {
		Colors []string `form:"colors[]"`
	}

	r.POST("/test", func(context *gin.Context) {
		var formColor FormColor
		err := context.ShouldBind(&formColor)
		if err != nil {
			panic(err)
		}
		res := gin.H{"colors": formColor.Colors}
		fmt.Println(res)
		context.JSON(200, gin.H{"colors": formColor.Colors})
	})

	// ###########
	/*
		r.POST("/postman", func(c *gin.Context) {
			post, err := io.ReadAll(c.Request.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(post))
			c.JSON(http.StatusOK, gin.H{
				"message": string(post),
			})
		})
	*/
	r.Run("192.168.1.3:2222") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
