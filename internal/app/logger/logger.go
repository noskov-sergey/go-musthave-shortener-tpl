package logger

import (
	"go.uber.org/zap"
)

var Sugar zap.SugaredLogger

func Initialize() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic(err)
	}
	defer logger.Sync()

	// делаем регистратор SugaredLogger
	Sugar = *logger.Sugar()
	return nil
}
