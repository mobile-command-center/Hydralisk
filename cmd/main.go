package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {

	u := &user.User{
		Id:       "",
		Password: "",
		Client: &http.Client{
			Jar: func() *cookiejar.Jar {
				c, _ := cookiejar.New(nil)
				return c
			}(),
		},
	}

	formValue := url.Values{}
	encoder := schema.NewEncoder()
	encoder.Encode(&goods.Membership{}, formValue)

	fmt.Println(formValue.Encode())

	u.Login("")

	u.Logout("")
}
