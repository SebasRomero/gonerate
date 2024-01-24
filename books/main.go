package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sebasromero/github.com/api-generator/books/server"
	"syscall"
)

func main() {
	ctx := context.Background()
	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8080")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server running on port", srv.Addr)
	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("Server stopped")
}
