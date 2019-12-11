package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/drkchiloll/ex-rest-static/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var entry, static, port string
	flag.StringVar(&entry, "entry", ".index.html", "the entrypoint")
	flag.StringVar(&static, "static", ".", "the dir to serve from")
	flag.StringVar(&port, "port", "8000", "the port to listen on")
	flag.Parse()

	r := mux.NewRouter()

	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/users", handler.GetUsersHandler).Methods("GET")

	r.PathPrefix("/dist").Handler(http.FileServer(http.Dir(static)))
	r.PathPrefix("/").HandlerFunc(handler.IndexHandler(entry))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// func indexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, entrypoint)
// 	}
// 	return http.HandlerFunc(fn)
// }

// func getUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	data := map[string]interface{}{
// 		"id": "12345",
// 		"ts": time.Now().Format(time.RFC3339),
// 	}

// 	b, err := json.Marshal(data)
// 	if err != nil {
// 		http.Error(w, err.Error(), 400)
// 		return
// 	}
// 	w.Write(b)
// }
