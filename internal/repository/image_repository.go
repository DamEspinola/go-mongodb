package repository

import (
	"context"
	"go-gin-image-store/internal/domain/models"
	"io"
)

type ImageRepository interface {
	Store(ctx context.Context, filename string, content io.Reader) (*models.Image, error)
	FindByID(ctx context.Context, id string) (*models.Image, io.Reader, error)
}
