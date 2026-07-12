package main

import (
	"LMSBitLab/internal/api"
	"LMSBitLab/internal/config"
	"LMSBitLab/internal/database"
	"LMSBitLab/internal/handler"
	"LMSBitLab/internal/repository"
	"LMSBitLab/internal/service"

	_ "LMSBitLab/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// @title LMS BitLab API
// @version 1.0
// @description API для управления курсами, главами и уроками платформы LMS BitLab.
// @host localhost:8080
// @BasePath /
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	cfg := config.Load()

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Warnf("Некорректный LOG_LEVEL=%q, используется info: %v", cfg.LogLevel, err)
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
