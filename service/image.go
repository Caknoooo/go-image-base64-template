package service

import (
	"context"
	"errors"
	"golang-base64-file-encryption-template/dto"
	"golang-base64-file-encryption-template/entities"
	"golang-base64-file-encryption-template/utils"

	"github.com/google/uuid"
)

type (
	ImageService interface {
		UploadImage(ctx context.Context, req dto.ImageRequest) (dto.ImageResponse, error)
	}
)

const PATH = "storage"

func UploadImage(ctx context.Context, req dto.ImageRequest) (dto.ImageResponse, error) {
	if req.Image == nil {
		return dto.ImageResponse{}, errors.New("image is required")
	}

	base64encrypted, err := utils.EncodeBase64(req.Image)
	if err != nil {
		return dto.ImageResponse{}, err
	}

	imageId := uuid.New()
	_ = utils.SaveImage(base64encrypted, PATH, imageId.String())
	imageName := utils.GenerateFilename(PATH, imageId.String())

	image := entities.Image{
		ID:       imageId,
		Filename: req.Image.Filename,
		Path:     imageName,
	}

	res := dto.ImageResponse{
		ID:       image.ID.String(),
		Filename: image.Filename,
		Path:     image.Path,
	}

	return res, nil
}