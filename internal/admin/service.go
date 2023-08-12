package admin

import (
	"context"
	"errors"

	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CreateAdmin(ctx context.Context, dto AdminRequest) (*Admin, *response.Error)
	Login(ctx context.Context, dto AdminRequest) (*Admin, *response.Error)
}

type adminService struct {
	repo AdminRepository
}

// Login implements AdminService.
func (s *adminService) Login(ctx context.Context, dto AdminRequest) (*Admin, *response.Error) {
	dataAdmin, err := s.repo.FindAdminByUsername(ctx, dto.Username)
	if err != nil {
		return nil, err
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(dataAdmin.Password), []byte(dto.Password))
	if errBcrypt != nil {
		return nil, response.NewError(400, errors.New("username or password is invalid"))
	}

	return dataAdmin, nil
}

// CreateAdmin implements AdminService.
func (s *adminService) CreateAdmin(ctx context.Context, dto AdminRequest) (*Admin, *response.Error) {
	checkUsername, _ := s.repo.FindAdminByUsername(ctx, dto.Username)
	if checkUsername != nil {
		return nil, response.NewError(409, errors.New("username is already exists"))
	}

	hashedPass, errHashed := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if errHashed != nil {
		return nil, response.NewError(500, errHashed)
	}

	entity := Admin{
		Username: dto.Username,
		Password: string(hashedPass),
	}

	data, err := s.repo.InsertAdmin(ctx, entity)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewAdminService(repo AdminRepository) AdminService {
	return &adminService{repo}
}
