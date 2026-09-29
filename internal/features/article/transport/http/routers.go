package article_http_router

import (
	"blog-service/internal/core/domain"
	core_transport_http_server "blog-service/internal/core/transport/http/server"
	article_usecases "blog-service/internal/features/article/usecases"
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, article article_usecases.CreateArticleService) (*domain.Article, error)
	DeleteArcticle(ctx context.Context, articleId uuid.UUID, authorId uuid.UUID) error
	GetArticleById(ctx context.Context, articleId uuid.UUID) (*domain.Article, error)
	UpdateArticle(ctx context.Context, article article_usecases.ArticleService) (*domain.Article, error)
}

type ArticleRouter struct {
	services  ArticleService
	validator *validator.Validate
}

func NewArticleRouters(services ArticleService, validator *validator.Validate) ArticleRouter {
	return ArticleRouter{
		services:  services,
		validator: validator,
	}
}

func (r *ArticleRouter) Routers() []core_transport_http_server.Route {
	return []core_transport_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "article",
			Handler: r.CreateArticle,
		},
		{
			Method:  http.MethodGet,
			Path:    "article",
			Handler: r.GetArticleById,
		},
		{
			Method:  http.MethodPut,
			Path:    "article",
			Handler: r.UpdateArticle,
		},
		{
			Method:  http.MethodDelete,
			Path:    "article",
			Handler: r.DeleteArticle,
		},
	}
}
