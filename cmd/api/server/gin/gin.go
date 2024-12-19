package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	Addr   string
	Debug  bool
	Engine *gin.Engine
}

func NewGinEngine(addr string, port int, debug bool) *GinEngine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return &GinEngine{
		Addr:   fmt.Sprintf("%s:%d", addr, port),
		Debug:  debug,
		Engine: gin.New(),
	}
}
