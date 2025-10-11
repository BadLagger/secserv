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
	// Сохраняем в куки
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token.AccessToken,
		Path:     "/",
		MaxAge:   3600, // 1 час
		HttpOnly: true, // защита от XSS
		Secure:   true, // только HTTPS
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "user_email",
		Value:    userInfo.Email,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: false, // чтобы JS мог прочитать если нужно
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	// Возвращаем страницу которая закроет popup
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
    <!DOCTYPE html>
    <html>
    <script>
        // Просто закрываем окно - данные уже в куках!
        window.close();
    </script>
    <body>
        Авторизация успешна! Закрываю окно...
    </body>
    </html>
    `)
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

func (c *Controller) PrivateCabPageWithYandexHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET privatecabpage with yandex request from: %s", r.RemoteAddr)
	emailCookie, err := r.Cookie("user_email")
	if err != nil {
		c.log.Error("No email cookie!!")
		emailCookie.Value = "unknown"
	}
	c.log.Debug("User email: %s", emailCookie.Value)
	c.htmlView.GetPrivateCabWithYandexPage(r.RemoteAddr, emailCookie.Value, &w)
}

func (c *Controller) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Info("GET logout request from: %s", r.RemoteAddr)
	cookies := []string{"auth_token", "user_email"}
	for _, name := range cookies {
		http.SetCookie(w, &http.Cookie{
			Name:   name,
			Value:  "",
			Path:   "/",
			MaxAge: -1, // Удалить куку
		})
	}

	// Редирект на главную
	http.Redirect(w, r, "/", http.StatusFound)
}
