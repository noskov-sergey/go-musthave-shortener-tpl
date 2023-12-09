package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRedirect(t *testing.T) {
	type want struct {
		expectedCode int
		location     string
	}
	tests := []struct {
		name       string
		existedUrl string
		want       want
	}{
		{
			name:       "Redirect func test for true",
			existedUrl: "https://www.e1.ru/",
			want: want{
				expectedCode: 307,
				location:     "https://www.e1.ru/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			new_shortkey := storage.Add(tt.existedUrl)
			req := httptest.NewRequest(http.MethodGet, "/"+new_shortkey, nil)
			w := httptest.NewRecorder()
			RouteRedirect(w, req)
			res := w.Result()
			assert.Equal(t, tt.want.expectedCode, res.StatusCode)
			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
		})
	}
}

func TestCreateRedirect(t *testing.T) {
	type want struct {
		expectedCode  int
		location      string
		contenttype   string
		contentlength int64
		proto         string
	}
	tests := []struct {
		name   string
		reqUrl string
		want   want
	}{
		{
			name:   "test for CreateRedirect true",
			reqUrl: "/",
			want: want{
				expectedCode:  201,
				location:      "https://www.e1.ru/",
				contenttype:   "text/plain",
				contentlength: 30,
				proto:         "HTTP/1.1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bodyReader := strings.NewReader("https://www.e1.ru/")
			r := httptest.NewRequest(http.MethodPost, tt.reqUrl, bodyReader)
			w := httptest.NewRecorder()
			RouteRedirect(w, r)
			res := w.Result()
			assert.Equal(t, tt.want.proto, res.Proto)
			assert.Equal(t, tt.want.expectedCode, res.StatusCode)
		})
	}
}
