package service

import (
	"context"
	"fmt"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type AddWord struct {
	ExeceDB store.Execer
	QueryDB store.Queryer
	Repo    WordAdder
}

func (a *AddWord) AddWord(ctx context.Context, title string, desc string, lab string, articleIDs []int) (*entity.Word, error) {
	t := &entity.Word{
		Title:      title,
		Desc:       desc,
		Lab:        lab,
		ArticleIDs: articleIDs,
	}

	err := a.Repo.AddWord(ctx, a.ExeceDB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	err = a.Repo.BindArticleToWords(ctx, a.ExeceDB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to bind article to words: %w", err)
	}

	return t, nil
}

func (a *AddWord) SearchArticleID(ctx context.Context, id int) error {
	err := a.Repo.SearchArticleID(ctx, a.QueryDB, id)
	if err != nil {
		return fmt.Errorf("failed to get article id: %w", err)
	}

	return nil
}
