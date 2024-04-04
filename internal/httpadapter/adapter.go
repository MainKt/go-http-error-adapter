package httpadapter

import "net/http"

type HandlerFuncErr func(http.ResponseWriter, *http.Request) error

type ErrorHandler func(HandlerFuncErr) http.HandlerFunc

type ServeMuxErr struct {
	http.ServeMux
	errorHandler ErrorHandler
}

func (mux *ServeMuxErr) HandleFuncErr(pattern string, handler HandlerFuncErr) {
	mux.Handle(pattern, mux.errorHandler(handler))
}

func NewServeMuxErr(errorHandler ErrorHandler) *ServeMuxErr {
	return &ServeMuxErr{
		ServeMux:     http.ServeMux{},
		errorHandler: errorHandler,
	}
}
