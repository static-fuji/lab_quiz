package service

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/store"
)

type WordAdder interface {
	AddWord(ctx context.Context, db store.Execer, t *entity.Word) error
	SearchArticleID(ctx context.Context, db store.Queryer, id int) error
	BindArticleToWords(ctx context.Context, db store.Execer, t *entity.Word) error
}

type WordLister interface {
	ListWords(ctx context.Context, db store.Queryer) (entity.Words, error)
}

type ArticleAdder interface {
	AddArticle(ctx context.Context, db store.Execer, t *entity.Article) error
}

type ArticleLister interface {
	ListArticle(ctx context.Context, db store.Queryer) (entity.Articles, error)
}

type ArticleSearcher interface {
	SearchArticleID(ctx context.Context, db store.Queryer, id int) error
}

type WordSearcher interface {
	ListBind(ctx context.Context, db store.Queryer, id int) (entity.Words, error)
}
