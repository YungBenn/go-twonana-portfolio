//go:build wireinject
// +build wireinject

package di

import (
	"github.com/YungBenn/go-twonana-portfolio/config"
	"github.com/YungBenn/go-twonana-portfolio/internal/api/handler"
	"github.com/YungBenn/go-twonana-portfolio/internal/infra/cloudinary"
	"github.com/YungBenn/go-twonana-portfolio/internal/nft"
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