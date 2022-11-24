package services

import (
	"context"
	"github.com/vageeshabr/cinema-service/internal/models"
)

type Movie interface {
	Get(ctx context.Context, pageNo int) ([]*models.Movie, error)
	Count(ctx context.Context) (int, error)
}
