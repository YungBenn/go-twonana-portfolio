//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/YungBenn/go-twonana-portfolio/config"
	"github.com/YungBenn/go-twonana-portfolio/internal/http/handler"
	"github.com/YungBenn/go-twonana-portfolio/internal/nft"
	"github.com/YungBenn/go-twonana-portfolio/internal/infra/cloudinary"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitNft(db *mongo.Database, env config.EnvVars) *handler.NftHandler {
	wire.Build(
		nft.NewNftRepository,
		nft.NewNftService,
		cloudinary.NewMediaService,
		handler.NewNftHandler,
	)

	return &handler.NftHandler{}
}