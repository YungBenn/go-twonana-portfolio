//go:build wireinject
// +build wireinject

package di

import (
	"github.com/YungBenn/go-twonana-portfolio/internal/admin"
	"github.com/YungBenn/go-twonana-portfolio/internal/api/handler"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/wire"
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