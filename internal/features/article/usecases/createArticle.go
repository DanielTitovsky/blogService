package article_usecases

import (
	"github.com/google/uuid"
)

type CreateArticleRequest struct {
	Title    string    `json:"title" validate:"required,max=255"`
	AuthorID uuid.UUID `json:"authorID"`
	Content  string    `json:"content" validate:"required"`
}

type CreateArticleService struct {
	AuthorID uuid.UUID
	Title    string
	Content  string
}
