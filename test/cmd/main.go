/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package main

import (
	"context"
	"fmt"
)

// Handler defines the handler invoked by Middleware.
type Handler func(ctx context.Context, req interface{}) (interface{}, error)

// Middleware is HTTP/gRPC transport middleware.
type Middleware func(Handler) Handler

type ServerOption func(*Server)

type Server struct {
	middlewares []Middleware
	router      map[string]Handler
}

func NewServer(options ...ServerOption) *Server {

	s := &Server{router: make(map[string]Handler)}

	for _, option := range options {
		option(s)
	}

	return s
}

func (s *Server) AddRouter(path string, handler Handler) {
	if _, ok := s.router[path]; ok {
		panic("route already exists")
	}

	s.router[path] = handler
}

func (s *Server) Use(middlewares ...Middleware) {
	s.middlewares = append(s.middlewares, middlewares...)
}

func chainUnaryInterceptors(interceptors []Middleware) Middleware {
	return func(handler Handler) Handler {
		return interceptors[0](getChainUnaryHandler(interceptors, 0, handler))
	}
}

func getChainUnaryHandler(interceptors []Middleware, curr int, finalHandler Handler) Handler {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		handler := interceptors[curr+1](getChainUnaryHandler(interceptors, curr+1, finalHandler))
		return handler(ctx, req)
	}
}

func WithAuth() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			fmt.Println("auth --- start")
			res, err := next(ctx, req)

			fmt.Println("auth --- end")
			return res, err
		}
	}
}

func WithLogging() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			fmt.Println("logging --- start")
			res, err := next(ctx, req)

			fmt.Println("logging --- end")
			return res, err
		}
	}
}

func WithErr() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			fmt.Println("err --- start")
			res, err := next(ctx, req)

			fmt.Println("err --- end")
			return res, err
		}
	}
}

func main() {
	svr := NewServer()

	svr.Use(WithAuth(), WithLogging(), WithErr())

	svr.AddRouter("/foo", func(ctx context.Context, req interface{}) (interface{}, error) {
		fmt.Println("foo")
		return nil, nil
	})

	svr.AddRouter("/bar", func(ctx context.Context, req interface{}) (interface{}, error) {
		fmt.Println("bar")
		return nil, nil
	})

	// 模拟执行过程
	if handle, ok := svr.router["/foo"]; ok {
		interceptors := chainUnaryInterceptors(svr.middlewares)
		handler := interceptors(handle)
		handler(context.Background(), nil)
	}
}
