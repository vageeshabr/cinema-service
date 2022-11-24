package movie

import (
	"context"
	"database/sql"
	"github.com/vageeshabr/cinema-service/internal/models"
)

type Store struct {
	db *sql.DB
}

func (s *Store) Count(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Get(ctx context.Context, pageNo int) ([]*models.Movie, error) {

	//s.db.QueryContext(ctx, "select id, name, ratings from movies limit 20 offset ")
	panic("not implemented")
}
