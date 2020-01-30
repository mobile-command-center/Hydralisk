package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Config struct {
	Endpoint    string `json:"endpoint"`
	LoginUrl    string `json:"login"`
	LogoutUrl   string `json:"logout"`
	RegisterUrl string `json:"register"`
	Id          string `json:"id"`
	Password    string `json:"password"`
}

func main() {
	conf, err := ioutil.ReadFile("info.json")
	if err != nil {
		panic(err)
	}

	c := &Config{}
	if err := json.Unmarshal(conf, c); err != nil {
		panic(err)
	}

	u := &user.User{
		Id:       c.Id,
		Password: c.Password,
		Client: &http.Client{
			Jar: func() *cookiejar.Jar {
				c, _ := cookiejar.New(nil)
				return c
			}(),
		},
	}

	formValue := url.Values{}
	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	_ = encoder.Encode(&goods.Membership{}, formValue)

	fmt.Println(formValue.Encode())

	u.Login(c.LoginUrl)

	u.Logout(c.LogoutUrl)
}
