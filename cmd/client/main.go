package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Aws string `json:"aws"`
}

var (
	c = &Config{}
)

func init() {
	conf, err := ioutil.ReadFile("aws.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(conf, c); err != nil {
		panic(err)
	}
}

func main() {
	conf, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Aws)
	req, _ := http.NewRequest("POST", c.Aws, bytes.NewReader(conf))

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	var b []byte
	res.Body.Read(b)
	fmt.Println(string(b))
}
