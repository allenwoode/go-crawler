package controller

import (
	"feilin.com/gocourse/goaction/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strings"
	"strconv"
	"feilin.com/gocourse/goaction/frontend/model"
	"context"
	"reflect"
	"feilin.com/gocourse/goaction/crawler/engine"
	"regexp"
	"log"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return SearchResultHandler{}
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	//q = formatQueryString(q)
	log.Println(q)

	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	//fmt.Fprintf(w, "q=%s, from=%d", q, from)
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("dating_profile").
							Query(elastic.NewQueryStringQuery(formatQueryString(q))).
							From(from).
							Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Query = q
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func formatQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}