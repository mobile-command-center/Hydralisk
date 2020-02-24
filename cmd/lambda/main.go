package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"github.com/sirupsen/logrus"
	"net/http"
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
	log = logrus.New()
)

type Response events.APIGatewayProxyResponse

func newResponse(status int) Response {
	return Response{
		StatusCode:      status,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "text/plain;charset-utf-8",
		},
	}
}

func Handler(ctx context.Context, client client.Client) (Response, error) {

	log.Debug(client)

	converter := goods.NewConverter(client)
	if err := converter.Convert(membership); err != nil {
		return newResponse(http.StatusInternalServerError), err
	}

	u := user.NewUser(c.Id, c.Password)

	formValue := url.Values{}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	err := encoder.Encode(membership, formValue)
	if err != nil {
		return newResponse(http.StatusInternalServerError), err
	}

	status, err := u.Login(c.LoginUrl)
	if err != nil {
		return newResponse(status), err
	}
	defer u.Logout(c.LogoutUrl)

	log.Debug(formValue.Encode())

	status, err = u.Register(c.RegisterUrl, formValue)
	if err != nil {
		return newResponse(status), err
	}

	return newResponse(http.StatusOK), nil
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	membership.AdminInformation.Yuchi = c.Id
	membership.AdminInformation.Jupsu = c.Id
}

func main() {
	lambda.Start(Handler)
}
