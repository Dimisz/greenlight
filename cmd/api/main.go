package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// version specifies the current version of the app.
// TODO: autogenerate when deploying
const version = "1.0.0"

// config struct holds all the config settings for the app.
type config struct {
	port int
	env  string
}

// application struct holds the dependencies for HTTP handlers, helpers and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// read the value of port and env command-line flags
	// into the config struct
	// defaults if no flags provided:
	// port 4000 and env "development"
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// initialize a new logger that writes messages to Stdout
	// prefixed with the current date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// declare an instance of application struct
	// containing the config struct and the logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// declare a HTTP server with timeout settings
	// which listens on the port specified in config
	// and uses our servemux
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start the HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
