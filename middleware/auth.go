package middleware

import (
	"Anvarjon-33/Nuxt_Go/db"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func TestMid() gin.HandlerFunc {
	return func(c *gin.Context) {

		db.Redis.Set(nil, "Example_key", "Example_val", time.Second*10)

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
