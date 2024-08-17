package main

import (
	"Anvarjon-33/Nuxt_Go/api/auth"
	"Anvarjon-33/Nuxt_Go/db"
	"Anvarjon-33/Nuxt_Go/middleware"
	"Anvarjon-33/Nuxt_Go/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

var (
	par    chan map[string]string
	secret string
	mode   string
)

func init() {
	secret = os.Getenv("SECRET")
	par = make(chan map[string]string, 1)
}

func main() {

	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/csrf-header", routes.Csrf())

	r.GET("/data", routes.Csrf(), func(c *gin.Context) {
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
			c.JSON(http.StatusAccepted, gin.H{
				"message": res,
			})
		}
	}, middleware.Debugger())
	auth.Auth(r)
	r.Run(":2222") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, HEAD")
		c.Writer.Header().Set("X-Requested-With", "XMLHttpRequest")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

//while true;
//do
//nc -q 1 192.168.1.3 1213;
//nc -q 1 192.168.1.3 1212;
//nc -q 1 192.168.1.3 1234;
//done
