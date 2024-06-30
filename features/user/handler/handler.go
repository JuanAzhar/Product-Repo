package handler

import (
	"net/http"
	"product-rest-api/features/user/entity"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase entity.UserUseCaseInterface
}

func New(userUC entity.UserUseCaseInterface) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (handler *UserController) Register(e echo.Context) error {
	input := new(UserRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.UserCore{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	row, errUser := handler.userUsecase.Register(data)
	if errUser != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create user",
			"error":   errUser.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create user",
		"data":    row,
	})
}

func (handler *UserController) Login(e echo.Context) error {
	input := new(UserRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.UserCore{
		Email:    input.Email,
		Password: input.Password,
	}

	data, token, err := handler.userUsecase.Login(data.Email, data.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error login",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"email":   data.Email,
		"token":   token,
	})
}

func (handler *UserController) DeleteUser(e echo.Context) error {
	idParams := e.Param("id")
	err := handler.userUsecase.DeleteUser(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting user",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func (handler *UserController) ReadSpecificUser(e echo.Context) error {
	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "user not found",
		})
	}

	data, err := handler.userUsecase.ReadSpecificUser(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific user",
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get user",
		"data":    data,
	})
}
