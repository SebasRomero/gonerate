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

func initTopic(name string) string {
	capitalized := cases.Title(language.Und).String(name)
	return `package server
	
	type ` + capitalized + ` struct {
		yourFirstProperty   string
		yourSecondProperty 	string
	}
	
	var ` + name + `Array []*` + capitalized + ` = []*` + capitalized + `{}
	`
}

func initRoutes(name string) string {
	capitalized := cases.Title(language.Und).String(name)
	return `package server
		func initRoutes() {
			http.HandleFunc("/", index)
			http.HandleFunc("/` + name + `",func(w http.ResponseWriter, r *http.Request){
				switch r.Method {
				case http.MethodGet:
					get` + capitalized + `(w,r)
	
				case http.MethodPost:
					//addMethod
	
				case http.MethodDelete:
					//addMethod
				
				case http.MethodPut:
					//addMethod
				
				default: 
					w.WriteHeader(http.StatusMethodNotAllowed)
					fmt.Fprintf(w, "Method not allowed")
					return
				}
			})
		}
		`
}

func initHandlers(name string) string {
	capitalized := cases.Title(language.Und).String(name)
	return `package server

	func index(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
		fmt.Fprintf(w, "Hello there")
	}

	func get` + capitalized + `(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(` + name + `Array)
	}
	`
}

func initRoutes(name string) string {

	if len(name) >= 3 {
		modifiedString := name[:len(name)-3]
		return `package server
		func initRoutes() {
			http.HandleFunc("/", index)
			http.HandleFunc("/` + string(modifiedString) + `",func(w http.ResponseWriter, r *http.Request){
				switch r.Method {
				case http.MethodGet:
					//addMethod
	
				case http.MethodPost:
					//addMethod
	
				case http.MethodDelete:
					//addMethod
				
				case http.MethodPut:
					//addMethod
				
				default: 
					w.WriteHeader(http.StatusMethodNotAllowed)
					fmt.Fprintf(w, "Method not allowed")
					return
				}
			})
		}
		`
	}

	return ""
}
