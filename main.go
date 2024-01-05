package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, data := range movies {
		id := strconv.Itoa(data.Id)
		if id == params["id"] {
			json.NewEncoder(w).Encode(data)
			return
		}
	}
}

func deleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := params["id"]

	for index, data := range movies {
		if id == strconv.Itoa(data.Id) {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	const message = "{'message':'Error, Index already exist'}"
	for _, data := range movies {
		if data.Id == movie.Id {
			json.NewEncoder(w).Encode(message)
			return
		}
	}

	movies = append(movies, movie)
	fmt.Println(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id, _ = strconv.Atoi(params["id"])
	movies = append(movies[:movie.Id-1], movies[movie.Id:]...)
	movies = append(movies, movie)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: 1, Name: "DDLJ", Title: "blah blah", Director: &Director{FirstName: "Priyanshu", LastName: "Sharma"}})
	movies = append(movies, Movie{Id: 2, Name: "Koi Mil Gaya", Title: "blah blah", Director: &Director{FirstName: "Navya", LastName: "Sharma"}})
	fmt.Println(movies)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies/delete/{id}", deleteMovieById).Methods("DELETE")
	r.HandleFunc("/movies/createMovie", createMovie).Methods("POST")
	r.HandleFunc("/movies/updateMovie/{id}", updateMovie).Methods("PUT")

	http.ListenAndServe(":8080", r)
}
