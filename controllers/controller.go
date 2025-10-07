package controllers

import (
	"net/http"
	"secserv/models"
	"secserv/view"
)

type Controller struct {
	counter  *models.CountService
	htmlView *view.HtmlView
}

func NewCountroller(countServ *models.CountService, htmlView *view.HtmlView) *Controller {
	return &Controller{
		counter:  countServ,
		htmlView: htmlView,
	}
}

func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {
	c.htmlView.GetIndexPage(c.counter.IncrementAndGet(), &w)
}
