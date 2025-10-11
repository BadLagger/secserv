package controllers

import (
	"fmt"
	"net/http"
	"secserv/models"
	"secserv/utils"
	"secserv/view"
)

type Controller struct {
	log      *utils.Logger
	counter  *models.CountService
	ya       *models.YandexService
	htmlView *view.HtmlView
}

func NewCountroller(countServ *models.CountService, yaServ *models.YandexService, htmlView *view.HtmlView) *Controller {
	return &Controller{
		log:      utils.GlobalLogger(),
		counter:  countServ,
		ya:       yaServ,
		htmlView: htmlView,
	}
}

func (c *Controller) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	c.htmlView.GetNotFoundPage(&w)
}

func (c *Controller) MockHandler(w http.ResponseWriter, r *http.Request) {
	c.htmlView.GetMockPage("Test Mock Page", &w)
}

/*func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {
	//c.htmlView.GetIndexPage(c.counter.IncrementAndGet(), &w)
	c.htmlView.GetAuthPage(c.strData.GetWelcomeMsg(), c.strData.YandexId, c.strData.YandexRedirectURI, &w)
}*/

func (c *Controller) YandexAuthHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET yandex auth request from: %s", r.RemoteAddr)
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	c.log.Debug("Yandex code: %s", code)

	state := r.URL.Query().Get("state")
	c.log.Debug("Yandex state: %s", state)
	// Validate state
	token, err := c.ya.ExchangeCodeToken(code)
	if err != nil {
		http.Error(w, "Get token error", http.StatusBadRequest)
	}

	c.log.Debug("Token: %s", token.TokenType)

	userInfo, err := c.ya.GetUserInfo(token)
	if err != nil {
		http.Error(w, "Get user_info error", http.StatusBadRequest)
	}

	c.log.Debug("UserInfo: %s", userInfo.Email)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Авторизация успешна</title>
        <script>
            // Отправляем сообщение в родительское окно
            window.opener.postMessage({
                type: 'yandex_oauth_success',
                token: '%s',
                email: '%s',
                user_id: '%s'
            }, '*');
            
            // Закрываем окно
            window.close();
        </script>
    </head>
    <body>
        <p>Авторизация успешна! Закрываю окно...</p>
    </body>
    </html>
    `, token.AccessToken, userInfo.Email, userInfo.ID)
}

func (c *Controller) MainPageHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET mainpage request from: %s", r.RemoteAddr)
	c.htmlView.GetMainPage(r.RemoteAddr, &w)
}

func (c *Controller) MainPageWithYandexHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET mainpage with yandex request from: %s", r.RemoteAddr)
	c.htmlView.GetMainPageWithYandex(r.RemoteAddr, c.ya.YandexId, c.ya.YandexRedirectURI, &w)
}

func (c *Controller) PrivateCabPageHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET privatecabpage request from: %s", r.RemoteAddr)
	c.htmlView.GetPrivateCabPage(r.RemoteAddr, &w)
}
