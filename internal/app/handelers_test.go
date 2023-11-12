package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	type want struct {
		expectedCode int
		location     string
	}
	tests := []struct {
		name       string
		reqUrl     string
		existedUrl Storage
		want       want
	}{
		{
			name:   "Redirect func test for true",
			reqUrl: "/rsaKyawW",
			existedUrl: Storage{
				links: map[string]string{
					"rsaKyawW": "https://www.e1.ru/",
				},
			},
			want: want{
				expectedCode: 307,
				location:     "https://www.e1.ru/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, tt.reqUrl, nil)
			w := httptest.NewRecorder()
			Redirect(w, r)
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
		contentlength string
	}
	tests := []struct {
		name       string
		reqUrl     string
		body       string
		existedUrl Storage
		want       want
	}{
		{
			name:   "test for CreateRedirect true",
			reqUrl: "/",
			body:   "https://www.e1.ru/",
			existedUrl: Storage{
				links: map[string]string{
					"rsaKyawW": "https://www.e1.ru/",
				},
			},
			want: want{
				expectedCode:  201,
				location:      "https://www.e1.ru/",
				contenttype:   "text/plain",
				contentlength: "30",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, tt.reqUrl, nil)
			w := httptest.NewRecorder()
			Redirect(w, r)
			res := w.Result()
			assert.Equal(t, tt.want.expectedCode, res.StatusCode)
			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
		})
	}
}
