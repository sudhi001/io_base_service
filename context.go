package io_base_service

// Context defines a framework-agnostic request context
type Context interface {
	JSON(data interface{}) error
	BindJSON(dest interface{}) error
	GetLocal(key string) interface{}
	GetHeader(name string) string
	SetLocal(key string, value interface{})
	JSONResponse(statusCode int, data interface{}) error
	Next() error
	Path() string
	Method() string
	IP() string
}
