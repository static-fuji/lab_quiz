package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/static-fuji/lab_quiz/entity"
	"github.com/static-fuji/lab_quiz/testutil"
)

func TestAddTask(t *testing.T) {
	t.Parallel()
	type want struct {
		status  int
		rspFile string
	}

	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/add_task/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/add_task/ok_rsp.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/add_task/bad_req_rsp.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/words",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			moq := &AddWordServiceMock{}
			moq.AddWordFunc = func(
				ctx context.Context, title string, desc string, lab string, articleID int,
			) (*entity.Word, error) {
				if tt.want.status == http.StatusOK {
					return &entity.Word{ID: 1}, nil
				}
				return nil, errors.New("error from mock")
			}

			sut := AddWord{
				Service: moq,
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
