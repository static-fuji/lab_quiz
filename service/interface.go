package service

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type WordAdder interface {
	AddWord(ctx context.Context, db store.Execer, t *entity.Word) error
}

type WordLister interface {
	ListWords(ctx context.Context, db store.Queryer) (entity.Words, error)
}
