package main

import (
	"bytes"
	"encoding/base64"
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
	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
    enc.Write([]byte("user:passw0rd"))
	enc.Close()
	encondedCreds, err := buffer.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		log.Fatalln("failed to read encoded buffer")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s",encondedCreds))
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("nao foi possivel lero body")
	}

	fmt.Println(string(content))
	fmt.Println(res.StatusCode)

}
