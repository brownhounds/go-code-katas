package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Param  string `json:"param,omitempty"`
}

func handler_middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handler middleware...")
		next.ServeHTTP(w, r)
	})
}

func global_middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware...")
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	res := Response{
		Id:     5,
		Status: "OK",
	}

	json.NewEncoder(w).Encode(res)
}

func handler_route_params(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	res := Response{
		Id:     5,
		Status: "OK",
		Param:  r.PathValue("id"),
	}

	json.NewEncoder(w).Encode(res)
}

func Run() {
	mux := http.NewServeMux()
	port := strconv.Itoa(8000)

	mux.Handle("/", handler_middleware(http.HandlerFunc(handler)))
	mux.HandleFunc("/{id}", handler_route_params)

	log.Println("Server listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, global_middleware(mux)))
}
