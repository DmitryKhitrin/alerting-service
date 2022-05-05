package main

import (
	"context"
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/server"
	"github.com/DmitryKhitrin/alerting-service/internal/server/config"
	"github.com/caarlos0/env/v6"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	app := server.NewApp(&cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: server.GetRouter(app),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil {
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("context was canceled")
	case s := <-quit:
		log.Println("signal was provided: ", s)
		cancel()
	}

}
