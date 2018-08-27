package view

import (
	"html/template"
	"io"
	"frontend/model"
)

type SearchResultView struct {
	view *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		view: template.Must(template.ParseFiles(filename)),
	}
}

func (v SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return v.view.Execute(w, data)
}
