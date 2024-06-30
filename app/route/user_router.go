package route

import (
	"product-rest-api/features/user/handler"
	"product-rest-api/features/user/repository"
	"product-rest-api/features/user/service"
	m "product-rest-api/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitUserRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := repository.New(db)
	userUseCase := service.New(userRepository)
	userController := handler.New(userUseCase)

	e.POST("/user", userController.Register)
	e.POST("/user/login", userController.Login)
	e.GET("/user/:id", userController.ReadSpecificUser, m.JWTMiddleware())
	e.DELETE("/user/:id", userController.DeleteUser, m.JWTMiddleware())
}
