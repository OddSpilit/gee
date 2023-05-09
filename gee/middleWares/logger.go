package middleWares

import (
	"gee/gee"
	"log"
	"time"
)

func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v \n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
