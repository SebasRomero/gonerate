package filegenerator

func initMainWriter() string {
	return `
	package main
	
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
	return `
	package server

	func New(addr string) *http.Server {

	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}
	`
}
