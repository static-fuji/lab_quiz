package service

import (
	"context"
	"fmt"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type AddWord struct {
	DB   store.Execer
	Repo WordAdder
}

func (a *AddWord) AddWord(ctx context.Context, title string, desc string) (*entity.Word, error) {
	t := &entity.Word{
		Title: title,
		Desc:  desc,
	}

	err := a.Repo.AddWord(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return t, nil
}
