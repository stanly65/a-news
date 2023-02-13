package api

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestArticle_FormatPublishedDate(t *testing.T) {
	type fields struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		}
		Author      string
		Title       string
		Description string
		URL         string
		URLToImage  string
		PublishedAt time.Time
		Content     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Article{
				Source:      tt.fields.Source,
				Author:      tt.fields.Author,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				URL:         tt.fields.URL,
				URLToImage:  tt.fields.URLToImage,
				PublishedAt: tt.fields.PublishedAt,
				Content:     tt.fields.Content,
			}
			if got := a.FormatPublishedDate(); got != tt.want {
				t.Errorf("FormatPublishedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FetchEverything(t *testing.T) {
	type fields struct {
		c        *http.Client
		key      string
		PageSize int
	}
	type args struct {
		query string
		page  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Results
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c:        tt.fields.c,
				key:      tt.fields.key,
				PageSize: tt.fields.PageSize,
			}
			got, err := c.FetchEverything(tt.args.query, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchEverything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchEverything() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		c        *http.Client
		key      string
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.c, tt.args.key, tt.args.pageSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
