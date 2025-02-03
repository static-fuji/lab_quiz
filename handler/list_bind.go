package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/static-fuji/lab_quiz/entity"
)

type ListBind struct {
	Service WordSearchService
}

type word_article struct {
	ID    entity.WordID `json:"id"`
	Title string        `json:"title"`
	Desc  string        `json:"desc"`
	Lab   string        `json:"lab"`
}

func (lb *ListBind) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	articleIDStr := chi.URLParam(r, "articleID")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: "invalid article ID",
		}, http.StatusBadRequest)
		return
	}

	words, err := lb.Service.ListBind(ctx, articleID)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := []word_article{}
	for _, t := range words {
		rsp = append(rsp, word_article{
			ID:    t.ID,
			Title: t.Title,
			Desc:  t.Desc,
			Lab:   t.Lab,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
