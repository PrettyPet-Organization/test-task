package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KozlovNikolai/test-task/internal/handlers"
	"github.com/KozlovNikolai/test-task/internal/middlewares"
	"github.com/KozlovNikolai/test-task/internal/pkg/config"
	"github.com/KozlovNikolai/test-task/internal/store"
	"github.com/KozlovNikolai/test-task/internal/store/inmemory"
	"github.com/KozlovNikolai/test-task/internal/store/pgstore"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// Server is ...
type Server struct {
	router *gin.Engine
	logger *zap.Logger
}

// NewServer is ...
func NewServer() *Server {
	// Инициализация логгера Zap
	//	logger, err := zap.NewProduction()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	var repoWR store.IUserRepository
	var repoRO store.IUserRepository

	// Выбор репозитория
	switch config.Cfg.RepoType {
	case "postgres":
		poolWR, err := pgxpool.Connect(context.Background(), config.Cfg.StorageWR)
		if err != nil {
			logger.Fatal("Unable to connect to database master", zap.Error(err))
		}
		repoWR = pgstore.NewUserRepository(poolWR, logger)

		poolRO, err := pgxpool.Connect(context.Background(), config.Cfg.StorageRO)
		if err != nil {
			logger.Fatal("Unable to connect to database replica", zap.Error(err))
		}
		repoRO = pgstore.NewUserRepository(poolRO, logger)
	case "inmemory":
		repoWR = inmemory.NewUserRepository(logger)
		repoRO = repoWR
	default:
		logger.Fatal("Invalid repository type")
	}

	// Создание сервера
	server := &Server{
		router: gin.Default(),
		logger: logger,
	}

	// Middleware
	server.router.Use(middlewares.LoggerMiddleware(logger))
	server.router.Use(middlewares.RequestIDMiddleware())

	// CORS
	server.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:8443", "https://127.0.0.1:8443"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Инициализация обработчиков
	userHandler := handlers.NewUserHandler(logger, repoWR, repoRO)

	// add swagger
	server.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CRUD маршруты для User
	openRouter := server.router.Group("/user")
	openRouter.GET("/:id", userHandler.GetUser)
	openRouter.POST("/register", userHandler.CreateUser)
	openRouter.POST("/login", userHandler.LoginUser)

	authorized := server.router.Group("/")
	authorized.Use(middlewares.AuthMiddleware())

	authorized.GET("/users", userHandler.GetUsers)
	// openRouter.GET("/:id", orderHandler.GetOrderByID)
	// openRouter.GET("/list", orderHandler.GetAllOrdersList)
	// openRouter.GET("/bypatient/:patient_id/:is_active", orderHandler.GetOrdersByPatientID)
	// openRouter.PUT("/:id", orderHandler.UpdateOrder)
	// openRouter.PATCH("/addservices", orderHandler.AddServicesToOrder)
	// openRouter.DELETE("/:id", orderHandler.DeleteOrder)

	// CRUD маршруты для Orders
	// baseRouter := server.router.Group("/orders")
	// baseRouter.POST("", orderHandler.CreateOrder)
	// baseRouter.GET("/:id", orderHandler.GetOrderByID)
	// baseRouter.GET("/list", orderHandler.GetAllOrdersList)
	// baseRouter.GET("/bypatient/:patient_id/:is_active", orderHandler.GetOrdersByPatientID)
	// baseRouter.PUT("/:id", orderHandler.UpdateOrder)
	// baseRouter.PATCH("/addservices", orderHandler.AddServicesToOrder)
	// baseRouter.DELETE("/:id", orderHandler.DeleteOrder)
	return server
}

// Run is ...
func (s *Server) Run() {
	defer func() {
		_ = s.logger.Sync() // flushes buffer, if any
	}()
	// Настройка сервера с таймаутами
	server := &http.Server{
		Addr:         config.Cfg.Address,
		Handler:      s.router,
		ReadTimeout:  config.Cfg.Timeout,
		WriteTimeout: config.Cfg.Timeout,
		IdleTimeout:  config.Cfg.IdleTimout,
	}

	if err := server.ListenAndServeTLS(config.CertFile, config.KeyFile); err != nil && err != http.ErrServerClosed {
		s.logger.Fatal(fmt.Sprintf("Could not listen on %s", config.Cfg.Address), zap.Error(err))
	}
}
