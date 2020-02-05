package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//Config 함수는 ERP 서버 정보를 갖는 구조체이다.
type Config struct {
	Endpoint    string `json:"endpoint"` //ERP URL endpoint
	LoginUrl    string `json:"login"`    //Login URL
	LogoutUrl   string `json:"logout"`   //Logout URL
	RegisterUrl string `json:"register"` //Register URL
	Id          string `json:"id"`       //Admin id
	Password    string `json:"password"` //Admin password
}

var (
	c = &Config{}
)

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(reqBody))
	membership := &goods.Membership{}
	if err := decoder.Decode(membership); err != nil {
		fmt.Fprintln(w, http.StatusBadRequest)
	} else {
		u := user.NewUser(c.Id, c.Password)

		formValue := url.Values{}

		encoder := schema.NewEncoder()
		encoder.SetAliasTag("form")

		_ = encoder.Encode(membership, formValue)

		u.Login(c.LoginUrl)
		u.Register(c.RegisterUrl, formValue)
		defer u.Logout(c.LogoutUrl)

		fmt.Fprintf(w, "%s", http.StatusOK)
	}
}

func init() {
	log.Println("Hydralisk ERP server starting...")
	log.Println("Loading access information...")
	conf, err := ioutil.ReadFile("info.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(conf, c); err != nil {
		panic(err)
	}
	log.Println("Initialization success...")
}

func main() {
	http.HandleFunc("/register", registrationHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
