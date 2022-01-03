package deploy

import "github.com/gin-gonic/gin"

type Builder struct {
	C IConstructor
}

type IConstructor interface {
	NewHttpServer() *gin.Engine
	NewClient()
	NewService()
	NewDataServer()
}

func (b *Builder) InitApp() *app {
	b.C.NewDataServer()
	b.C.NewClient()
	b.C.NewService()
	e := b.C.NewHttpServer()
	return newApp(e)
}
