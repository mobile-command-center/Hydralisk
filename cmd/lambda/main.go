package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"github.com/sirupsen/logrus"
)

//Config 구조체는 Ajung 툴 접속 정보를 갖는 구조체
type Config struct {
	LoginURL    string `json:"login"`    //Login URL
	LogoutURL   string `json:"logout"`   //Logout URL
	RegisterURL string `json:"register"` //Register URL
	ID          string `json:"id"`       //Admin id
	Password    string `json:"password"` //Admin password
}

var (
	conf = &Config{
		LoginURL:    os.Getenv("LOGIN"),
		LogoutURL:   os.Getenv("LOGOUT"),
		RegisterURL: os.Getenv("REGISTRATION"),
		ID:          os.Getenv("ID"),
		Password:    os.Getenv("PASSWORD"),
	}
	membership = emptyMembership()
	log        = logrus.New()
)

func emptyMembership() *goods.Membership {
	return &goods.Membership{
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
}

//Handler 는 Aws lambda 핸들러이다.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Content-Type"] = "text/plain; charset=utf-8"
	resp.StatusCode = http.StatusOK
	resp.IsBase64Encoded = false
	resp.Body = "가입 신청서가 성공적으로 전송되었습니다."

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

	c := &client.Client{}
	err = decoder.Decode(c, r.PostForm)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Post form decoding failed"
		return resp, err
	}

	t := template.Must(template.New("Ajung").Parse(client.RequestTmpl))
	var rawData bytes.Buffer
	err = t.Execute(&rawData, c)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Template generation failed"
		return resp, err
	}

	converter := goods.NewConverter(*c)
	err = converter.Convert(membership)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Item converting failed"
		return resp, err
	}

	u := user.NewUser(conf.ID, conf.Password)

	formValue := url.Values{}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	err = encoder.Encode(membership, formValue)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Form value for ERP encoding failed"
		return resp, err
	}

	log.Infof("%+v\n", membership)

	_, err = u.Login(conf.LoginURL)
	if err != nil {
		resp.StatusCode = http.StatusUnauthorized
		resp.Body = "Login failed"
		return resp, err
	}
	defer u.Logout(conf.LogoutURL)

	_, err = u.Register(conf.RegisterURL, formValue, rawData)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = "Data sending failed"
		return resp, err
	}

	membership = emptyMembership()
	return resp, nil
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	membership.AdminInformation.Yuchi = conf.ID
	membership.AdminInformation.Jupsu = conf.ID
}

func main() {
	lambda.Start(Handler)
}
