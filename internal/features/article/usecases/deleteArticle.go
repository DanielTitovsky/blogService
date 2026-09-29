package article_usecases

import (
	"github.com/google/uuid"
)

type DeleteArticleRequest struct {
	Id       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"authorID"`
}
