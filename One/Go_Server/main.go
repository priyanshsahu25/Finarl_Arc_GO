package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r* http.Request){
fmt.Fprintf(w,"Hello you are at Root page")
}

func handleHome (w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello you are at home page")
}

type User struct{
	Name string `json:"name"`
}
func createUser(w http.ResponseWriter,r*http.Request){

if r.Method!=http.MethodPost{
	http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
}

defer r.Body.Close()

var user User

  err2:= json.NewDecoder(r.Body).Decode(&user)
  if err2 != nil {
	http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	return
}

if user.Name==""{
	http.Error(w, "Name field is required", http.StatusBadRequest)
		return
}

fmt.Fprintf(w, "User created successfully: %s", user.Name)


}

func main() {
mux:=http.NewServeMux()
mux.HandleFunc("/",handleRoot)
mux.HandleFunc("/home",handleHome)
mux.HandleFunc("POST /user",createUser)
fmt.Println("Serer has been started")
err := http.ListenAndServe(":8800", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}



}