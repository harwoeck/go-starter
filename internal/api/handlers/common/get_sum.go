package common

import (
	"fmt"
	"net/http"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/util"
	"github.com/labstack/echo/v4"
)

func GetSumRoute(s *api.Server) *echo.Route {
	return s.Router.Management.GET("/sum/:count", getSumHandler(s))
}

// :count "1" =>   1      => "1"
// :count "2" =>   1+2    => "3"
// :count "3" =>   1+2+3  => "6"
func getSumHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cnt int64
		if err := echo.PathParamsBinder(c).Int64("count", &cnt).BindError(); err != nil {
			return c.String(http.StatusBadRequest, "Please provide an integer.\n")
		}
		if cnt < 0 {
			return c.String(http.StatusBadRequest, "Please provide a positive integer.\n")
		}

		t := util.TriangularNumber(uint64(cnt))
		return c.String(http.StatusOK, fmt.Sprint(t, "\n"))
	}
}
