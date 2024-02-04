package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Sugar zap.SugaredLogger

func Initialize() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic(err)
	}
	defer logger.Sync()

	Sugar = *logger.Sugar()
	return nil
}

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			responseData := &responseData{
				status: 0,
				size:   0,
			}

			lw := loggingResponseWriter{
				ResponseWriter: w,
				responseData:   responseData,
			}

			next.ServeHTTP(&lw, r)

			Sugar.Infoln(
				"uri", r.RequestURI,
				"method", r.Method,
				"status", responseData.status,
				"duration", time.Since(start),
				"size", responseData.size,
			)

		}
		return http.HandlerFunc(fn)
	}
}
