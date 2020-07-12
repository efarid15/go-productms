package main

import (
	"context"
	"github.com/nicholasjackson/env"
	"gomicroservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)
var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address server")
func main() {
	_ = env.Parse()
	newlog := log.New(os.Stdout, "gomicroservice", log.LstdFlags)
    hp := handlers.NewProduct(newlog)

    sm := http.NewServeMux()
    sm.Handle("/products", hp)
	sm.Handle("/product/", hp)

	s := http.Server{
		Addr: *bindAddress,
		Handler: sm,
		ErrorLog: newlog,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	// start server
	go func() {
		newlog.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			newlog.Printf("Error starting server : %s\n\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.Shutdown(ctx)
}

