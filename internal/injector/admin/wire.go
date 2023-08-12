//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/YungBenn/go-twonana-portfolio/internal/http/handler"
	"github.com/YungBenn/go-twonana-portfolio/internal/admin"
	"github.com/google/wire"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitAdmin(db *mongo.Database, store *session.Store) *handler.AuthHandler {
	wire.Build(
		admin.NewAdminRepository,
		admin.NewAdminService,
		handler.NewAuthHandler,
	)

	return &handler.AuthHandler{}
}