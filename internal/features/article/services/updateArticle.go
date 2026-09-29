package article_services

import (
	"blog-service/internal/core/domain"
	article_usecases "blog-service/internal/features/article/usecases"
	"context"
	"fmt"
)

func (s *ArticleService) UpdateArticle(ctx context.Context, article article_usecases.ArticleService) (*domain.Article, error) {

	domainArticle := &domain.Article{
		Id:      article.Id,
		AutorId: article.AuthorID,
		Title:   article.Title,
		Content: article.Content,
	}

	domainArticle, err := s.repository.UpdateArticle(ctx, *domainArticle)

	if err != nil {
		return nil, fmt.Errorf("Failed to updated article: %w", err)
	}

	return domainArticle, nil
}
