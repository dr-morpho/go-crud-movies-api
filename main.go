package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`

}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, elem := range movies {
		if elem.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, elem := range movies {
		if elem.ID == params["id"]{
			json.NewEncoder(w).Encode(elem)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var singleMovie Movie
	_= json.NewDecoder(r.Body).Decode(&singleMovie)
	singleMovie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, singleMovie)
	json.NewEncoder(w).Encode(singleMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, elem := range movies {
		if elem.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...)
			var movieSingle Movie
			_=json.NewDecoder(r.Body).Decode(&movieSingle)
			movieSingle.ID = params["id"]
			movies = append(movies, movieSingle)
			json.NewEncoder(w).Encode(movieSingle)
			return
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Desperado", Director: &Director{FirstName: "Mark", LastName: "Vanderloo"}})
	movies = append(movies, Movie{ID: "2", Isbn: "555012", Title: "Dreams", Director: &Director{FirstName: "Kate", LastName: "Moss"}} )

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server start, port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}