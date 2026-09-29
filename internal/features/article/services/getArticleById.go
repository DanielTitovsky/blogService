package article_services

import (
	"blog-service/internal/core/domain"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *ArticleService) GetArticleById(ctx context.Context, articleId uuid.UUID) (*domain.Article, error) {

	article, err := s.repository.GetArticleById(ctx, articleId)

	if err != nil {
		return nil, fmt.Errorf("Failed to get article by id: %w", err)
	}

	return article, nil
}
