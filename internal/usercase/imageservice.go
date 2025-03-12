package usercase

import (
	"context"
	"errors"
	"go-gin-image-store/internal/domain/models"
	"go-gin-image-store/internal/repository"
	"io"
)

type ImageService struct {
	repo repository.ImageRepository
}

func NewImageService(repo repository.ImageRepository) *ImageService {
	return &ImageService{repo: repo}
}

func (s *ImageService) UploadImage(ctx context.Context, filename string, content io.Reader) (*models.Image, error) {

	// Verificar entradas
	if content == nil {
		return nil, errors.New("content cannot be nil")
	}

	// Llamar al repositorio
	image, err := s.repo.Store(ctx, filename, content)

	if err != nil {
		return nil, err
	}

	// Nunca devuelvas nil sin error
	if image == nil {
		return nil, errors.New("repository returned nil image without error")
	}

	return image, nil
}

func (s *ImageService) GetImage(ctx context.Context, id string) (*models.Image, io.Reader, error) {
	return s.repo.FindByID(ctx, id)
}
