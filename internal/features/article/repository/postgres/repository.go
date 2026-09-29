package repository_postgres_article

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ArticleRepository struct {
	pool *pgxpool.Pool
}

func NewArticleRepository(pool *pgxpool.Pool) ArticleRepository {
	return ArticleRepository{
		pool: pool,
	}
}
