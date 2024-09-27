package server

import (
	"strconv"
	"time"
)

type Option func(*Server)

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

func Addr(addr string) Option {
	return func(s *Server) {
		s.server.Addr = addr
	}
}

func HostPortAddr(host string, port int) Option {
	return func(s *Server) {
		s.server.Addr = host + ":" + strconv.Itoa(port)
	}
}
