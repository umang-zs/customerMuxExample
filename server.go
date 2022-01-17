package muxExample

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter() //created a new instance of router
	router.HandleFunc("/", handler)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func handler(writer http.ResponseWriter, request *http.Request) {

	name := map[string]int{
		"Umang":  0,
		"Akshat": 1,
		"Jayesh": 2,
	}

	query := request.URL.Query().Get("name")
	if _, ok := name[query]; ok {
		writer.WriteHeader(http.StatusOK)

		_, err := writer.Write([]byte("Welcome " + query))
		if err != nil {
			log.Println("error in writing response")
		}

		return
	}

	writer.WriteHeader(http.StatusNotFound)

}
