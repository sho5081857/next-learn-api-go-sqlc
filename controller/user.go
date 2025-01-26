package controller

import (
	"net/http"
	"next-learn-go-sqlc/infrastructure/database/sqlc"
	"next-learn-go-sqlc/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	GetUserById(c echo.Context) error
	GetUserByEmail(c echo.Context) error
}

type userController struct {
	uu usecase.UserUseCase
}

func NewUserController(uu usecase.UserUseCase) UserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := sqlc.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "User created")
}

func (uc *userController) LogIn(c echo.Context) error {
	user := sqlc.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tokenString)

}

func (uc *userController) GetUserById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	userRes, err := uc.uu.GetUserById(userId.(pgtype.UUID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) GetUserByEmail(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	userRes, err := uc.uu.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}
