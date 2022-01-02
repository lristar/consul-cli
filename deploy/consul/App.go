package main

import (
	"context"
	"github.com/gin-gonic/gin"
)

type app struct {
	g      *gin.Engine
	ctx    context.Context
	cancel func()
}

func newApp(engine *gin.Engine) *app {
	return &app{g: engine}
}
