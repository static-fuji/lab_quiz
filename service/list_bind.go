package service

import (
	"context"
	"fmt"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type ListBind struct {
	DB   store.Queryer
	Repo WordSearcher
}

func (l *ListBind) ListBind(ctx context.Context, id int) (entity.Words, error) {
	ts, err := l.Repo.ListBind(ctx, l.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
