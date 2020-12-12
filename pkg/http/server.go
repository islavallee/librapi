package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
}

// NewHandler create a new http handler with routing handlers
func NewHandler() *Handler {

	h := &Handler{
		engine: gin.Default(),
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	h.engine.Use(cors.New(corsConfig))

	h.engine.GET("/hc", h.Healthcheck)

	return h
}

func (h *Handler) Engine() *gin.Engine {
	return h.engine
}

type Server struct {
	handler *Handler
}

// NewServer create a new http server
func NewServer(h *Handler) *Server {
	return &Server{
		handler: h,
	}
}

// Serve launch the web server
func (s *Server) Serve(port int) {
	if err := s.handler.Engine().Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
