package article_services

import (
	"blog-service/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

type ArticleService struct {
	repository ArticleRepository
}

type ArticleRepository interface {
	CreateArticle(ctx context.Context, article domain.Article) (*domain.Article, error)
	DeleteArcticle(ctx context.Context, articleId uuid.UUID, authorId uuid.UUID) error
	GetArticleById(ctx context.Context, articleId uuid.UUID) (*domain.Article, error)
	UpdateArticle(ctx context.Context, article domain.Article) (*domain.Article, error)
}

func NewArticleService(repository ArticleRepository) ArticleService {
	return ArticleService{
		repository: repository,
	}
}
