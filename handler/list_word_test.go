package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/testutil"
)

func TestListWord(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		words []*entity.Word
		want  want
	}{
		"ok": {
			words: []*entity.Word{
				{
					ID:    1,
					Title: "test1",
				},
				{
					ID:    2,
					Title: "test2",
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_word/ok_rsp.json.golden",
			},
		},
		"empty": {
			words: []*entity.Word{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_word/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/words", nil)

			moq := &ListWordsServiceMock{}
			moq.ListWordsFunc = func(ctx context.Context) (entity.Words, error) {
				if tt.words != nil {
					return tt.words, nil
				}
				return nil, errors.New("error from mock")
			}
			sut := ListWord{Service: moq}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
