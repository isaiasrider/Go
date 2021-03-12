package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/p@ssword", nil)
	if err != nil {
		log.Fatalln("não foi possível guardar a request")
	}
    req.SetBasicAuth("user","p@ssword")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("nao foi possivel ler o body")
	}

	fmt.Println(string(content))
	fmt.Println(res.StatusCode)

}
