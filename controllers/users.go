package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"userregistry/common/errors"
	"userregistry/models"
	"userregistry/service"
)

type updateResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type UsersController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
}

type usersController struct {
	usersService service.UserService
}

func NewUsersController(usersService service.UserService) UsersController {
	return &usersController{usersService: usersService}
}

func (uc *usersController) Get(c echo.Context) error {
	userName := c.Param("username")
	if userName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username must not be empty")
	}

	user, err := uc.usersService.Get(context.Background(), userName)
	if err != nil {
		if err == errors.ErrUserNotFound {
			return c.JSON(http.StatusOK, struct {
				status  int
				Message string
			}{
				status:  http.StatusNotFound,
				Message: "user not found",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *usersController) Create(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println("user", user, "is saved")
	expiryTime := time.Unix(user.ExpiryDate, 0)
	fmt.Println("Expiry Date", expiryTime)

	err := uc.usersService.Create(context.Background(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		status  int
		Message string
	}{
		status:  http.StatusCreated,
		Message: "user created",
	})
}

func (uc *usersController) Update(c echo.Context) error {
	userName := c.Param("username")
	if userName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username must not be empty")
	}

	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	expiryTime := time.Unix(user.ExpiryDate, 0)
	fmt.Println("Expiry Date", expiryTime)

	err := uc.usersService.Update(context.Background(), user, userName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ur := updateResponse{
		Status:  true,
		Message: fmt.Sprintf("The user %s has been successfully updated!", userName),
	}
	return c.JSON(http.StatusCreated, ur)
}
