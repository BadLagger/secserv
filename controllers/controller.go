package controllers

import (
	"net/http"
	"secserv/models"
	"secserv/utils"
	"secserv/view"
)

type Controller struct {
	log      *utils.Logger
	counter  *models.CountService
	strData  *models.StringService
	htmlView *view.HtmlView
}

func NewCountroller(countServ *models.CountService, strServ *models.StringService, htmlView *view.HtmlView) *Controller {
	return &Controller{
		log:      utils.GlobalLogger(),
		counter:  countServ,
		strData:  strServ,
		htmlView: htmlView,
	}
}

func (c *Controller) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	c.htmlView.GetNotFoundPage(&w)
}

func (c *Controller) MockHandler(w http.ResponseWriter, r *http.Request) {
	c.htmlView.GetMockPage("Test Mock Page", &w)
}

func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {
	//c.htmlView.GetIndexPage(c.counter.IncrementAndGet(), &w)
	c.htmlView.GetAuthPage(c.strData.GetWelcomeMsg(), c.strData.YandexId, c.strData.YandexRedirectURI, &w)
}

func (c *Controller) YandexAuthHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) MainPageHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET mainpage request from: %s", r.RemoteAddr)
	c.htmlView.GetMainPage(r.RemoteAddr, &w)
}

func (c *Controller) MainPageWithYandexHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET mainpage with yandex request from: %s", r.RemoteAddr)
	c.htmlView.GetMainPageWithYandex(r.RemoteAddr, c.strData.YandexId, c.strData.YandexRedirectURI, &w)
}

func (c *Controller) PrivateCabPageHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET privatecabpage request from: %s", r.RemoteAddr)
	c.htmlView.GetPrivateCabPage(r.RemoteAddr, &w)
}
