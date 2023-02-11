package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stanly65/a-news/api"
	te "html/template"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// Create new Template with base name and parsed content; panic error !=0
var tpl = te.Must(te.ParseFiles("./templates/index.html"))
var nc *api.Client

type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *api.Results
}

func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}

func baseHandler(w http.ResponseWriter, _ *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		return
	}

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	results, err := nc.FetchEverything(searchQuery, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nextPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	search := &Search{
		Query:      searchQuery,
		NextPage:   nextPage,
		TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(nc.PageSize))),
		Results:    results,
	}

	if ok := !search.IsLastPage(); ok {
		search.NextPage++
	}

	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		return
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	nc = api.NewClient(myClient, apiKey, 2)

	// instantiate a file server object
	fs := http.FileServer(http.Dir("static"))

	// create an HTTP request multiplexer
	// matches the URL of incoming requests against
	//a list of registered patterns
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/", baseHandler)

	fmt.Println("Server started at port 6789")
	err = http.ListenAndServe(":6789", mux)
	if err != nil {
		return
	}
}
