package cloudinary

import (
	"context"
	"errors"

	"github.com/YungBenn/go-twonana-portfolio/config"
	"github.com/YungBenn/go-twonana-portfolio/pkg/response"
	"github.com/YungBenn/go-twonana-portfolio/pkg/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type MediaService interface {
	Upload(file string) (*string, *response.Error)
	Delete(file string) (*string, *response.Error)
}

type mediaService struct {
	env config.EnvVars
}

// Delete implements MediaService
func (s *mediaService) Delete(file string) (*string, *response.Error) {
	cld, err := cloudinary.NewFromParams(s.env.CLOUDINARYCLOUDNAME, s.env.CLOUDINARYAPIKEY, s.env.CLOUDINARYSECRET)
	if err != nil {
		return nil, response.NewError(500, err)
	}

	var ctx = context.Background()

	fileName := utils.GetFileName(file)

	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: fileName,
	})

	if err != nil {
		return nil, response.NewError(500, err)
	}

	return &res.Result, nil
}

// Upload implements MediaService
func (s *mediaService) Upload(file string) (*string, *response.Error) {
	cld, err := cloudinary.NewFromParams(s.env.CLOUDINARYCLOUDNAME, s.env.CLOUDINARYAPIKEY, s.env.CLOUDINARYSECRET)
	if err != nil {
		return nil, response.NewError(500, err)
	}

	var ctx = context.Background()

	if file != "" {
		uploadResult, err := cld.Upload.Upload(
			ctx,
			file,
			uploader.UploadParams{
				PublicID: uuid.New().String(),
				Folder:   s.env.CLOUDINARYFOLDER,
			},
		)

		if err != nil {
			return nil, response.NewError(500, err)
		}

		return &uploadResult.SecureURL, nil
	}

	return nil, response.NewError(500, errors.New("cannot read base64 file"))
}

func NewMediaService(env config.EnvVars) MediaService {
	return &mediaService{env}
}
