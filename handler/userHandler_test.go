package handler_test

import (
	"apiTestt/handler"
	"apiTestt/repo"
	"apiTestt/service/user"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestArticleHandler_HandleCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"firstname":"firstname", "lastname":"lastname", "age":20, "email":"email@fmail.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	db, err := sqlx.Connect("mysql", "root:@(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatalln(err)
	}

	apiWithoutJWT := e.Group("")
	userRepo := repo.NewuserRepository(db)
	userService := user.NewUserService(userRepo)
	h := handler.NewUserHandler(apiWithoutJWT, userService)

	// Assertions
	if assert.NoError(t, h.CreateUserHdl(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"Create":"success"}`, strings.Trim(rec.Body.String(), "\n"))
	}
}
