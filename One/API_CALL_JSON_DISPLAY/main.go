package main

// USING API TO FETCH DATA AND MARSHAL AND UNMARSHAL
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Joke struct {
	Id        int    `json:"id"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	res, err := http.Get("https://official-joke-api.appspot.com/random_ten")
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Println("Error1:", err)
	}
	defer res.Body.Close()

	var jokes []Joke

	DecodeErr := json.NewDecoder(res.Body).Decode(&jokes)
	if DecodeErr != nil {
		fmt.Println("Error3", DecodeErr)
	}

	for _, val := range jokes {
		fmt.Println(val.Setup)
		fmt.Println(val.Punchline)
		fmt.Println("------------------")
	}

}
