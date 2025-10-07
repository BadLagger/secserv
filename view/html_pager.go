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
