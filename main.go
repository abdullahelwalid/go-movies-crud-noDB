package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"math/rand"
	"strconv"
)





type movie struct {
	
	ID string `json:"ID"`
	Isbn string `json:"Isbn"`
	Title string `json:"Title"`
	Director *director `json:"Director"`

}

type director struct {
	
	Firstname string `json:"Firstname"`
	Lastname string `json:"Lastname"`

}

var movies []movie 


func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	return
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}

	}
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Contenct-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	var movie movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main(){
	r := mux.NewRouter()
	movies = append(movies, movie{ID: "1", Isbn: "54456", Title: "SpiderMan 1", Director: &director{Firstname: "Someone", Lastname: "somebody"}})
	movies = append(movies, movie{ID: "2", Isbn: "68987", Title: "SpiderMan 2", Director: &director{Firstname: "Someone", Lastname: "somebody"}})
	movies = append(movies, movie{ID: "3", Isbn: "32165", Title: "SpiderMan 3", Director: &director{Firstname: "Someone", Lastname: "somebody"}})
	movies = append(movies, movie{ID: "4", Isbn: "46253", Title: "SpiderMan 4", Director: &director{Firstname: "Someone", Lastname: "somebody"}})

	r.HandleFunc("/", getMovies).Methods("GET")
	
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movie/{id}", deleteMovies).Methods("DELETE")

	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movie", createMovie).Methods("POST")


	fmt.Printf("Starting server on prot :8080\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}

