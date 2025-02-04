package service

import (
	"context"
	"fmt"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type ListArticle struct {
	DB   store.Queryer
	Repo ArticleLister
}

func (l *ListArticle) ListArticles(ctx context.Context) (entity.Articles, error) {
	ts, err := l.Repo.ListArticle(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
