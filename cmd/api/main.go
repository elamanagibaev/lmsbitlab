package main

import (
	"LMSBitLab/internal/api"
	"LMSBitLab/internal/config"
	"LMSBitLab/internal/database"
	"LMSBitLab/internal/handler"
	"LMSBitLab/internal/repository"
	"LMSBitLab/internal/service"

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

	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	chapterRepo := repository.NewChapterRepository(db)
	chapterService := service.NewChapterService(chapterRepo)
	chapterHandler := handler.NewChapterHandler(chapterService)

	lessonRepo := repository.NewLessonRepository(db)
	lessonService := service.NewLessonService(lessonRepo)
	lessonHandler := handler.NewLessonHandler(lessonService)

	router := gin.Default()
	router.Use(api.ErrorMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	courses := router.Group("/courses")
	{
		courses.POST("", courseHandler.Create)
		courses.GET("", courseHandler.GetAll)
		courses.GET("/:id", courseHandler.GetByID)
		courses.PUT("/:id", courseHandler.Update)
		courses.DELETE("/:id", courseHandler.Delete)
		courses.GET("/:id/chapters", chapterHandler.GetByCourseID)
	}

	chapters := router.Group("/chapters")
	{
		chapters.POST("", chapterHandler.Create)
		chapters.GET("/:id", chapterHandler.GetByID)
		chapters.PUT("/:id", chapterHandler.Update)
		chapters.DELETE("/:id", chapterHandler.Delete)
		chapters.GET("/:id/lessons", lessonHandler.GetByChapterID)
	}

	lessons := router.Group("/lessons")
	{
		lessons.POST("", lessonHandler.Create)
		lessons.GET("/:id", lessonHandler.GetByID)
		lessons.PUT("/:id", lessonHandler.Update)
		lessons.DELETE("/:id", lessonHandler.Delete)
	}

	errStart := router.Run(":" + cfg.AppPort)
	if errStart != nil {
		logrus.Fatalf("Не удалось запустить приложение: %v", errStart)
	}
}
