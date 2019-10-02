package main

import (
	"flag"
	"log"
	"time"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
)

var debug *bool

func main() {
	// Set flags
	var port = flag.String("p", "9000", "server port number")
	debug = flag.Bool("d", false, "debug mode")
	flag.Parse()

	log.Printf("Starting C3PO Rest API service on port %s\n", *port)

	if *debug {
		log.Printf("in debug mode\n")
	}

	// Setup server
	e := echo.New()
	e.Server.Addr = ":" + *port
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted, allows requests from any origin with GET, PUT, POST, DELETE or OPTIONS method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	// Set Handlers
	e.GET("/api/category", GetCategories)
	// e.GET("/api/ncms", GetNcms)
	// e.GET("/api/ncms/:id", GetNcmsByUser)
	// e.GET("/api/leads/trial", GetLeadsTrial)

	// Start server like a boss
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
