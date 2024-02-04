package server

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRedirect(t *testing.T) {
	ts := httptest.NewServer(LinkRouter())
	defer ts.Close()

	type want struct {
		expectedCode int
		location     string
	}
	tests := []struct {
		name    string
		existed string
		want    want
	}{
		{
			name:    "Redirect func test for true",
			existed: "https://www.e1.ru/",
			want: want{
				expectedCode: 307,
				location:     "https://www.e1.ru/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewShortKey, err := storage.Add(tt.existed)
			shorter := "/" + NewShortKey
			if err != nil {
				t.Errorf("Error. Can't add url to storage")
			}
			req, err := http.NewRequest(http.MethodGet, ts.URL+shorter, nil)
			if err != nil {
				t.Errorf("Error. Can't add url to storage")
			}
			resp, err := ts.Client().Do(req)
			if err != nil {
				t.Errorf("Error. Can't add url to storage")
			}
			defer resp.Body.Close()
			assert.Equal(t, tt.want.expectedCode, resp.Request.Response.StatusCode)
			assert.Equal(t, tt.want.location, resp.Request.Response.Header.Get("Location"))
		})
	}
}

func TestCreateRedirect(t *testing.T) {
	ts := httptest.NewServer(LinkRouter())
	defer ts.Close()

	type want struct {
		expectedCode int
		location     string
		proto        string
	}
	tests := []struct {
		name        string
		req         string
		linkForBody string
		want        want
	}{
		{
			name:        "test for CreateRedirect - worked data",
			req:         "/",
			linkForBody: "https://www.e1.ru/",
			want: want{
				expectedCode: 201,
				location:     "https://www.e1.ru/",
				proto:        "HTTP/1.1",
			},
		},
		{
			name:        "test for CreateRedirect - empty data",
			req:         "/",
			linkForBody: "",
			want: want{
				expectedCode: 400,
				location:     "",
				proto:        "HTTP/1.1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, ts.URL+tt.req, strings.NewReader(tt.linkForBody))
			w := httptest.NewRecorder()
			CreateRedirect(w, r)
			res := w.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			shortLink := strings.ReplaceAll(string(body), "http://localhost:8080/", "")
			url1, _ := storage.Get(shortLink)
			assert.Equal(t, tt.want.proto, res.Proto)
			assert.Equal(t, tt.want.location, url1)
			assert.Equal(t, tt.want.expectedCode, res.StatusCode)
		})
	}
}

func TestAPIShorten(t *testing.T) {
	return
}
