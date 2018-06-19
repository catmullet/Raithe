package queue

import (
	"github.com/labstack/echo"
	"github.com/catmullet/Raithe/app/auth/services"
)

func Routes(e *echo.Echo) {
	g := e.Group("/queue")

	g.POST("/push", Push)
	g.POST("/pop", Pop)

	auth := e.Group("/auth")

	auth.GET("/dump_tokens", services.DumpTokens)
	auth.DELETE("/invalidate_tokens", services.InvalidateTokens)
	auth.POST("/register", services.RegisterAsAgent)
}
