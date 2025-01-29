package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
}


var Movies = []Movie{
	{1, "DDLJ", "Romance", "Priyansh"},
	{2, "KPMG", "Horror", "Sahu"},
	{3, "SAMSUNG", "COmedy", "JIVESH"},
}

func handleRoot ( w http.ResponseWriter, r *http.Request ){

	fmt.Fprintf(w,"Welcomr to Home")
}

func getMovies (w http.ResponseWriter, r *http.Request){
	if r.Method!= http.MethodGet{
		fmt.Println("Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)

}

func getMovie ( w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		fmt.Println("GET method allowed only",http.StatusMethodNotAllowed)
	}
	vars:=mux.Vars(r)
	movieId,err:=strconv.Atoi(vars["id"])
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		for _,movie:=range Movies{
			if movie.ID==movieId{
				fmt.Fprintf(w,"Movie is %s",movie.Name)
			}

		}
	}
}

func delteMovie(w http.ResponseWriter,r * http.Request){

	vars:=mux.Vars(r)
	movieID,err:=strconv.Atoi(vars["id"])
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		var temp [] Movie 
		for _,movie:=range Movies{
			if movie.ID!=movieID{
				temp = append(temp, movie)
			}
		}
		Movies=temp
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Movies)
		
	}

}

func startServer (){

	r := mux.NewRouter()
	r.HandleFunc("/",handleRoot).Methods("GET")
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}",delteMovie).Methods("DELETE")

	// r.HandleFunc("/movies",createMovie).Methods("POST")
	// r.HandleFunc("/movie/{id}",updateMovie).Methods("PUT")

	fmt.Println("Starting server at 8085 port")
	err:=http.ListenAndServe(":8085",r)
	if err!=nil{
		fmt.Println(err.Error())
	}
}

func main() {
	startServer()

}