package middleware

import (
	"net/http"
	"os"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/web"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {

	return func (c *gin.Context) {

		token := c.GetHeader("TOKEN")
		if token == "" {
			web.NewApiError(c, http.StatusUnauthorized, "unauthorized", "token not found")
			c.Abort()
			return
		}

		if token != os.Getenv("TOKEN") {
			web.NewApiError(c, http.StatusUnauthorized, "unauthorized", "invalid token")
			c.Abort()
			return
		}

		c.Next()

	}

}
