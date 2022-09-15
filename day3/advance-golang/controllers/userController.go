package controllers

import (
	"advance-golang/config"
	"advance-golang/lib/database"
	"advance-golang/middlewares"
	"advance-golang/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetSpecificUser(c echo.Context) error {
	id := c.Param("id")

	user, e := database.GetUserByID(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func AddNewUser(c echo.Context) error {

	user := models.Users{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success created",
		"user":    user,
	})
}

func UpdateUser(c echo.Context) error {

	id := c.Param("id")

	user := models.Users{}
	c.Bind(&user)

	if err := config.DB.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updated",
		"user":    user,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	_, e := database.DeleteUserByID(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success deleted",
	})
}

func LoginUser(c echo.Context) error {
	users := models.Users{}
	c.Bind(&users)

	user, e := database.UserLogin(&users)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "login success",
		"users":  user,
	})
}

func Me(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := database.DetailUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "get profile data",
		"user":   users,
	})
}

func MyProfile(c echo.Context) error {
	userIdHasLogged := middlewares.ExtractTokenUser(c)

	idUser, _ := strconv.Atoi(userIdHasLogged)

	if idUser == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "User Not Found")
	}

	users, err := database.DetailUser(idUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "get profile data",
		"user":   users,
	})
}
