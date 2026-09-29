package article_usecases

import (
	"time"

	"github.com/google/uuid"
)

type ArticleRequest struct {
	Id      uuid.UUID `json:"id"`
	AutorId uuid.UUID `json:"autorId"`
	Title   string    `json:"title" validate:"required,max=255"`
	Content string    `json:"content" validate:"required"`
}

type ArticleService struct {
	Id       uuid.UUID
	AuthorID uuid.UUID
	Title    string
	Content  string
}

type ArticleResponse struct {
	ID        uuid.UUID `json:"id"`
	AuthorID  uuid.UUID `json:"authorId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
