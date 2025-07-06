package main

import (
	"fmt"
	"log"
	"net/http"

	"my-service/config"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("erro ao carregar configuração: %v", err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	e := echo.New()
	e.HideBanner = true

	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	addr := fmt.Sprintf(":%d", config.Cfg.HTTPPort)
	logger.Info("iniciando servidor", zap.String("addr", addr))
	if err := e.Start(addr); err != nil {
		logger.Fatal("erro ao iniciar servidor", zap.Error(err))
	}
}
