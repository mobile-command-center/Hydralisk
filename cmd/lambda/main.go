package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"net/url"
	"os"
)

type Config struct {
	LoginUrl    string `json:"login"`    //Login URL
	LogoutUrl   string `json:"logout"`   //Logout URL
	RegisterUrl string `json:"register"` //Register URL
	Id          string `json:"id"`       //Admin id
	Password    string `json:"password"` //Admin password
}

type MyEvent struct {
	MemberShip goods.Membership `json:"membership"`
}

func Handler(ctx context.Context, myEvent MyEvent) error {
	c := Config{
		LoginUrl:    os.Getenv("LOGIN"),
		LogoutUrl:   os.Getenv("LOGOUT"),
		RegisterUrl: os.Getenv("REGISTRATION"),
		Id:          os.Getenv("ID"),
		Password:    os.Getenv("PASSWORD"),
	}

	u := user.NewUser(c.Id, c.Password)

	formValue := url.Values{}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")

	_ = encoder.Encode(myEvent.MemberShip, formValue)

	u.Login(c.LoginUrl)
	defer u.Logout(c.LogoutUrl)
	u.Register(c.RegisterUrl, formValue)

	return nil
}

func main() {
	lambda.Start(Handler)
}
