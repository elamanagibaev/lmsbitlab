package main

import (
	"LMSBitLab/internal/config"
	"LMSBitLab/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	errStart := router.Run(":" + cfg.AppPort)
	if errStart != nil {
		log.Fatalf("Не удалось запустить приложение: %v", errStart)
	}
}
