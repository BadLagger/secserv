package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"secserv/controllers"
	"secserv/models"
	"secserv/utils"
	"secserv/view"
	"syscall"
)

func main() {
	log := utils.GlobalLogger().SetLevel(utils.Debug)
	defer log.Log(utils.Info, "App DONE!!!")

	log.Log(utils.Info, "App run!!!")

	appCfg := utils.CfgLoad("SecServ")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainView := view.NewHtmlView()
	countServ := models.NewCountService()
	strServ := models.NewStringService("877224072a4d47958acc083523b6f397", "http://95.165.171.184:8089/yandex_oauth")
	mainCtrl := controllers.NewCountroller(countServ, strServ, mainView)

	server := &http.Server{Addr: appCfg.HostAddress}
	serverErr := make(chan error, 1)

	go func() {
		http.HandleFunc("/", mainCtrl.IndexHandler)

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
