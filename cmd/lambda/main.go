package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
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

var (
	c = &Config{
		LoginUrl:    os.Getenv("LOGIN"),
		LogoutUrl:   os.Getenv("LOGOUT"),
		RegisterUrl: os.Getenv("REGISTRATION"),
		Id:          os.Getenv("ID"),
		Password:    os.Getenv("PASSWORD"),
	}
	membership = &goods.Membership{
		PaymentInformation: goods.PaymentInformation{
			AccountTransfer: goods.AccountTransfer{
				Bank: "0",
			},
			CreditCard: goods.CreditCard{
				Company: "0",
			},
		},
		ItemInformation: goods.ItemInformation{
			FirstItem: &goods.FirstItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
			SecondItem: &goods.SecondItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
			ThirdItem: &goods.ThirdItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
		},
	}
)

func Handler(ctx context.Context, client client.Client) error {
	converter := goods.NewConverter(client)
	if err := converter.Convert(membership); err != nil {
		return errors.New("Membership converting error")
	}

	u := user.NewUser(c.Id, c.Password)

	formValue := url.Values{}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	encoder.Encode(membership, formValue)

	u.Login(c.LoginUrl)
	defer u.Logout(c.LogoutUrl)
	u.Register(c.RegisterUrl, formValue)

	return nil
}

func init() {
	membership.AdminInformation.Yuchi = c.Id
	membership.AdminInformation.Jupsu = c.Id
}

func main() {
	lambda.Start(Handler)
}
