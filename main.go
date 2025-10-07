package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"secserv/utils"
	"syscall"
)

func main() {
	log := utils.GlobalLogger().SetLevel(utils.Debug)
	defer log.Log(utils.Info, "App DONE!!!")

	log.Log(utils.Info, "App run!!!")

	appCfg := utils.CfgLoad("SecServ")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := &http.Server{Addr: appCfg.HostAddress}
	serverErr := make(chan error, 1)

	go func() {
		log.Info("Try to start server...")
		err := server.ListenAndServe()
		if err != nil {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		server.Shutdown(ctx)
		log.Info("Escape signal received! Gentel shutdown!")
	case err := <-serverErr:
		log.Critical("Server can't run: %v", err)
	case <-ctx.Done():
		log.Error("Context DONE!!! Ouu my... this is terrible!")
	}
}
