package httpframework

import (
	"github.com/nehal1992/Go-Clean-Architecture/postService"
	"github.com/nehal1992/Go-Clean-Architecture/repository"
)

type HttpFramework interface {
	HandelRoutes(m *Modules)
}
type Modules struct {
	PostService *postService.PostService
}

func NewHttpRoutesEngine(db repository.PostRepo) *Modules {
	return &Modules{
		PostService: &postService.PostService{
			Engin: db,
		},
	}
}
