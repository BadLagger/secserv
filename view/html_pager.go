package view

import (
	"html/template"
	"net/http"
)

type HtmlView struct {
	indexPath string
}

func NewHtmlView() *HtmlView {
	return &HtmlView{
		indexPath: "templates/index.html",
	}
}

func (v *HtmlView) GetNotFoundPage(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusNotFound)
	v.GetMockPage("Page Not Found 404", w)
}

func (v *HtmlView) GetIndexPage(count int, w *http.ResponseWriter) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(*w, "Failed get index.html", http.StatusInternalServerError)
		return
	}

	t.Execute(*w, map[string]interface{}{
		"Count": count,
	})
}

func (v *HtmlView) GetAuthPage(msg, id, uri string, w *http.ResponseWriter) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(*w, "Failed get index.html", http.StatusInternalServerError)
		return
	}

	t.Execute(*w, map[string]interface{}{
		"AuthWelcomeMsg": msg,
		"YaClientId":     id,
		"YaRedirectURI":  uri,
	})
}

func (v *HtmlView) GetMockPage(msg string, w *http.ResponseWriter) {
	t, err := template.ParseFiles("templates/moc.html")
	if err != nil {
		http.Error(*w, "Failed get moc.html", http.StatusInternalServerError)
		return
	}

	t.Execute(*w, map[string]interface{}{
		"MockMsg": msg,
	})
}

func (v *HtmlView) GetMainPage(ip string, w *http.ResponseWriter) {
	t, err := template.ParseFiles("templates/mifi_page1.html")
	if err != nil {
		http.Error(*w, "Failed get mifi_page1.html", http.StatusInternalServerError)
		return
	}

	t.Execute(*w, map[string]interface{}{
		"UserIp": ip,
	})
}

func (v *HtmlView) GetPrivateCabPage(ip string, w *http.ResponseWriter) {
	t, err := template.ParseFiles("templates/mifi_page2.html")
	if err != nil {
		http.Error(*w, "Failed get mifi_page2.html", http.StatusInternalServerError)
		return
	}

	t.Execute(*w, map[string]interface{}{
		"UserIp": ip,
	})
}
