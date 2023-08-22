package repository

import (
	"context"
	"golang-base64-file-encryption-template/entities"

	"gorm.io/gorm"
)

type (
	ImageRepository interface {
		UploadImage(ctx context.Context, image entities.Image) (entities.Image, error)
		GetImage(ctx context.Context, imageId string) (entities.Image, error)
	}

	imageRepository struct {
		db *gorm.DB
	}
)

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{
		db: db,
	}
}

func (ir *imageRepository) UploadImage(ctx context.Context, image entities.Image) (entities.Image, error) {
	err := ir.db.WithContext(ctx).Create(&image).Error
	if err != nil {
		return entities.Image{}, err
	}

	return image, nil
}

func (ir *imageRepository) GetImage(ctx context.Context, imageId string) (entities.Image, error) {
	var image entities.Image
	err := ir.db.WithContext(ctx).Where("id = ?", imageId).First(&image).Error
	if err != nil {
		return entities.Image{}, err
	}

	return image, nil
}