package controllers

import (
	"net/http"
	"secserv/models"
	"secserv/view"
)

type Controller struct {
	counter  *models.CountService
	strData  *models.StringService
	htmlView *view.HtmlView
}

func NewCountroller(countServ *models.CountService, strServ *models.StringService, htmlView *view.HtmlView) *Controller {
	return &Controller{
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
