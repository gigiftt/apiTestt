package main

import (
	"apiTestt/handler"
	"apiTestt/repo"
	"apiTestt/service/user"
	"log"

	// "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	// "google.golang.org/grpc/profiling/service"
	// "google.golang.org/grpc/profiling/service"
)


func main() {
    e := echo.New()

    db, err := sqlx.Connect("mysql", "root:@(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatalln(err)
	}

    apiWithoutJWT := e.Group("")

    userRepo := repo.NewuserRepository(db)

    userService := user.NewUserService(userRepo)

    handler.NewUserHandler(apiWithoutJWT,userService)

    e.Logger.Fatal(e.Start(":8080"))
}

