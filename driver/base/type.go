package base

import (
	"io"
	"net/http"
)

const (
	TypeString = "string"
	TypeSelect = "select"
	TypeBool   = "bool"
	TypeNumber = "number"
	TypeText   = "text"
)

const (
	Get = iota
	Post
	Put
	Delete
	Patch
)

type Json map[string]any

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Link struct {
	Url      string   `json:"url"`
	Headers  []Header `json:"headers"`
	Data     io.ReadCloser
	FilePath string `json:"path"` // for native
	Status   int
	Header   http.Header
}
