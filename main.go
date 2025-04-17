package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-contrib/cors"
	db "github.com/troodinc/trood-front-hackathon/database"
	_ "github.com/troodinc/trood-front-hackathon/docs"
	"github.com/troodinc/trood-front-hackathon/handlers"
)

// @title Trood Front Hackathon API
// @version 1.0
// @description This is the API documentation for the Trood Front Hackathon. Welcome to hell.
// @host localhost:8080
// @BasePath /

func main() {
	// Инициализация базы данных и проектов
	db.InitDatabase()
	handlers.InitProjects()

	// Создаем экземпляр Gin
	r := gin.Default()

	// Настройка CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешаем все источники
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Разрешаем только необходимые методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешаем необходимые заголовки
		AllowCredentials: true, // Разрешаем отправку cookie
	}))

	// Роуты для Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Роуты для проектов
	r.GET("/projects", handlers.GetProjects)
	r.GET("/projects/:id", handlers.GetProjectByID)
	r.POST("/projects", handlers.CreateProject)
	r.PUT("/projects/:id", handlers.EditProject)
	r.DELETE("/projects/:id", handlers.DeleteProject)

	// Роуты для вакансий
	r.GET("/projects/:id/vacancies", handlers.GetVacancies)
	r.POST("/projects/:id/vacancies", handlers.CreateVacancy)
	r.PUT("/vacancies/:id", handlers.EditVacancy)
	r.DELETE("/vacancies/:id", handlers.DeleteVacancy)

	// Запуск сервера
	port := "8080"
	log.Println("Server running on http://localhost:" + port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
