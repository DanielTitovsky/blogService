package repository_postgres_article

import (
	"blog-service/internal/core/domain"
	"context"
	"fmt"
	"time"
)

func (r *ArticleRepository) CreateArticle(ctx context.Context, article domain.Article) (*domain.Article, error) {
	createdArticle := &domain.Article{}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	sql := `
	INSERT INTO public.articles (
	    author_id,
	    title,
	    content
	)
	VALUES ($1, $2, $3)
	RETURNING
    	id,
    	author_id,
    	title,
    	content,
    	created_at,
    	updated_at;
	`

	err := r.pool.QueryRow(
		ctx,
		sql,
		article.AutorId,
		article.Title,
		article.Content,
	).Scan(
		&createdArticle.Id,
		&createdArticle.AutorId,
		&createdArticle.Title,
		&createdArticle.Content,
		&createdArticle.Created_at,
		&createdArticle.Updated_at,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to created article: %W", err)
	}

	return createdArticle, nil
}
