package io_base_service

import "github.com/gofiber/fiber/v2"

// FiberContext is an adapter to make Fiber work with the framework.Context
type FiberContext struct {
	Ctx *fiber.Ctx
}

// JSON sends a JSON response
func (f *FiberContext) JSON(data interface{}) error {
	return f.Ctx.JSON(data)
}

// BindJSON parses request JSON body
func (f *FiberContext) BindJSON(dest interface{}) error {
	return f.Ctx.BodyParser(dest)
}

// GetLocal retrieves a value set in the request context
func (f *FiberContext) GetLocal(key string) interface{} {
	return f.Ctx.Locals(key)
}

// Next calls the next middleware
func (f *FiberContext) Next() error {
	return f.Ctx.Next()
}

// SetLocal sets a local variable
func (fc *FiberContext) SetLocal(key string, value interface{}) {
	fc.Ctx.Locals(key, value)
}

// GetHeader gets a header value
func (fc *FiberContext) GetHeader(name string) string {
	return fc.Ctx.Get(name)
}

// JSONResponse sends a JSON response
func (fc *FiberContext) JSONResponse(statusCode int, data interface{}) error {
	return fc.Ctx.Status(statusCode).JSON(data)
}

// Path returns the request path
func (c *FiberContext) Path() string {
	return c.Ctx.Path()
}

// Method returns the HTTP method
func (c *FiberContext) Method() string {
	return c.Ctx.Method()
}

// IP returns the client's IP address
func (c *FiberContext) IP() string {
	return c.Ctx.IP()
}
