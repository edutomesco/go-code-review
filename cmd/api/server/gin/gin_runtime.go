package gin

import (
	"context"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type GinRuntime struct {
	Name     string
	Client   *GinEngine
	Shutdown time.Duration
}

func NewGinRuntime(name string, c *GinEngine, shutdownTime time.Duration) *GinRuntime {
	e := GinRuntime{
		Name:     name,
		Client:   c,
		Shutdown: shutdownTime,
	}

	return &e
}

func (e *GinRuntime) SetHandlers(f func(*gin.Engine)) {
	e.Client.Engine.Use(ErrorMiddleware())

	f(e.Client.Engine)
}

func (e *GinRuntime) Run(ctx context.Context) {
	srv := &http.Server{
		Addr:    e.Client.Addr,
		Handler: e.Client.Engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server running on %s", e.Client.Addr)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), e.Shutdown)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error while closing gin runtime: %s", err.Error())
	} else {
		log.Printf("gin runtime closed successfully")
	}
}
