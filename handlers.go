package main

import (
	"net/http"
	"github.com/labstack/echo"
)

// GetCategories
func GetCategories(c echo.Context) (err error) {
	if categories, count := dbGetCategories(); count > 0 {
		n := Categories{categories}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}

func GetNcms(c echo.Context) (err error) {
	if ncms, count := dbGetNcms(); count > 0 {
		n := NCMS{ncms}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}

func GetNcmsByUser(c echo.Context) (err error) {
	userID := c.Param("id")

	if ncms, count := dbGetNcmsByUser(userID); count > 0 {
		n := NCMS{ncms}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}

func GetLeadsTrial(c echo.Context) (err error) {
	if leads, count := dbGetLeadsTrial(); count > 0 {
		l := Leads{leads}
		err = c.JSON(http.StatusOK, l)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}
