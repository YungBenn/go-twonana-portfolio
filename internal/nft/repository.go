package nft

import (
	"context"
	"errors"

	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
	"github.com/YungBenn/go-twonana-portfolio/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NftRepository interface {
	InsertNft(ctx context.Context, nft Nft) (*Nft, *response.Error)
	FindAllNft(ctx context.Context, limit int, page int) ([]Nft, *response.Error)
	FindAllCategory(ctx context.Context) ([]string, *response.Error)
	FindOneNft(ctx context.Context, id string) (*Nft, *response.Error)
	FindNftByCategory(ctx context.Context, category string) ([]Nft, *response.Error)
	// FindNftByTitle(ctx context.Context, title string) (*Nft, *response.Error)
	UpdateNft(ctx context.Context, id string, nft Nft) (*Nft, *response.Error)
	DeleteNft(ctx context.Context, id string) *response.Error
}

type nftRepository struct {
	db *mongo.Database
	coll *mongo.Collection
}

// FindAllCategory implements NftRepository.
func (repo *nftRepository) FindAllCategory(ctx context.Context) ([]string, *response.Error) {
	coll := repo.db.Collection("nfts")

	res, err := coll.Distinct(ctx, "category", bson.D{})
	if err != nil {
		return nil, response.NewError(500, err)
	}

	categories := make([]string, len(res))
	for i, r := range res {
		if category, ok := r.(string); ok {
			categories[i] = category
		} else {
			return nil, response.NewError(500, err)
		}
	}

	return categories, nil

}

// FindNftByCategory implements NftRepository.
func (repo *nftRepository) FindNftByCategory(ctx context.Context, category string) ([]Nft, *response.Error) {
	coll := repo.db.Collection("nfts")

	res, err := coll.Find(ctx, bson.M{"category": category})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, response.NewError(404, err)
		}
	}

	nfts := make([]Nft, 0)
	if err = res.All(ctx, &nfts); err != nil {
		return nil, response.NewError(500, err)
	}

	return nfts, nil
}

func NewNftRepository(db *mongo.Database) NftRepository {
	return &nftRepository{
		db:   db,
		coll: db.Collection("nfts"),
	}
}

func (repo *nftRepository) InsertNft(ctx context.Context, nft Nft) (*Nft, *response.Error) {
	coll := repo.coll

	res, err := coll.InsertOne(ctx, nft)
	if err != nil {
		return nil, response.NewError(500, err)
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, response.NewError(500, errors.New("failed to get inserted ID"))
	}

	newNFT := &Nft{
		ID:             newID,
		Title:          nft.Title,
		Category:       nft.Category,
		Description:    nft.Description,
		Software:       nft.Software,
		Size:           nft.Size,
		NftFormat:      nft.NftFormat,
		MarketplaceUrl: nft.MarketplaceUrl,
		CountdownDays:  nft.CountdownDays,
		ImageUrl:       nft.ImageUrl,
	}

	return newNFT, nil
}

func (repo *nftRepository) FindAllNft(ctx context.Context, limit int, page int) ([]Nft, *response.Error) {
	coll := repo.coll

	res, err := coll.Find(ctx, bson.M{}, utils.NewMongoPaginate(limit, page).GetPaginatedOpts())
	if err != nil {
		return nil, response.NewError(500, err)
	}

	nfts := make([]Nft, 0)
	if err = res.All(ctx, &nfts); err != nil {
		return nil, response.NewError(500, err)
	}

	return nfts, nil
}

func (repo *nftRepository) FindOneNft(ctx context.Context, id string) (*Nft, *response.Error) {
	coll := repo.coll

	objectId, _ := primitive.ObjectIDFromHex(id)

	var nft Nft
	err := coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&nft)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, response.NewError(404, err)
		}
	}

	return &nft, nil
}

func (repo *nftRepository) UpdateNft(ctx context.Context, id string, nft Nft) (*Nft, *response.Error) {
	coll := repo.coll

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": nft}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, response.NewError(500, err)
	}

	var updatedNft Nft
	err = coll.FindOne(ctx, filter).Decode(&updatedNft)
	if err != nil {
		return nil, response.NewError(500, err)
	}

	return &updatedNft, nil
}

func (repo *nftRepository) DeleteNft(ctx context.Context, id string) *response.Error {
	coll := repo.coll

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return response.NewError(500, err)
	}

	return nil
}
