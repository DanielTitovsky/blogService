package article_usecases

import "github.com/google/uuid"

type ArticleId struct {
	ArticleId uuid.UUID `json:"id" validate:"required"`
}
