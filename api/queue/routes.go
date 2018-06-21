package queue

import (
	"github.com/catmullet/Raithe/app/services/registration"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	g := e.Group("/queue")

	g.POST("/push", Push)
	g.POST("/pop", Pop)

	auth := e.Group("/auth")

	auth.GET("/dump_tokens", registration.DumpTokens)
	auth.DELETE("/invalidate_tokens", registration.InvalidateTokens)
	auth.POST("/register", registration.RegisterAsAgent)
}
