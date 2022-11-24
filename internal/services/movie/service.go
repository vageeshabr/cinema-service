package movie

import (
	"context"
	"fmt"
	"github.com/vageeshabr/cinema-service/internal/models"
	"github.com/vageeshabr/cinema-service/internal/services"
	"github.com/vageeshabr/cinema-service/internal/stores"
)

type Service struct {
	store stores.MovieStorer
}

type InvalidParam struct {
	Param string
}

func (e InvalidParam) Error() string {
	return fmt.Sprintf("param `%s` is invalid", e.Param)
}

func (s *Service) Get(ctx context.Context, pageNo int) ([]*models.Movie, error) {
	if pageNo <= 0 {
		return nil, InvalidParam{Param: "page"}
	}

	return s.store.Get(ctx, pageNo)
}

func (s *Service) Count(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

func New(store stores.MovieStorer) services.Movie {
	return &Service{store: store}
}
