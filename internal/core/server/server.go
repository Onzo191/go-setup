package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	config "go-setup/internal/core/config"
)

type Module interface {
	RegisterRoutes(router *gin.RouterGroup)
	Name() string
}

type Server struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func New(cfg *config.Config) *Server {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery()) // panic -> 500

	return &Server{
		config:  cfg,
		router:  router,
		modules: []Module{},
	}
}

func (s *Server) RegisterModules(modules ...Module) {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	s.router.LoadHTMLGlob("./public/html/*")
	s.router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"goVersion":  runtime.Version(),
			"ginVersion": gin.Version,
		})
	})
	s.router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "not-found.html", gin.H{"error": "Not Found"})
	})

	// API version group
	v1 := s.router.Group("/api/v1")

	// Register routes for each module
	for _, module := range modules {
		module.RegisterRoutes(v1)
		s.modules = append(s.modules, module)
	}
}

func (s *Server) Start() {
	srv := &http.Server{
		Addr:         ":" + s.config.Port,
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			println("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		println("Server forced to shutdown")
	}
}
