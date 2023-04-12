package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func (c *gin.Context) {

		t := time.Now()
		verb := c.Request.Method
		path := c.Request.RequestURI

		c.Next()

		var sizeInBytes int
		if c.Writer != nil {
			sizeInBytes = c.Writer.Size()
		}

		elapsed := time.Since(t)

		fmt.Println("--------------------- LOGGER ---------------------")
		fmt.Printf("time: %v\npath: %s\nverb: %s\nsize: %d\nelapsed time: %v", t, path, verb, sizeInBytes, elapsed)
		fmt.Println("\n--------------------------------------------------")

	}

}