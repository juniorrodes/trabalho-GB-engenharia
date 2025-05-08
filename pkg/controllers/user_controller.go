package controllers

import (
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	userDb "github.com/juniorrodes/trab-GB-engenharia/pkg/db/user"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	queries *userDb.Queries
}

type User struct {
	Name string `json:"name"`
}

func NewUserController(db *pgxpool.Pool) *UserController {
	return &UserController{
		queries: userDb.New(db),
	}
}

func (c *UserController) GetUser(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return ctx.String(http.StatusBadRequest, "Expect a number")
	}

	user, err := c.queries.GetUser(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Oops")
	}

	return ctx.JSON(http.StatusOK, User{Name: user.Name})
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	var user userDb.CreateUserParams
	ctx.Bind(&user)

	newUser, err := c.queries.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, "Oops")
	}

	return ctx.JSON(http.StatusOK, User{Name: newUser.Name})
}
