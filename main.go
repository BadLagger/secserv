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

	if len(appCfg.FullchainPemPath) == 0 {
		log.Critical("You should set environment FULLCHAIN_PEM")
		return
	}

	if len(appCfg.PrivateSSLPath) == 0 {
		log.Critical("You should set environment PRIVATE_SSL_PATH")
		return
	}

	if len(appCfg.YandexId) == 0 {
		log.Critical("You should set enviromnent YANDEX_ID")
		return
	}

	if len(appCfg.YandexRedirectURL) == 0 {
		log.Critical("You should set enviromnent YANDEX_REDIRECT_URL")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainView := view.NewHtmlView()
	countServ := models.NewCountService()
	strServ := models.NewStringService(appCfg.YandexId, appCfg.YandexRedirectURL)
	mainCtrl := controllers.NewCountroller(countServ, strServ, mainView)

	server := &http.Server{Addr: appCfg.HostAddress}
	serverErr := make(chan error, 1)

	go func() {
		http.HandleFunc("/", mainCtrl.IndexHandler)
		http.HandleFunc("/yandex_oauth", mainCtrl.YandexAuthHandler)

		log.Info("Try to start server...")
		err := server.ListenAndServeTLS(appCfg.FullchainPemPath, appCfg.PrivateSSLPath)
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
