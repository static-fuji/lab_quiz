package handler

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListWordsService AddWordService AddArticleService
type ListWordsService interface {
	ListWords(ctx context.Context) (entity.Words, error)
}

type AddWordService interface {
	AddWord(ctx context.Context, title string, desc string, lab string, articleID []int) (*entity.Word, error)
	SearchArticleID(ctx context.Context, id int) error
}

type ListArticlesService interface {
	ListArticle(ctx context.Context) (entity.Articles, error)
}

type AddArticleService interface {
	AddArticle(ctx context.Context, title string, author string) (*entity.Article, error)
}

type WordSearchService interface {
	ListBind(ctx context.Context, id int) (entity.Words, error)
}
