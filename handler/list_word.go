package handler

import (
	"net/http"

	"github.com/static-fuji/lab_quiz/entity"
)

type ListWord struct {
	Service ListWordsService
}

type word struct {
	ID    entity.WordID `json:"id"`
	Title string        `json:"title"`
}

func (lt *ListWord) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	words, err := lt.Service.ListWords(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := []word{}
	for _, t := range words {
		rsp = append(rsp, word{
			ID:    t.ID,
			Title: t.Title,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
