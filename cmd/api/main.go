package main

import (
	"github.com/edutomesco/coupons/cmd/api/server"
	"github.com/edutomesco/coupons/cmd/api/server/gin"
	"github.com/edutomesco/coupons/internal/config"
	"github.com/edutomesco/coupons/internal/controllers"
	"github.com/edutomesco/coupons/internal/datasources/memdb"
	"github.com/edutomesco/coupons/internal/services"
	"log"
	"sync"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := gin.NewGracefullShutdown()

	// Initialize repository
	cr := memdb.NewCouponRepository()

	// Initialize services
	cs := services.NewCouponService(cr)

	// Initialize controllers
	hc := controllers.NewHealthController()
	cc := controllers.NewCouponController(cs)

	engine := gin.NewGinEngine(cfg.Host, cfg.Port, cfg.Debug)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		srv, err := server.New(cfg.App, engine, hc, cc)
		if err != nil {
			log.Fatalln(err.Error())
		}
		srv.Run(ctx)
	}()

	// Run other components...

	wg.Wait()
}
