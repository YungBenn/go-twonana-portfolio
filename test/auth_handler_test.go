package test

import (
	"testing"

	"github.com/YungBenn/go-twonana-portfolio/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

func TestAuthHandler_Register(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.AuthHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Register(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHandler.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthHandler_Login(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.AuthHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHandler.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthHandler_Logout(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *handler.AuthHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Logout(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHandler.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
