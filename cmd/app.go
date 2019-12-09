package main

import (
	"achuala.in/payhub/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Configuration
	conf, err := config.GetConfig(os.etenv("BUILD_ENV"), nil)
	if err != nil {
		panic(err) // othing we can do
	}
	if conf.GetString("environment") != "production" {
		conf.Debug()
	}

	// Logger
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level.UnmarshalText([]byte(conf.GetString("log.level")))
	zapConfig.Development = conf.GetString("environment") != "production"
	// Expose log level Prometheus metrics
	hook := promZap.MustNewPrometheusHook([]zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel,
		zapcore.ErrorLevel, zapcore.FatalLevel, zapcore.PanicLevel, zapcore.DebugLevel})
	logger, _ := zapConfig.Build(zap.Hooks(hook))
}
