package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type GetResponse struct {
	ORIGIN string `json:"origin"`
	URL string     `json:"url"`
	HEADERS map[string]string `json:"headers"`
}

func (r *GetResponse) ToString() string {
	s:= fmt.Sprintf("THIS IS WHAT WE HAVE ON OUR REQUEST\n Origin IP: %s\nURL: %s", r.ORIGIN, r.URL)
	for k, v := range r.HEADERS {
		s += fmt.Sprintf("Headers: %s = %s\n", k,v)
	}
	return s
}
func main() {
	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to read response")
	}
	defer res.Body.Close()
	content , err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Unable to read the body")
	}

	respContent := GetResponse{}
	json.Unmarshal(content, &respContent)
    fmt.Println(respContent.ToString())
}




