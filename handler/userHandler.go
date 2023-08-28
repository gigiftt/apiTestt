package handler

import (
	"apiTestt/service/user"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type userHandler struct {
	userService user.UserServiceInt
}

func NewUserHandler(version *echo.Group, userService user.UserServiceInt) userHandler {
	handler := &userHandler{
		userService: userService,
	}

	userRoutes := version.Group("/user")

	userRoutes.GET("/", handler.GetAllUserHdl)
	userRoutes.POST("/", handler.CreateUserHdl)
	userRoutes.POST("/:name", handler.UpdateUserHdl)
	userRoutes.POST("/delete", handler.DeleteUserHdl)
	userRoutes.POST("/form", handler.GetUserByFormHdl)
	userRoutes.POST("/file", handler.GetFileHdl)

	return *handler

}

func (uh *userHandler) GetAllUserHdl(c echo.Context) error {

	users, err := uh.userService.GetAllUserServ(c.Request().Context())
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (uh *userHandler) CreateUserHdl(c echo.Context) error {

	var user user.UserModel

	err := c.Bind(&user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	err = uh.userService.CreateUserServ(c.Request().Context(), user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"Create": "success"})
}

func (uh *userHandler) UpdateUserHdl(c echo.Context) error {

	var user user.UserModel

	err := c.Bind(&user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	err = uh.userService.UpdateUserServ(c.Request().Context(), user, c.Param("name"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"Update": "success"})
}

func (uh *userHandler) DeleteUserHdl(c echo.Context) error {

	log.Print(c.QueryParam("name"))

	err := uh.userService.DeleteUserServ(c.Request().Context(), c.QueryParam("name"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"Delete": "success"})
}

func (uh *userHandler) GetUserByFormHdl(c echo.Context) error {
	log.Printf(c.FormValue("name"))

	user, err := uh.userService.GetUserByFormServ(c.Request().Context(), c.FormValue("name"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *userHandler) GetFileHdl(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	file := form.File["file"]

	localFileName := "./uploads/" + file[0].Filename
	out, err := os.Create(localFileName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	defer out.Close()

	readFile, err := file[0].Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	_, err = io.Copy(out, readFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"Message": "success"})

}
