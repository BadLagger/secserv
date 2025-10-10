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

	"github.com/gorilla/mux"
)

func main() {
	log := utils.GlobalLogger().SetLevel(utils.Debug)
	defer log.Log(utils.Info, "App DONE!!!")

	log.Log(utils.Info, "App run!!!")

	appCfg := utils.CfgLoad("SecServ")

	if appCfg.SSLenable {
		log.Info("SSL enable!")
		if len(appCfg.FullchainPemPath) == 0 {
			log.Critical("You should set environment FULLCHAIN_PEM")
			return
		}

		if len(appCfg.PrivateSSLPath) == 0 {
			log.Critical("You should set environment PRIVATE_SSL_PATH")
			return
		}
	} else {
		log.Info("SSL disable!")
	}

	if appCfg.YandexEnable {
		if len(appCfg.YandexId) == 0 {
			log.Critical("You should set enviromnent YANDEX_ID")
			return
		}

		if len(appCfg.YandexRedirectURL) == 0 {
			log.Critical("You should set enviromnent YANDEX_REDIRECT_URL")
			return
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainView := view.NewHtmlView()
	countServ := models.NewCountService()
	strServ := models.NewStringService(appCfg.YandexId, appCfg.YandexRedirectURL)
	mainCtrl := controllers.NewCountroller(countServ, strServ, mainView)

	router := mux.NewRouter()
	server := &http.Server{
		Addr:    appCfg.HostAddress,
		Handler: router,
	}
	serverErr := make(chan error, 1)

	go func() {

		router.NotFoundHandler = http.HandlerFunc(mainCtrl.NotFoundHandler)

		router.HandleFunc("/", mainCtrl.MainPageHandler).Methods("GET")
		router.HandleFunc("/enter", mainCtrl.PrivateCabPageHandler).Methods("GET")
		/*if appCfg.YandexEnable {
			router.HandleFunc("/", mainCtrl.IndexHandler).Methods("GET")
			router.HandleFunc("/yandex_oauth", mainCtrl.YandexAuthHandler).Methods("GET")
		} else {
			router.HandleFunc("/", mainCtrl.MockHandler).Methods("GET")
		}*/

		log.Info("Try to start server...")
		var err error
		if appCfg.SSLenable {
			err = server.ListenAndServeTLS(appCfg.FullchainPemPath, appCfg.PrivateSSLPath)
		} else {
			err = server.ListenAndServe()
		}
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
