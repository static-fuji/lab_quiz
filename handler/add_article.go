/*
*論文を追加するハンドラ
 */
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AddArticle struct {
	Service   AddArticleService
	Validator *validator.Validate
}

func (aa *AddArticle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	err := aa.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t, err := aa.Service.AddArticle(ctx, b.Title, b.Author)
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
