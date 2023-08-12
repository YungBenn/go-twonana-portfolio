package admin

import (
	"context"
	"errors"

	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository interface {
	FindAdminByUsername(ctx context.Context, username string) (*Admin, *response.Error)
	InsertAdmin(ctx context.Context, admin Admin) (*Admin, *response.Error)
}

type adminRepository struct {
	db *mongo.Database
}

// FindAdminByUsername implements AdminRepository.
func (repo *adminRepository) FindAdminByUsername(ctx context.Context, username string) (*Admin, *response.Error) {
	coll := repo.db.Collection("admins")

	var admin Admin
	err := coll.FindOne(ctx, bson.M{"username": username}).Decode(&admin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, response.NewError(500, err)
		}
	}

	return &admin, nil
}

// InsertAdmin implements AdminRepository.
func (repo *adminRepository) InsertAdmin(ctx context.Context, admin Admin) (*Admin, *response.Error) {
	coll := repo.db.Collection("admins")

	res, err := coll.InsertOne(ctx, admin)
	if err != nil {
		return nil, response.NewError(500, err)
	}

	newID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, response.NewError(500, errors.New("failed to get inserted ID"))
	}

	newAdmin := &Admin{
		ID:       newID,
		Username: admin.Username,
		Password: admin.Password,
	}

	return newAdmin, nil
}

func NewAdminRepository(db *mongo.Database) AdminRepository {
	return &adminRepository{db}
}
