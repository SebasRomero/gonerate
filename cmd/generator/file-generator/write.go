package filegenerator

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func initMainWriter() string {
	return `package main
	
	func main(){
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
	`
}

func initServerWriter() string {
	return `package server

	func New(addr string) *http.Server {

	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}
	`
}

func initTopic(name string) string { // This isn't working well
	capitalized := cases.Title(language.Und).String(name)
	return `package server
	
	type ` + capitalized + ` struct {
		yourFirstProperty   string
		yourSecondProperty 	string
	}
	
	var ` + name + `array []*` + capitalized + ` = []*` + capitalized + `{}
	`
}
