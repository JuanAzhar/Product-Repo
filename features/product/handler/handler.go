package handler

import (
	"net/http"
	"product-rest-api/features/product/entity"
	middlewares "product-rest-api/utils/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type productController struct {
	productUsecase entity.ProductsUseCaseInterface
}

func New(productUC entity.ProductsUseCaseInterface) *productController {
	return &productController{
		productUsecase: productUC,
	}
}

func (handler *productController) PostProduct(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(c)

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error get userId, login first",
		})
	}

	input := ProductRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error bind data",
		})
	}

	image, err := c.FormFile("Image")
	if err != nil {
		if err == http.ErrMissingFile {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "No file uploaded",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error uploading file",
		})
	}

	data := entity.ProductsCore{
		Product_image: input.Product_image,
		Product_name:  input.Product_name,
		Description:   input.Description,
		Quantity:      input.Quantity,
		Price:         input.Price,
	}

	_, errproduct := handler.productUsecase.PostProduct(data, image)
	if errproduct != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error post product",
			"error":   errproduct.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success post product",
		"data":    data,
	})
}

func (handler *productController) ReadAllProduct(c echo.Context) error {
	data, err := handler.productUsecase.ReadAllProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all product",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all product",
		"data":    data,
	})
}

func (handler *productController) ReadSpecificProduct(c echo.Context) error {
	idParamstr := c.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "product not found",
		})
	}

	data, err := handler.productUsecase.ReadSpecificProduct(idParams.String())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific product",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get product",
		"data":    data,
	})
}

func (handler *productController) UpdateProduct(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get userId, login first, login first",
		})
	}

	idParams := c.Param("id")

	data := new(ProductRequest)
	if errBind := c.Bind(data); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	image, err := c.FormFile("Image")
	if err != nil {
		if err == http.ErrMissingFile {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "No file uploaded",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error uploading file",
		})
	}

	productData := entity.ProductsCore{
		Product_image: data.Product_image,
		Product_name:  data.Product_name,
		Description:   data.Description,
		Quantity:      data.Quantity,
		Price:         data.Price,
	}

	updatedProduct, err := handler.productUsecase.UpdateProduct(idParams, productData, image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating product",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product updated successfully",
		"data":    updatedProduct,
	})

}

func (handler *productController) DeleteProduct(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get userId, login first",
		})
	}

	idParams := c.Param("id")
	err := handler.productUsecase.DeleteProduct(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting product",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product deleted successfully",
	})
}
