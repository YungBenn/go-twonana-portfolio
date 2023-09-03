package test

import (
	"testing"

	"github.com/YungBenn/go-twonana-portfolio/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

func TestNftHandler_CreateNft(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.CreateNft(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.CreateNft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNftHandler_GetAllNft(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAllNft(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.GetAllNft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNftHandler_GetOneNft(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetOneNft(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.GetOneNft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNftHandler_UpdateNft(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.UpdateNft(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.UpdateNft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNftHandler_DeleteNft(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.DeleteNft(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.DeleteNft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNftHandler_GetAllCategory(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAllCategory(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.GetAllCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNftHandler_GetAllNftByCategory(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAllNftByCategory(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.GetAllNftByCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNftHandler_GetNftByTitle(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.NftHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetNftByTitle(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("NftHandler.GetNftByTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
