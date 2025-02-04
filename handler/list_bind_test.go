package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/testutil"
)

func TestListBind(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		articleID int
		words     []*entity.Word
		want      want
	}{
		"ok": {
			articleID: 1,
			words: []*entity.Word{
				{
					ID:    1,
					Title: "word1",
					Desc:  "description1",
					Lab:   "lab1",
				},
				{
					ID:    2,
					Title: "word2",
					Desc:  "description2",
					Lab:   "lab2",
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_bind/ok_rsp.json.golden",
			},
		},
		"empty": {
			articleID: 2,
			words:     []*entity.Word{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_bind/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			moq := &WordSearchServiceMock{}
			moq.ListBindFunc = func(ctx context.Context, articleID int) (entity.Words, error) {
				return tt.words, nil
			}

			sut := ListBind{
				Service: moq,
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/articles/{articleID}/words", nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("articleID", strconv.Itoa(tt.articleID))
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
		})
	}
}
