package main

import (
	"Anvarjon-33/Nuxt_Go.port/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var par = make(chan map[string]string, 1)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	//r.GET("/data", func(c *gin.Context) {
	//	var res = make(map[string]string)
	//	params := c.Request.URL.Query()
	//	for key, val := range params {
	//		res[key] = strings.Join(val, "")
	//	}
	//	par <- res
	//})

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
