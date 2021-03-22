package postService

import (
	"context"
	"math/rand"

	"github.com/nehal1992/Go-Clean-Architecture/models"
	"github.com/nehal1992/Go-Clean-Architecture/repository"
)

type PostService struct {
	Engin repository.PostRepo
}

func (p *PostService) List(ctx context.Context) ([]models.Post, error) {
	return p.Engin.List(ctx)
}
func (p *PostService) Get(ctx context.Context, id int) (models.Post, error) {
	return p.Engin.Get(ctx, id)
}
func (p *PostService) Add(ctx context.Context, post models.Post) error {
	post.ID = rand.Int()
	_, err := p.Engin.Create(ctx, post)
	return err
}

func (p *PostService) Update(ctx context.Context, id int, post models.Post) error {
	return p.Engin.Update(ctx, id, post)
}

func (p *PostService) Delete(ctx context.Context, id int) (bool, error) {
	return p.Engin.Delete(ctx, id)
}
