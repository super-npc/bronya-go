package resp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResultVo struct {
	Status     int         `json:"status"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	DataDetail struct {
		BeanType string `json:"beanType"`
		IsArray  bool   `json:"isArray"`
	} `json:"dataDetail"`
	Cookies interface{} `json:"cookies"`
}

func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, ResultVo{Status: 200, Msg: "success", Data: data})
}

func Fail(c echo.Context, msg string) error {
	return c.JSON(http.StatusOK, ResultVo{Status: 500, Msg: msg})
}
