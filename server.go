package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Portnumber        float64
	IdleTimeout       string
	ReadHeaderTimeout string
	Tls               bool
	Tlspath           string
	Persist           bool
	Persistpath       string
}

func run(c Config) {
	// Spin up the goroutine that will own the data
	requestChannel := make(chan query)
	go dataOwner(requestChannel)

	// Set up the multiplexer
	mux := http.NewServeMux()
	setupHandlers(mux, requestChannel)

	// Configure the HTTP server
	idleDur, err := time.ParseDuration(c.IdleTimeout)
	if err != nil {
		fmt.Println(err)
		return
	}

	readHeadDur, err := time.ParseDuration(c.ReadHeaderTimeout)
	if err != nil {
		fmt.Println(err)
		return
	}

	srv := &http.Server{
		Addr:              ":" + strconv.Itoa(int(math.Floor(c.Portnumber))),
		Handler:           mux,
		IdleTimeout:       idleDur,
		ReadHeaderTimeout: readHeadDur,
	}

	// Run the server
	srv.ListenAndServe()
}
