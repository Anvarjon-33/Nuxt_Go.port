package middleware

import (
	"Anvarjon-33/Nuxt_Go/db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func TestMid() gin.HandlerFunc {
	return func(c *gin.Context) {

		db.Redis.Set(context.Background(), "Example_key", "Example_val", time.Second*10)

		t := time.Now()

		// before request

		c.Next()

		// after request

		latency := time.Since(t)
		log.Print("Time for response: -------------------------------->  ", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func Debugger() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := c.Request.Cookies()
		auth := c.Request.Header.Get("Authorization")
		csrf := c.Request.Header.Get("X-CSRF-Token")
		csrf2 := c.Request.Header.Get("CSRF-Token")

		fmt.Println("SESS : ", sess)
		fmt.Println("AUTH : ", auth)
		fmt.Println("CSRF : ", csrf)
		fmt.Println("CSRF2 : ", csrf2)
	}
}
