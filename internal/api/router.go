package api

import (
	"rbp/internal/api/handlers"
	"rbp/internal/api/middleware"
	"rbp/internal/db/sqlc"
	"rbp/internal/repository"
	"rbp/internal/service"
	"rbp/pkg/cache"

	"github.com/gin-gonic/gin"
)

func NewRouter(db sqlc.DBTX, redisClient *cache.RedisClient) *gin.Engine {
	router := gin.Default()

	querier := sqlc.New(db)

	userRepo := repository.NewUserRepository(querier)
	productRepo := repository.NewProductRepository(querier)
	orderRepo := repository.NewOrderRepository(querier)
	archiveRepo := repository.NewArchiveRepository(querier)

	authService := service.NewAuthService(userRepo, redisClient)
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo)
	archiveService := service.NewArchiveService(archiveRepo)

	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)
	archiveHandler := handlers.NewArchiveHandler(archiveService)

	router.Use(middleware.RateLimiter(redisClient.Client))

	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(authService))
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProductByID)

		api.POST("/orders", orderHandler.CreateOrder)
		api.GET("/orders", orderHandler.GetOrders)

		api.POST("/archives", archiveHandler.AddToArchive)
		api.GET("/archives", archiveHandler.GetArchivesByUserID)
	}

	return router
}
