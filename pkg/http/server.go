package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/islavallee/librapi/pkg/application"
	"github.com/islavallee/librapi/pkg/storage"
)

type Handler struct {
	engine  *gin.Engine
	librapi *application.Librapi
}

// NewHandler create a new http handler with routing handlers
func NewHandler() *Handler {
	repository := storage.NewBoltDB("/var/librapi/librapi.db")

	h := &Handler{
		engine: gin.Default(),
		librapi: application.NewLibrapi(
			repository,
		),
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	h.engine.Use(cors.New(corsConfig))

	h.engine.GET("/hc", h.Healthcheck)
	h.engine.POST("/datas", h.PostData)
	h.engine.GET("/datas/:key", h.GetData)
	h.engine.PUT("/datas/:key", h.PutData)
	h.engine.DELETE("/datas/:key", h.DeleteData)

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
