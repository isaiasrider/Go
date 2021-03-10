package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main () {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalln("Não foi possível realizar a chamada")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Não foi possível ler o conteúdo")
	}
	fmt.Println(string(content))
}