package middle_ware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
)

// CustomHttpErrorHandler 自定义的全局错误处理函数
// 兼容以下两种情况：
// 1. 普通错误：return errors.New("不能存在记录")
// 2. panic异常：panic("无法绑定请求参数")
func CustomHttpErrorHandler(err error, c echo.Context) {
	// 获取Echo的上下文日志
	logger := c.Logger()

	// 记录详细的错误信息
	logger.Errorf("Handler error: %v", err)

	// 先尝试将错误转换为 *echo.HTTPError
	var he *echo.HTTPError
	ok := errors.As(err, &he)

	if !ok {
		// 如果不是 echo.HTTPError，检查是否是panic错误
		// 由于middleware.Recover()会将panic转换为*echo.HTTPError，
		// 这里我们直接处理原始错误信息
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(), // 使用原始错误信息而不是通用文本
		}
	} else {
		// 如果是echo.HTTPError，但Message是默认的，也使用原始错误
		if he.Message == http.StatusText(he.Code) {
			he.Message = err.Error()
		}
	}

	// 对于panic错误，记录堆栈信息
	if he.Code == http.StatusInternalServerError {
		stack := debug.Stack()
		logger.Errorf("Stack trace: %s", stack)
	}

	// 返回 JSON 格式的错误响应
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			// 使用 resp.ResultVo 作为统一的返回格式
			// 将 interface{} 转换为 string 类型
			result := resp.ResultVo{
				Status: he.Code,
				Msg:    fmt.Sprintf("%v", he.Message),
				Data:   nil,
			}
			err = c.JSON(http.StatusOK, result)
		}
		if err != nil {
			logger.Error(err)
		}
	}
}
