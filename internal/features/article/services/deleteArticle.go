package article_services

import (
	"context"

	"github.com/google/uuid"
)

func (s *ArticleService) DeleteArcticle(ctx context.Context, articleId uuid.UUID, authorId uuid.UUID) error {

	err := s.repository.DeleteArcticle(ctx, articleId, authorId)

	return err
}
