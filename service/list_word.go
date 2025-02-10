/*
*専門用語一覧を取得するサービス
 */
package service

import (
	"context"
	"fmt"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type ListWord struct {
	DB   store.Queryer
	Repo WordLister
}

func (l *ListWord) ListWords(ctx context.Context) (entity.Words, error) {
	ts, err := l.Repo.ListWords(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
