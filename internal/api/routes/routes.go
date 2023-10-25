package routes

import (
	"github.com/YungBenn/go-twonana-portfolio/config"
	"github.com/YungBenn/go-twonana-portfolio/internal/api/middleware"
	injectorAdmin "github.com/YungBenn/go-twonana-portfolio/internal/di/admin"
	injectorNft "github.com/YungBenn/go-twonana-portfolio/internal/di/nft"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupNFTRoutes(app *fiber.App, db *mongo.Database, env config.EnvVars, store *session.Store) {
	nft := injectorNft.InitNft(db, env)

	nftRouter := app.Group("/api/v1/nft")
	nftRouter.Use(middleware.IsAuth(store))

	nftRouter.Post("/", nft.CreateNft)
	nftRouter.Get("/", nft.GetAllNft)
	nftRouter.Get("/category/:category", nft.GetAllNftByCategory)
	nftRouter.Get("/category", nft.GetAllCategory)
	nftRouter.Get("/:id", nft.GetOneNft)
	nftRouter.Put("/:id", nft.UpdateNft)
	nftRouter.Delete("/:id", nft.DeleteNft)
}

func SetupAuthRoutes(app *fiber.App, db *mongo.Database, store *session.Store) {
	admin := injectorAdmin.InitAdmin(db, store)

	authRouter := app.Group("api/v1/auth")

	authRouter.Post("/register", admin.Register)
	authRouter.Post("/login", admin.Login)
	authRouter.Delete("/logout", admin.Logout)
}