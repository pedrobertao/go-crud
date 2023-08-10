package logging

import "go.uber.org/zap"

var L *zap.SugaredLogger
var Logger *zap.Logger

func Start() error {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	L = Logger.Sugar()
	return nil
}
