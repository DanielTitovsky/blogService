package repository_postgres_article

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *ArticleRepository) DeleteArcticle(ctx context.Context, articleId uuid.UUID, authorId uuid.UUID) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	sql := `
	DELETE FROM public.articles
	WHERE id = $1
	  AND author_id = $2
	RETURNING id;
	`

	_, err := r.pool.Exec(
		ctx,
		sql,
		articleId,
		authorId,
	)

	if err != nil {
		return fmt.Errorf("Failed to delete article: %w", err)
	}

	return nil
}
