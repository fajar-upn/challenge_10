package router

import (
	"challenge_10/handler"
	"challenge_10/middleware"
	"challenge_10/repositories"
	"challenge_10/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	bookRepository := repositories.NewRepository(db)
	bookService := service.NewService(bookRepository)
	bookHandler := handler.NewHandler(bookService)

	userRepository := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// route book product
	book := api.Group("/book")
	book.Use(middleware.Authentication())
	book.POST("/create", bookHandler.CreateBook)
	book.GET("/", bookHandler.GetBooks)
	book.GET("/:bookID", bookHandler.GetBookByID)
	book.PUT("/:bookID", bookHandler.UpdateBookByID)
	book.DELETE("/:bookID", bookHandler.DeleteBookByID)

	// route user admin/ client
	user := api.Group("/user")
	user.POST("/client", userHandler.CreateClientUser)
	user.POST("/admin", userHandler.CreateAdminUser)
	user.POST("/validate", userHandler.ValidateUser)

	return router
}
