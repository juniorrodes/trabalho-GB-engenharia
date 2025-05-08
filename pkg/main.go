package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/juniorrodes/trab-GB-engenharia/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbPool, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@db:5432/mydb")
	if err != nil {
		log.Fatal("failed to connect to DB :(")
	}

	userController := controllers.NewUserController(dbPool)

	e.GET("/users/:id", userController.GetUser)
	e.POST("/users", userController.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
