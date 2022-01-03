package deploy

import (
	"context"
	"github.com/gin-gonic/gin"
)

type app struct {
	G      *gin.Engine
	Ctx    context.Context
	Cancel func()
}

func newApp(engine *gin.Engine) *app {
	return &app{G: engine}
}
