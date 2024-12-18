package handler

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListWordsService AddWordService
type ListWordsService interface {
	ListWords(ctx context.Context) (entity.Words, error)
}

type AddWordService interface {
	AddWord(ctx context.Context, title string) (*entity.Word, error)
}
