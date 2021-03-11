package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	if err != nil {
		log.Fatalln("nao foi possivel se autenticar")
	}
    req.SetBasicAuth("user","passw0rd")
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("nao foi possivel lero body")
	}

	fmt.Println(string(content))
	fmt.Println(res.StatusCode)

}
