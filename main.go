// Package main running http and file servers; implementing the received data into the web template
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

// Create new Template with base name and parsed content; panic if error !=0
var tpl = te.Must(te.ParseFiles("./templates/index.html"))
var nc *api.Client

type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *api.Results
}

// IsLastPage checks if this is the last page
func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

// CurrentPage returns the current page
func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

// PreviousPage returns the previous page
func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}

// baseHandler - response loads an empty template with no data
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

// searchHandler processes the results and loads a data template
func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// q represents the user’s query, and page is used to page through the results
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
	// fmt.Printf("%+v", results) // test display of the result in the console

	// Rendering the results
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
	// execute the template, passing the search struct as the data object
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

	adr := os.Getenv("ADDRESS")
	if adr == "" {
		adr = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("NEWS_API_KEY must be set in .env file")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	nc = api.NewClient(myClient, apiKey, 2)

	// instantiate a file server object
	fs := http.FileServer(http.Dir("static"))

	// create an HTTP request multiplexer
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/", baseHandler)

	// start http server
	fmt.Printf("Starting server on http://%v:%v", adr, port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println("Failed to start server:")
		log.Fatal(err)
	}
}
