package main

import (
	"net/http"
	"testing"
)

func TestSearch_CurrentPage(t *testing.T) {
	type fields struct {
		Query      string
		NextPage   int
		TotalPages int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "1", fields: fields{Query: "war", NextPage: 1, TotalPages: 4}, want: 1},
		{name: "2", fields: fields{Query: "war", NextPage: 3, TotalPages: 4}, want: 2},
		{name: "3", fields: fields{Query: "war", NextPage: 4, TotalPages: 4}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Query:      tt.fields.Query,
				NextPage:   tt.fields.NextPage,
				TotalPages: tt.fields.TotalPages,
			}
			if got := s.CurrentPage(); got != tt.want {
				t.Errorf("CurrentPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_IsLastPage(t *testing.T) {
	type fields struct {
		Query      string
		NextPage   int
		TotalPages int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "2", fields: fields{Query: "war", NextPage: 2, TotalPages: 4}, want: false},
		{name: "4", fields: fields{Query: "war", NextPage: 4, TotalPages: 4}, want: true},
		{name: "5", fields: fields{Query: "war", NextPage: 5, TotalPages: 4}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Query:      tt.fields.Query,
				NextPage:   tt.fields.NextPage,
				TotalPages: tt.fields.TotalPages,
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
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "5", fields: fields{Query: "war", NextPage: 5, TotalPages: 4}, want: 3},
		{name: "4", fields: fields{Query: "war", NextPage: 4, TotalPages: 4}, want: 2},
		{name: "3", fields: fields{Query: "war", NextPage: 3, TotalPages: 4}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Query:      tt.fields.Query,
				NextPage:   tt.fields.NextPage,
				TotalPages: tt.fields.TotalPages,
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
	}{}
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
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchHandler(tt.args.w, tt.args.r)
		})
	}
}
