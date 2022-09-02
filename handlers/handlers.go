package handlers

import "net/http"

type Handler interface {
	Handler() http.HandlerFunc
}
