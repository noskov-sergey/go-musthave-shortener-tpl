package storage

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"go-musthave-shortener-tpl/internal/app/backup"
	"go-musthave-shortener-tpl/internal/app/logger"
	"go-musthave-shortener-tpl/internal/app/models"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRedirect(t *testing.T) {
	logger.Initialize()
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
			NewShortKey, err := storage.RealStorage.Add(tt.existed)
			shorter := "/" + NewShortKey
			if err != nil {
				t.Errorf("Error. Can't get short uri from storage")
			}
			req, err := http.NewRequest(http.MethodGet, ts.URL+shorter, nil)
			if err != nil {
				t.Errorf("Error. Can't make request")
			}
			resp, err := ts.Client().Do(req)
			if err != nil {
				t.Errorf("Error. After trying to get response %s", err)
			}
			assert.Equal(t, tt.want.expectedCode, resp.Request.Response.StatusCode)
			assert.Equal(t, tt.want.location, resp.Request.Response.Header.Get("Location"))
			resp.Body.Close()
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
			url1, _ := storage.RealStorage.Get(shortLink)
			assert.Equal(t, tt.want.proto, res.Proto)
			assert.Equal(t, tt.want.location, url1)
			assert.Equal(t, tt.want.expectedCode, res.StatusCode)
		})
	}
}

func TestAPIShorten(t *testing.T) {
	logger.Initialize()
	ts := httptest.NewServer(LinkRouter())
	defer ts.Close()

	var resAPI models.ResponseShorten

	exampleBody := `{
  		"url": "https://practicum.yandex.ru"
	}`
	exampleLink := "https://practicum.yandex.ru"

	testCases := []struct {
		name         string // добавляем название тестов
		method       string
		body         string // добавляем тело запроса в табличные тесты
		expectedCode int
		expectedBody string
	}{
		{
			name:         "method_get",
			method:       http.MethodGet,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "method_put",
			method:       http.MethodPut,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "method_delete",
			method:       http.MethodDelete,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "method_post_without_body",
			method:       http.MethodPost,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tc.method
			req.URL = ts.URL + "/api/shorten"

			if len(tc.body) > 0 {
				req.SetHeader("Content-Type", "application/json")
				req.SetBody(tc.body)
			}

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")
			assert.Equal(t, tc.expectedCode, resp.StatusCode(), "Response code didn't match expected")
		})
	}
	t.Run("Post method full check", func(t *testing.T) {
		req := resty.New().R()
		req.Method = http.MethodPost
		req.URL = ts.URL + "/api/shorten"
		req.SetHeader("Content-Type", "application/json")
		req.SetBody(exampleBody)
		resp, err := req.Send()

		assert.NoError(t, err, "error making HTTP request")
		assert.Equal(t, http.StatusCreated, resp.StatusCode(), "Response code didn't match expected")

		json.Unmarshal(resp.Body(), &resAPI)

		reqnext := resty.New().R()
		reqnext.Method = http.MethodGet
		reqnext.URL = ts.URL + resAPI.Result[21:]
		respAPI, err := reqnext.Send()

		uri, _ := storage.RealStorage.Get(resAPI.Result[22:])

		assert.Equal(t, uri, exampleLink, "error with link")
		assert.NoError(t, err, "error making HTTP request")
		assert.Equal(t, http.StatusTemporaryRedirect, respAPI.RawResponse.Request.Response.Request.Response.StatusCode, "Response code didn't match expected")
	})
}

func TestPingAPI(t *testing.T) {}
