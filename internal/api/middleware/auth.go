package middleware

import (
	"log"
	"net/http"

	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func IsAuth(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(response.Response(
				http.StatusInternalServerError, 
				http.StatusText(http.StatusInternalServerError), 
				err,
			))
		}
	
		username := sess.Get("username")
		if username == nil {
			return c.Status(http.StatusUnauthorized).JSON(response.Response(
				http.StatusUnauthorized, 
				http.StatusText(http.StatusUnauthorized), 
				err,
			))
		}

		if username != nil {
			log.Println("session ada")
		}
	
		return c.Next()
	}
}