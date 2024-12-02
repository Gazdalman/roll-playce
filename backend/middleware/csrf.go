package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"net/http"
)

func CSRFMiddleware(secretKey string) gin.HandlerFunc {
	csrfMiddleware := csrf.Protect(
		[]byte(secretKey),
		csrf.Secure(false), // Set to true for HTTPS
		csrf.RequestHeader("X-CSRF-Token"),
	)

	return func(c *gin.Context) {
		csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the CSRF token in Gin's context
			token := csrf.Token(r)
			c.Set("csrf_token", token)

			// Pass the request and response back to Gin's context
			c.Request = r
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	}
}
