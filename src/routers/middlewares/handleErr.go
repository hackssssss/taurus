package middlewares

import (
	"net/http"
	"taurus/src/common"

	"github.com/labstack/echo"
)

func HandleErr(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if err = next(c); err != nil {
			if !c.Response().Committed {
				d := StdResp{
					Code: common.ErrCodeCommonErr,
					Msg:  err.Error(),
					Data: GetOutResponse(c),
				}
				return c.JSON(http.StatusOK, d)
			}
		}
		return
	}
}
