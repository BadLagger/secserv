package view

import (
	"html/template"
	"net/http"
)

type HtmlView struct {
	indexPath  string
	simplePath string
}

func NewHtmlView() *HtmlView {
	return &HtmlView{
		indexPath:  "templates/index.html",
		simplePath: "templates/simple.html",
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

func (v *HtmlView) GetSimplePage(title string, btn1Msg string, btn2Msg string, w *http.ResponseWriter) {
	t, err := template.ParseFiles(v.simplePath)
	if err != nil {
		http.Error(*w, "Failed get simple path", http.StatusInternalServerError)
		return
	}

	t.Execute(*w, map[string]interface{}{
		"MockMsg":     title,
		"Button1Name": btn1Msg,
		"Button2Name": btn2Msg,
	})
}
