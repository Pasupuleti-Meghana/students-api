package main

import (
	"Pasupuleti-Meghana/students-api/config"
	"Pasupuleti-Meghana/students-api/internal/http/handlers/student"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	//load config
	cfg := config.MustLoad()

	//database setup

	//router setup
	router := http.NewServeMux()   //which function should handle which URL path. // create webserver, Handle requests from client, Send responses back
	fmt.Printf("Starting server on port %s..\n",cfg.HTTPServer.Address)

	router.HandleFunc("POST /students", student.New())


	//server setup
	//create the server instance
	server := http.Server {
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}

	
	fmt.Println("Server started successfully")

	// err :=  server.ListenAndServe()
	// if err != nil {
	// 	log.Fatalf("failed to start server: %s", err.Error())
	// }
	//After if server is started, it will keep running and listening for incoming requests until it is stopped or encounters an error.
	// so to handle the system shutdown error gracefully, we can use the os/signal package to listen for interrupt signals and shut down the server gracefully when such a signal is received.
	// so we are starting the server in a separate goroutine, and then we are listening for interrupt signals in the main goroutine. When an interrupt signal is received, we are shutting down the server gracefully using the Shutdown method of the http.Server struct.

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("Failed to start server", slog.String("error", err.Error()))
		}
	} ()

	<-done 

	slog.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shut down the server", slog.String("error", err.Error()))
	}

}