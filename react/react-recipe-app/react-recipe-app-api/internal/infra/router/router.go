package router

import (
	"kedamaonsen/react-recipe-app-api/internal/infra/router/route"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router echo.Echo

func CreateRouter() *Router {

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())

	e.GET("/search", route.Search)

	r := Router(*e)
	return &r
}

func (r *Router) StartApp() {

	e := echo.Echo(*r)

	//とりあえずポートは固定
	e.Start(":9999")
}
