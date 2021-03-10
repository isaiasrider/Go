package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main () {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://httpbin.org/get?search=teste", nil)
	post_req, err := http.NewRequest("POST","http://httpbin.org/post",strings.NewReader("teste"))
	resp, err := client.Do(req)
	resp_post, err := client.Do(post_req)
	if err != nil {
		log.Fatalln("Não foi possível realizar a chamada")
	}
	defer resp.Body.Close()
	defer resp_post.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Não foi possívl ler o conteúdo")
	}
	content_post, err := ioutil.ReadAll(resp_post.Body)
	if err != nil {
		log.Fatalln("não foi possível executar o post")
	}
	fmt.Println(string(content))
	fmt.Println(string("teste"))
	fmt.Println(string(content_post))
}