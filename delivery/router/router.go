// Package router provides functionality to set up and run the HTTP server,
// manage routes, and apply middleware based on access levels.
//
// It configures and initializes routes with varying access requirements:
// - Public routes: Accessible without authentication.
// - Protected routes: Require authentication.
// - Privileged routes: Require both authentication and admin privileges.
package router

import (
	"log"
	"net/http"

	"bail/config"
	"bail/delivery/core"
	authmiddleware "bail/delivery/middleware"
	ijwt "bail/usecases/core/i_jwt"

	"github.com/gin-gonic/gin"
)

// Router manages the HTTP server and its dependencies,
// including controllers and JWT authentication.
type Router struct {
	addr        string
	baseURL     string
	controllers []core.IController
	jwtService  ijwt.Service
}

func (r *Router) Use(f func(next http.Handler) http.Handler) {
	panic("unimplemented")
}

// Config holds configuration settings for creating a new Router instance.
type Config struct {
	Addr        string             // Address to listen on
	BaseURL     string             // Base URL for API routes
	Controllers []core.IController // List of controllers
	JwtService  ijwt.Service       // JWT service
}

// NewRouter creates a new Router instance with the given configuration.
// It initializes the router with address, base URL, controllers, and JWT service.
func NewRouter(config Config) *Router {
	return &Router{
		addr:        config.Addr,
		baseURL:     config.BaseURL,
		controllers: config.Controllers,
		jwtService:  config.JwtService,
	}
}

func (r *Router) Run() error {
	router := gin.Default()

	// Apply CORS middleware
	router.Use(CORSMiddleware())
	// router.SetTrustedProxies([]string{config.Envs.ClientHost, config.Envs.ClientHost1})

	// Setting up routes under baseURL
	api := router.Group(r.baseURL)
	{
		// Public routes
		publicRoutes := api.Group("/v1")
		for _, c := range r.controllers {
			c.RegisterPublic(publicRoutes)
		}

		// Protected routes
		protectedRoutes := api.Group("/v1")
		protectedRoutes.Use(authmiddleware.Authorize(r.jwtService, false))
		for _, c := range r.controllers {
			c.RegisterProtected(protectedRoutes)
		}

		// Privileged routes
		privilegedRoutes := api.Group("/v1")
		privilegedRoutes.Use(authmiddleware.Authorize(r.jwtService, true))
		for _, c := range r.controllers {
			c.RegisterPrivileged(privilegedRoutes)
		}
	}

	log.Println("Listening on", r.addr)
	return router.Run(r.addr)
}

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{config.Envs.ClientHost}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, accessToken, refreshToken, resetToken")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
