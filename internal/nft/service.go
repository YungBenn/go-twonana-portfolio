package nft

import (
	"context"

	"github.com/YungBenn/go-twonana-portfolio/internal/cloudinary"
	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
)

type NftService interface {
	CreateNft(ctx context.Context, dto CreateNftRequest) (*Nft, *response.Error)
	FindAllNft(ctx context.Context, limit int, page int) ([]Nft, *response.Error)
	FindAllCategory(ctx context.Context) ([]string, *response.Error)
	FindNftByCategory(ctx context.Context, category string) ([]Nft, *response.Error)
	FindOneNft(ctx context.Context, id string) (*Nft, *response.Error)
	UpdateNft(ctx context.Context, id string, dto UpdateNftRequest) (*Nft, *response.Error)
	DeleteNft(ctx context.Context, id string) *response.Error
}

type nftService struct {
	repo  NftRepository
	media cloudinary.MediaService
}

// FindNftByCategory implements NftService.
func (s *nftService) FindNftByCategory(ctx context.Context, category string) ([]Nft, *response.Error) {
	return s.repo.FindNftByCategory(ctx, category)
}

// FindAllCategory implements NftService.
func (s *nftService) FindAllCategory(ctx context.Context) ([]string, *response.Error) {
	return s.repo.FindAllCategory(ctx)
}

// CreateNft implements NftService.
func (s *nftService) CreateNft(ctx context.Context, dto CreateNftRequest) (*Nft, *response.Error) {
	entity := Nft{
		Title:          dto.Title,
		Category:       dto.Category,
		Description:    dto.Description,
		Software:       dto.Software,
		Size:           dto.Size,
		NftFormat:      dto.NftFormat,
		MarketplaceUrl: dto.MarketplaceUrl,
		CountdownDays:  dto.CountdownDays * 86400000,
	}

	if dto.ImageUrl != nil {
		image, err := s.media.Upload(*dto.ImageUrl)
		if err != nil {
			return nil, err
		}

		if image != nil {
			entity.ImageUrl = image
		}
	}

	data, err := s.repo.InsertNft(ctx, entity)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DeleteNft implements NftService.
func (s *nftService) DeleteNft(ctx context.Context, id string) *response.Error {
	_, err := s.repo.FindOneNft(ctx, id)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteNft(ctx, id); err != nil {
		return err
	}

	return nil
}

// FindAllNft implements NftService.
func (s *nftService) FindAllNft(ctx context.Context, limit int, page int) ([]Nft, *response.Error) {
	return s.repo.FindAllNft(ctx, limit, page)
}

// FindOneNft implements NftService.
func (s *nftService) FindOneNft(ctx context.Context, id string) (*Nft, *response.Error) {
	return s.repo.FindOneNft(ctx, id)
}

// UpdateNft implements NftService.
func (s *nftService) UpdateNft(ctx context.Context, id string, dto UpdateNftRequest) (*Nft, *response.Error) {
	entity := Nft{
		Title:          dto.Title,
		Category:       dto.Category,
		Description:    dto.Description,
		Software:       dto.Software,
		Size:           dto.Size,
		NftFormat:      dto.NftFormat,
		MarketplaceUrl: dto.MarketplaceUrl,
		CountdownDays:  dto.CountdownDays,
	}

	if dto.ImageUrl != nil {
		image, err := s.media.Upload(*dto.ImageUrl)
		if err != nil {
			return nil, err
		}

		if entity.ImageUrl != nil {
			_, err := s.media.Delete(*entity.ImageUrl)
			if err != nil {
				return nil, err
			}
		}

		if image != nil {
			entity.ImageUrl = image
		}
	}

	data, err := s.repo.UpdateNft(ctx, id, entity)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewNftService(repository NftRepository, media cloudinary.MediaService) NftService {
	return &nftService{repository, media}
}
