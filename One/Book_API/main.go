package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	ID int `json:"book_id"`
	Name string `json:"name"`
	Author string `json:"author"`
	URL string `json:"URL"`
}



func main() {

	URL := "https://hapi-books.p.rapidapi.com/nominees/romance/2020"
	req,err:=http.NewRequest("GET",URL,nil)
	if err != nil {
		fmt.Println("Error sending Req", err)
	}
	req.Header.Add("x-rapidapi-key", "72530930a5mshcfef10a174d3faep1d3410jsn07acb80c4022")
	req.Header.Add("x-rapidapi-host", "hapi-books.p.rapidapi.com")
	

	client:=http.Client{}
	res,resErr:=client.Do(req)
	if resErr!=nil{
		fmt.Println("Error in getting Response",resErr)
	}
	fmt.Println("Status Code:", res.StatusCode)
	defer res.Body.Close()
	 Books:= make([]Book,0)
	decodeErr:=json.NewDecoder(res.Body).Decode(&Books)
	if decodeErr != nil {
		fmt.Println("Deocde Error:", decodeErr)
	}
	for _,book:= range Books{
		fmt.Println("ID :",book.ID)
		fmt.Println("Name :",book.Name)
		fmt.Println("Author :",book.Author)
		fmt.Println("ID :",book.URL)
		fmt.Println("----------------------")
	}

	// data,dataErr:=ioutil.ReadAll(res.Body)
	// if dataErr!=nil{
	// 	fmt.Println("Error from res body",dataErr)
	// }
	// fmt.Println(string(data))


	

	// res, err := http.Get(URL)
	// if err != nil {
	// 	fmt.Println("Error1", err)
	// } else if http.StatusOK != res.StatusCode {
	// 	fmt.Println("Status code :", res.StatusCode)
	// }
	// defer res.Body.Close()
	// data, readErr := ioutil.ReadAll(res.Body)
	// if readErr != nil {
	// 	fmt.Println("Read Error", readErr)
	// }
	// fmt.Println(string(data))
}
