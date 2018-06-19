package queue

import (
	"github.com/labstack/echo"
	"github.com/catmullet/Raithe/Auth/Services"
)

func Routes(e *echo.Echo) {
	g := e.Group("/queue")

	g.POST("/push", Push)
	g.POST("/pop", Pop)

	auth := e.Group("/auth")

	auth.GET("/dump_tokens", Services.DumpTokens)
	auth.DELETE("/invalidate_tokens", Services.InvalidateTokens)
	auth.POST("/register", Services.RegisterAsAgent)
}
