package article_services

import (
	"blog-service/internal/core/domain"
	article_usecases "blog-service/internal/features/article/usecases"
	"context"
	"fmt"
)

func (s *ArticleService) CreateArticle(ctx context.Context, article article_usecases.CreateArticleService) (*domain.Article, error) {

	domainArticle := domain.Article{
		AutorId: article.AuthorID,
		Content: article.Content,
		Title:   article.Title,
	}

	createdArticle, err := s.repository.CreateArticle(ctx, domainArticle)

	if err != nil {
		return nil, fmt.Errorf("Failed to create article: %w", err)
	}

	return createdArticle, nil
}
