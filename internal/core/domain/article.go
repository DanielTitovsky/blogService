package domain

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	Id         uuid.UUID
	AutorId    uuid.UUID
	Title      string
	Content    string
	Created_at time.Time
	Updated_at time.Time
}
