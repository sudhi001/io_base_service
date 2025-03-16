package io_base_service

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// FiberRouter is a concrete implementation of Router using Fiber
type FiberRouter struct {
	app *fiber.App
}

// NewRouter initializes a new Fiber-based router
func NewRouter() Router {
	return &FiberRouter{app: fiber.New()}
}

// Get registers a GET route
func (r *FiberRouter) Get(path string, handler func(c Context) error) {
	r.app.Get(path, func(ctx *fiber.Ctx) error {
		return handler(&FiberContext{ctx})
	})
}

// Post registers a POST route
func (r *FiberRouter) Post(path string, handler func(c Context) error) {
	r.app.Post(path, func(ctx *fiber.Ctx) error {
		return handler(&FiberContext{ctx})
	})
}

// Use registers middleware
func (r *FiberRouter) Use(middleware func(c Context) error) {
	r.app.Use(func(ctx *fiber.Ctx) error {
		return middleware(&FiberContext{ctx})
	})
}

// Listen starts the server
func (r *FiberRouter) Listen(port string) error {
	return r.app.Listen(port)
}

// TestRequest allows testing HTTP requests
func (r *FiberRouter) TestRequest(req *http.Request) (*http.Response, error) {
	return r.app.Test(req)
}
