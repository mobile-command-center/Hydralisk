package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Config struct {
	Aws string `json:"aws"`
}

var (
	c = &Config{}
)

func makeMultiPart(v url.Values) (string, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.SetBoundary("ajnetbot")

	for key, r := range v {
		var fw io.Writer
		fw, _ = w.CreateFormField(key)
		for _, value := range r {
			_, _ = io.Copy(fw, strings.NewReader(value))
		}
	}
	w.Close()

	return b.String(), w.FormDataContentType()
}

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
	conf, err := ioutil.ReadFile("client.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(conf))
	client := &client.Client{}
	if err := decoder.Decode(client); err != nil {
		log.Printf("decode err : %s\n", err)
	}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	formValue := url.Values{}
	if err := encoder.Encode(client, formValue); err != nil {
		log.Printf("encode err : %s\n", err)
	}

	s, ct := makeMultiPart(formValue)
	//fmt.Printf("string : \n%s\n", s)
	//fmt.Printf("contents type : %s\n", ct)

	req, err := http.NewRequest(http.MethodPost, c.Aws, strings.NewReader(s))
	if err != nil {
		log.Printf("request err : %s\n", err)
	}
	req.Header.Set("Content-Type", ct)
	cl := strconv.FormatInt(req.ContentLength, 10)
	fmt.Println(cl)
	req.Header.Set("Content-Length", cl)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Printf("request send err : %s\n", err)
	}
	defer res.Body.Close()
}
