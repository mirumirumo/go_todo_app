package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/mirumirumo/go_todo_app/entity"
	"github.com/mirumirumo/go_todo_app/store"
	"github.com/mirumirumo/go_todo_app/testutil"
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
	}{"ok": {reqFile: "a", want: want{status: http.StatusOK, rspFile: "a"}},
		"badRequest": {reqFile: "a", want: want{status: http.StatusBadRequest, rspFile: "a"}}}
	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)
			sut := AddTask{Store: &store.TaskStore{
				Tasks: map[entity.TaskID]*entity.Task{}},
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
		})

	}
}
