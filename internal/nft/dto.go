package nft

import "mime/multipart"

type CreateNftRequest struct {
	Title          string                `form:"title" json:"title"`
	ImageUrl       *multipart.FileHeader `form:"image_url" json:"image_url" swaggerignore:"true"`
	Category       string                `form:"category" json:"category"`
	Description    string                `form:"description" json:"description"`
	Software       string                `form:"software" json:"software"`
	Size           string                `form:"size" json:"size"`
	NftFormat      string                `form:"nft_format" json:"nft_format"`
	MarketplaceUrl string                `form:"marketplace_url" json:"marketplace_url"`
	CountdownDays  int                   `form:"countdown_days" json:"countdown_days"`
}

type UpdateNftRequest struct {
	Title          string                `form:"title,omitempty" json:"title,omitempty"`
	ImageUrl       *multipart.FileHeader `form:"image_url,omitempty" json:"image_url,omitempty" swaggerignore:"true"`
	Category       string                `form:"category,omitempty" json:"category,omitempty"`
	Description    string                `form:"description,omitempty" json:"description,omitempty"`
	Software       string                `form:"software,omitempty" json:"software,omitempty"`
	Size           string                `form:"size,omitempty" json:"size,omitempty"`
	NftFormat      string                `form:"nft_format,omitempty" json:"nft_format,omitempty"`
	MarketplaceUrl string                `form:"marketplace_url,omitempty" json:"marketplace_url,omitempty"`
	CountdownDays  int                   `form:"countdown_days,omitempty" json:"countdown_days,omitempty"`
}
