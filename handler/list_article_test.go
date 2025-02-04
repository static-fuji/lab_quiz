package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/testutil"
)

func TestListArticle(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		articles []*entity.Article
		want     want
	}{
		"ok": {
			articles: []*entity.Article{
				{
					ID:     1,
					Title:  "test1",
					Author: "author1",
				},
				{
					ID:     2,
					Title:  "test2",
					Author: "author2",
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_article/ok_rsp.json.golden",
			},
		},
		"empty": {
			articles: []*entity.Article{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_article/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			moq := &ListArticlesServiceMock{}
			moq.ListArticlesFunc = func(ctx context.Context) (entity.Articles, error) {
				return tt.articles, nil
			}

			sut := ListArticle{
				Service: moq,
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/articles", nil)

			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
		})
	}
}
