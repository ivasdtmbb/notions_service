package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"notions_service/app/cmd/internal/app"
	"notions_service/app/cmd/internal/config"
)

func main() {
	cfg := config.GetConfig()

	logger := setupLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Running Application")
	a.Run()

}

func setupLogger(env string) *logrus.Logger {
	log := logrus.New()

	switch env {
	case config.ENV_LOCAL:
		setupPrettyLogrus(log)
	case config.ENV_DEV:
		// JSON + Debug
		setupDevLogrus(log)
	case config.ENV_PROD:
		// JSON + info
		setupProdLogrus(log)
	}

	return log
}

func setupPrettyLogrus(log *logrus.Logger) {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
}

func setupDevLogrus(log *logrus.Logger) {
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
}

func setupProdLogrus(log *logrus.Logger) {
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.InfoLevel)
	log.SetOutput(os.Stdout)
}
