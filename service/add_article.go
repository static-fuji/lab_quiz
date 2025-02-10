/*
*論文を追加するサービス
 */
package service

import (
	"context"
	"fmt"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type AddArticle struct {
	DB   store.Execer
	Repo ArticleAdder
}

func (a *AddArticle) AddArticle(ctx context.Context, title string, author string) (*entity.Article, error) {
	t := &entity.Article{
		Title:  title,
		Author: author,
	}

	err := a.Repo.AddArticle(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return t, nil
}
