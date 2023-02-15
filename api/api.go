// Package api create an API client for the purpose of working with the News API https://newsapi.org/
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Article struct mirrors the data when decoding the response body
type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

// FormatPublishedDate day, month with words, year
func (a *Article) FormatPublishedDate() string {
	year, month, day := a.PublishedAt.Date()
	return fmt.Sprintf("%d %v, %d", day, month, year)
}

// Results struct mirrors the data when decoding the response body
type Results struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Client struct represents the client for working with the News API
// "c" field points to the HTTP client that should be used to make requests
// "key" field holds the API key
// "PageSize" field holds the number of results to return per page
type Client struct {
	c        *http.Client
	key      string
	PageSize int
}

// FetchEverything endpoint accepts two arguments: the search query, and the page number
// These are appended to the request URL in addition to the API key and page size
func (c *Client) FetchEverything(query, page string) (*Results, error) {
	endpoint := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=en&searchin=title", url.QueryEscape(query), c.PageSize, page, c.key)

	resp, err := c.c.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errStr := fmt.Sprintf("Unable to read from body %s", err)
		return nil, errors.New(errStr)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &Results{}
	return res, json.Unmarshal(body, res)
}

// NewClient creates and returns a new Client instance
func NewClient(c *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{c, key, pageSize}
}
