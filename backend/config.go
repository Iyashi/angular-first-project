package main

import (
	"flag"
	"os"
)

var (
	developmentMode bool

	backendDomain = os.Getenv("BACKEND_DOMAIN")
	backendPort   = os.Getenv("BACKEND_PORT")

	frontendDomain = os.Getenv("FRONTEND_DOMAIN")
	frontendPort   = os.Getenv("FRONTEND_PORT")
)

func init() {
	flag.BoolVar(&developmentMode, "development", developmentMode, "Run in development mode")

	flag.StringVar(&backendDomain, "domain", backendDomain, "Domain at which the backend is accessible")
	flag.StringVar(&backendPort, "port", backendPort, "Port which backend listens on")

	flag.StringVar(&frontendDomain, "frontend-domain", frontendDomain, "Domain at which the frontend is accessible")
	flag.StringVar(&frontendPort, "frontend-port", frontendPort, "Port which frontend listens on")

	flag.Parse()
}
