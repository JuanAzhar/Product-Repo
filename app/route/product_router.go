package route

import (
	"product-rest-api/features/product/handler"
	"product-rest-api/features/product/repository"
	"product-rest-api/features/product/service"
	m "product-rest-api/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitProductRouter(db *gorm.DB, e *echo.Echo) {
	productRepository := repository.New(db)
	productUseCase := service.New(productRepository)
	productController := handler.New(productUseCase)

	e.POST("/product", productController.PostProduct, m.JWTMiddleware())
	e.GET("/product/:id", productController.ReadSpecificProduct, m.JWTMiddleware())
	e.GET("/product", productController.ReadAllProduct)
	e.DELETE("/product/:id", productController.DeleteProduct, m.JWTMiddleware())
	e.PUT("/product/:id", productController.UpdateProduct, m.JWTMiddleware())
}
