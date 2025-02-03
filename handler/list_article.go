package handler

import (
	"net/http"

	"github.com/static-fuji/lab_quiz/entity"
)

type ListArticle struct {
	Service ListArticlesService
}

type article struct {
	ID     entity.ArticleID `json:"id"`
	Title  string           `json:"title"`
	Author string           `json:"author"`
}

func (la *ListArticle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	articles, err := la.Service.ListArticle(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := []article{}
	for _, t := range articles {
		rsp = append(rsp, article{
			ID:     t.ID,
			Title:  t.Title,
			Author: t.Author,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
