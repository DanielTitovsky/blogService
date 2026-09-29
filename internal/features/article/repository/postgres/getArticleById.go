package repository_postgres_article

import (
	"blog-service/internal/core/domain"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *ArticleRepository) GetArticleById(ctx context.Context, articleId uuid.UUID) (*domain.Article, error) {

	article := &domain.Article{}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	sql := `
	SELECT
    	id,
    	author_id,
    	title,
    	content,
    	created_at,
    	updated_at
	FROM public.articles
	WHERE id = $1;
	`

	err := r.pool.QueryRow(
		ctx,
		sql,
		articleId,
	).Scan(
		&article.Id,
		&article.AutorId,
		&article.Title,
		&article.Content,
		&article.Created_at,
		&article.Updated_at,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve article: %w", err)
	}

	return article, nil
}
