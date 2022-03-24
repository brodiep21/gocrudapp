package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w Http.Responsewriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", Isbn: "438227", Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	movies = append(movies, Movie{
		ID: "2", Isbn: "422855", Title: "Interstellar", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})

	movies = append(movies, Movie{
		ID: "3", Isbn: "442311", Title: "Harry Potter", Director: &Director{Firstname: "Chris", Lastname: "Columbus"}})

	movies = append(movies, Movie{
		ID: "4", Isbn: "456990", Title: "Lord of the Rings: 				Fellowship of the Ring", Director: &Director{Firstname: "Peter", Lastname: "Jackson"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000:", r))
}
