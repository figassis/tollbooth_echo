package tollbooth_echo

import (
	tollbooth "github.com/didip/tollbooth/v6"
	limiter "github.com/didip/tollbooth/v6/limiter"
	"github.com/labstack/echo/v4"
)

func LimitMiddleware(lmt *limiter.Limiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			httpError := tollbooth.LimitByRequest(lmt, c.Response(), c.Request())
			if httpError != nil {
				return c.String(httpError.StatusCode, httpError.Message)
			}
			return next(c)
		})
	}
}

func LimitHandler(lmt *limiter.Limiter) echo.MiddlewareFunc {
	return LimitMiddleware(lmt)
}
