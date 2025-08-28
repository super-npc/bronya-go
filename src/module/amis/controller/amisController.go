package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/util"
)

func Page(c echo.Context) error {
	_, err := util.NewStructFromJSONAndName("", []byte{})
	if err != nil {
		return errors.New("")
	}
	return c.String(http.StatusOK, "Hello, World!")
}
