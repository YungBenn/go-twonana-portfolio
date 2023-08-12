package nft

import "go.mongodb.org/mongo-driver/bson/primitive"

type Nft struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title          string             `bson:"title,omitempty" json:"title,omitempty"`
	ImageUrl       *string            `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Category       string             `bson:"category,omitempty" json:"category,omitempty"`
	Description    string             `bson:"description,omitempty" json:"description,omitempty"`
	Software       string             `bson:"software,omitempty" json:"software,omitempty"`
	Size           string             `bson:"size,omitempty" json:"size,omitempty"`
	NftFormat      string             `bson:"nft_format,omitempty" json:"nft_format,omitempty"`
	MarketplaceUrl string             `bson:"marketplace_url,omitempty" json:"marketplace_url,omitempty"`
	CountdownDays  int                `bson:"countdown_days,omitempty" json:"countdown_days,omitempty"`
}
