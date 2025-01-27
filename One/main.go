package main

// USING API TO FETCH DATA AND MARSHAL AND UNMARSHAL
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
res,err:=http.Get("https://official-joke-api.appspot.com/random_ten")
if err!=nil || res.StatusCode!= http.StatusOK  {
	fmt.Println("Error1:",err)
}
defer res.Body.Close()
body,readErr:=ioutil.ReadAll(res.Body)
if readErr!=nil{
	fmt.Println("Error2",readErr)
}else{
	fmt.Println("Body :",string(body))
}

}