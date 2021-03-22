package repository

import (
	"context"

	"github.com/nehal1992/Go-Clean-Architecture/models"
)

type PostRepo interface {
	List(ctx context.Context) ([]models.Post, error)
	Create(context.Context, models.Post) (interface{}, error)
	Get(ctx context.Context, id int) (models.Post, error)
	Update(ctx context.Context, id int, p models.Post) error
	Delete(ctx context.Context, id int) (bool, error)
}
