package preact

import (
	"github.com/labstack/echo"
)

var httpRouter *Router

func init() {
	r, err := NewRouter()
	if err != nil {
		panic(err)
	}

	httpRouter = r
}

func HTTPHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path

		match := httpRouter.Match(path)

		if match == nil {
			return next(c)
		}

		return RenderPage(c, match)
	}
}
