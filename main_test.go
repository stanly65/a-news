package main

import (
	"github.com/stanly65/a-news/api"
	"net/http"
	"testing"
)

func TestSearch_IsLastPage(t *testing.T) {
	type fields struct {
		Query      string
		NextPage   int
		TotalPages int
		Results    *api.Results
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Query:      tt.fields.Query,
				NextPage:   tt.fields.NextPage,
				TotalPages: tt.fields.TotalPages,
				Results:    tt.fields.Results,
			}
			if got := s.IsLastPage(); got != tt.want {
				t.Errorf("IsLastPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_CurrentPage(t *testing.T) {
	type fields struct {
		Query      string
		NextPage   int
		TotalPages int
		Results    *api.Results
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Query:      tt.fields.Query,
				NextPage:   tt.fields.NextPage,
				TotalPages: tt.fields.TotalPages,
				Results:    tt.fields.Results,
			}
			if got := s.CurrentPage(); got != tt.want {
				t.Errorf("CurrentPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_IsLastPage1(t *testing.T) {
	type fields struct {
		Query      string
		NextPage   int
		TotalPages int
		Results    *api.Results
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Query:      tt.fields.Query,
				NextPage:   tt.fields.NextPage,
				TotalPages: tt.fields.TotalPages,
				Results:    tt.fields.Results,
			}
			if got := s.IsLastPage(); got != tt.want {
				t.Errorf("IsLastPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_PreviousPage(t *testing.T) {
	type fields struct {
		Query      string
		NextPage   int
		TotalPages int
		Results    *api.Results
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Query:      tt.fields.Query,
				NextPage:   tt.fields.NextPage,
				TotalPages: tt.fields.TotalPages,
				Results:    tt.fields.Results,
			}
			if got := s.PreviousPage(); got != tt.want {
				t.Errorf("PreviousPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		in1 *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseHandler(tt.args.w, tt.args.in1)
		})
	}
}

func Test_searchHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchHandler(tt.args.w, tt.args.r)
		})
	}
}
