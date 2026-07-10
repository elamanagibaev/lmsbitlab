package main

import (
	"LMSBitLab/internal/api"
	"LMSBitLab/internal/config"
	"LMSBitLab/internal/database"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	cfg := config.Load()

	db, err := database.NewConnection(cfg)
	if err != nil {
		logrus.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := gin.Default()
	router.Use(api.ErrorMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	errStart := router.Run(":" + cfg.AppPort)
	if errStart != nil {
		logrus.Fatalf("Не удалось запустить приложение: %v", errStart)
	}
}
