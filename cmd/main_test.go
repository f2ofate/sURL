package main

import (
	"net/http"
	"net/http/httptest"
	v1 "sURL/internal/controller/http/v1"
	"sURL/internal/storage/memory"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = memory.MemStorage{}

func TestStoreURL(t *testing.T) {
	type want struct {
		code int
		body string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Empty body",
			want: want{
				code: 400,
				body: "",
			},
		},
		{
			name: "Body with URL",
			want: want{
				code: 201,
				body: "https://google.com",
			},
		},
		{
			name: "Body with wrong URL",
			want: want{
				code: 400,
				body: "htt://google.com",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(test.want.body))
			w := httptest.NewRecorder()

			handler := v1.StoreURL(&s)
			handler(w, req)

			res := w.Result()
			assert.Equal(t, res.StatusCode, test.want.code)

			defer res.Body.Close()
		})
	}
}
