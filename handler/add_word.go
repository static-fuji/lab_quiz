/*
*専門用語を追加するハンドラ
 */
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type AddWord struct {
	Service   AddWordService
	Validator *validator.Validate
}

func (at *AddWord) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title      string `json:"title" validate:"required"`
		Desc       string `json:"desc" validate:"required"`
		Lab        string `json:"lab" validate:"required"`
		ArticleIDs []int  `json:"article_id" validate:"required"`
	}
	var AddArticleids []int
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	err := validate.StructCtx(ctx, b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	var validArticleIDFound bool
	for _, articleID := range b.ArticleIDs {
		err = at.Service.SearchArticleID(ctx, articleID)
		if err == nil {
			validArticleIDFound = true
			AddArticleids = append(AddArticleids, articleID)
		}
	}

	if !validArticleIDFound {
		RespondJSON(ctx, w, &ErrResponse{
			Message: "invalid article ID",
		}, http.StatusBadRequest)
		return
	}

	t, err := at.Service.AddWord(ctx, b.Title, b.Desc, b.Lab, AddArticleids)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID int `json:"id"`
	}{ID: int(t.ID)}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
