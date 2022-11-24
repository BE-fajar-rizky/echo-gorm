package controllers

import (
	// "be13/mvc/helper"
	// "be13/mvc/models"
	// "be13/mvc/repositories"

	"fajar/apiMvc/helper"
	"fajar/apiMvc/models"
	"fajar/apiMvc/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {

	result, err := repository.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("berhasil membaca  user", result))
}

func AddUserController(c echo.Context) error {
	//
	Inputuser := models.User{}
	errbind := c.Bind(&Inputuser)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errbind.Error()))
	}
	errInsert := repository.InsertUser(Inputuser)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("gagal insert data"))
	}
	return c.JSON(http.StatusOK, helper.PesanSuksesHelper("berhasil membuat user"))
}

func GetAllUserControllerId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := repository.GetAllUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("berhasil membaca user", result))
}

func UpddateUserController(c echo.Context) error {
	//

	data := models.User{}
	c.Bind(&data)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("gagal baca id"))
	}
	result, errupdate := repository.UpdateUserId(id, data)
	if errupdate != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("gagal hapus data"))
	}
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("berhasil MENGHAPUS user", result))
}

func DeleteUserContoller(c echo.Context) error {
	//
	id, _ := strconv.Atoi(c.Param("id"))

	errdelete := repository.DeleteUser(id)
	if errdelete != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("gagal hapus data"))
	}
	return c.JSON(http.StatusOK, helper.PesanSuksesHelper("berhasil membuat user"))
}
