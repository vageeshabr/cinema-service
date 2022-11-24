package stores

import (
	"context"
	"github.com/vageeshabr/cinema-service/internal/models"
)

type MovieStorer interface {
	Get(ctx context.Context, pageNo int) ([]*models.Movie, error)
	Count(ctx context.Context) (int, error)
}
