package controllers

import (
	"net/http"
	"secserv/models"
	"secserv/view"
)

type Controller struct {
	counter  *models.CountService
	pager    *models.PageService
	strData  *models.StringService
	htmlView *view.HtmlView
}

func NewCountroller(countServ *models.CountService, pageServ *models.PageService, strServ *models.StringService, htmlView *view.HtmlView) *Controller {
	return &Controller{
		counter:  countServ,
		pager:    pageServ,
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
	c.htmlView.GetAuthPage(c.strData.GetWelcomeMsg(), c.strData.YandexId, c.strData.YandexRedirectURI, &w)
}

func (c *Controller) SimplePageHandler(w http.ResponseWriter, r *http.Request) {
	var page *models.PageMain = c.pager.GetPageById(0)
	c.htmlView.GetSimplePage(page.GetTitle(), page.GetButton()[0], &w)
}

func (c *Controller) SimpleEnterPageHandler(w http.ResponseWriter, r *http.Request) {
	var page *models.PageMain = c.pager.GetPageById(1)
	c.htmlView.GetSimplePage(page.GetTitle(), page.GetButton()[0], &w)
}

/*func (c *Controller) SimpleExitPageHandler(w http.ResponseWriter, r *http.Request) {

}*/

func (c *Controller) YandexAuthHandler(w http.ResponseWriter, r *http.Request) {

}
