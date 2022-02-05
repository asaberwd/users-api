package api

import (
	"net/http"

	"github.com/asaberwd/users-api/internal/user"
	"github.com/labstack/echo/v4"
)

// UserHandler ...
type UserHandler struct {
	UserManager user.Manager
}

// NewUserHandler ...
func NewUserHandler(userManager user.Manager) *UserHandler {
	return &UserHandler{UserManager: userManager}
}

// List ...
func (e *UserHandler) List(ctx echo.Context) error {
	res, err := e.UserManager.List(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

// Show ...
func (e *UserHandler) Show(ctx echo.Context) error {
	email := ctx.Param("email")
	res, err := e.UserManager.Show(email)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

// Add ...
func (e *UserHandler) Add(ctx echo.Context) error {
	u := &user.User{
		Email:     ctx.Param("email"),
		Firstname: ctx.Param("firstName"),
		Lastname:  ctx.Param("lastName"),
		Isactive:  true,
		Role:      ctx.Param("role"),
	}
	err := e.UserManager.Add(u)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, u)
}

// Update ...
func (e *UserHandler) Update(ctx echo.Context) error {
	u := &user.User{
		Email:     ctx.Param("email"),
		Firstname: ctx.Param("firstName"),
		Lastname:  ctx.Param("lastName"),
		Isactive:  true,
		Role:      ctx.Param("role"),
	}

	err := e.UserManager.Update(u)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, u)
}

// Router ...
func Router(e *echo.Echo, handler *UserHandler) {
	e.POST("/users", handler.Add)
	e.POST("/users/:id", handler.Update)

	e.GET("/users", func(c echo.Context) error { return handler.List(c) })
	e.GET("/users/:id", handler.Show)
}
