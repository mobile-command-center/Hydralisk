package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	conf, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	res, err := http.Post("http://localhost:9090/register", "application/json", bytes.NewReader(conf))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	var b []byte
	res.Body.Read(b)
	fmt.Println(string(b))
}
