package queue

import "github.com/labstack/echo"

func Routes(e *echo.Echo) {

	g := e.Group("/queue")

	g.POST("/push", Push)
	g.GET("/pop/:queue", Pop)
}
