package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"github.com/sirupsen/logrus"
	"io/ioutil"
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

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.StatusCode = http.StatusOK

	r := http.Request{}
	r.Header = make(map[string][]string)
	for k, v := range request.Headers {
		if k == "content-type" || k == "Content-Type" {
			r.Header.Set(k, v)
		}
	}

	log.Info(request.Body)

	var body []byte
	var err error

	if request.IsBase64Encoded {
		body, err = base64.StdEncoding.DecodeString(request.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	} else {
		r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(request.Body)))
	}

	if err != nil {
		resp.StatusCode = 403
		resp.Body = "Could not read request body"
		return resp, err
	}

	maxMemory := int64(len(request.Body))
	err = r.ParseMultipartForm(maxMemory)
	if err != nil {
		log.Debug(err)
		resp.StatusCode = 500
		resp.Body = "Parsing multi part form failed"
		return resp, err
	}

	decoder := schema.NewDecoder()
	decoder.SetAliasTag("form")
	decoder.IgnoreUnknownKeys(true)

	client := &client.Client{}
	err = decoder.Decode(client, r.PostForm)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Post form decoding failed"
		return resp, err
	}

	log.Info(client)

	converter := goods.NewConverter(*client)
	err = converter.Convert(membership)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Item converting failed"
		return resp, err
	}

	u := user.NewUser(c.Id, c.Password)

	formValue := url.Values{}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	err = encoder.Encode(membership, formValue)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Form value for ERP encoding failed"
		return resp, err
	}

	log.Info(membership)

	_, err = u.Login(c.LoginUrl)
	if err != nil {
		resp.StatusCode = http.StatusUnauthorized
		resp.Body = "Login failed"
		return resp, err
	}
	defer u.Logout(c.LogoutUrl)

	_, err = u.Register(c.RegisterUrl, formValue)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Data sending failed"
		return resp, err
	}
	return resp, nil
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	membership.AdminInformation.Yuchi = c.Id
	membership.AdminInformation.Jupsu = c.Id
}

func main() {
	lambda.Start(Handler)
}
