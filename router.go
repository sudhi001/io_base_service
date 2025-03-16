package io_base_service

import "net/http"

// Router defines a framework-independent router
type Router interface {
	Get(path string, handler func(c Context) error)
	Post(path string, handler func(c Context) error)
	Use(middleware func(c Context) error)
	Listen(port string) error
	TestRequest(req *http.Request) (*http.Response, error)
}
