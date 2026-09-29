package repository_postgres_article

import (
	"blog-service/internal/core/domain"
	"context"
	"fmt"
	"time"
)

func (r *ArticleRepository) UpdateArticle(ctx context.Context, article domain.Article) (*domain.Article, error) {
	updatedActicle := &domain.Article{}

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	sql := `
	UPDATE public.articles
	SET
	    title = $3,
	    content = $4,
	    updated_at = now()
	WHERE id = $1
	  AND author_id = $2
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
		article.Id,
		article.AutorId,
		article.Title,
		article.Content,
	).Scan(
		&updatedActicle.Id,
		&updatedActicle.AutorId,
		&updatedActicle.Title,
		&updatedActicle.Content,
		&updatedActicle.Created_at,
		&updatedActicle.Updated_at,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to update article: %w", err)
	}

	return updatedActicle, nil
}
