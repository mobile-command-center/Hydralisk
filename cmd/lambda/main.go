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
	membership = goods.EmptyMembership()
)

func NewEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	return encoder
}

func NewDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.SetAliasTag("form")
	decoder.IgnoreUnknownKeys(true)
	return decoder
}

func MakeResponse() events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Content-Type"] = "text/plain; charset=utf-8"
	resp.StatusCode = http.StatusOK
	resp.IsBase64Encoded = false
	resp.Body = "가입 신청서가 성공적으로 전송되었습니다."

	return resp
}

func MakeRequest(req events.APIGatewayProxyRequest) (r http.Request, err error) {
	r = http.Request{}
	r.Header = make(map[string][]string)
	for k, v := range req.Headers {
		if k == "content-type" || k == "Content-Type" {
			r.Header.Set(k, v)
		}
	}

	var body []byte
	if req.IsBase64Encoded {
		body, err = base64.StdEncoding.DecodeString(req.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	} else {
		r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(req.Body)))
		err = nil
	}
	return
}

//Handler 는 Aws lambda 핸들러이다.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := MakeResponse()
	req, err := MakeRequest(request)
	if err != nil {
		resp.StatusCode = http.StatusForbidden
		resp.Body = "Could not read request body"
		return resp, err
	}

	maxMemory := int64(len(request.Body))
	err = req.ParseMultipartForm(maxMemory)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = "Parsing multi part form failed"
		return resp, err
	}

	c := &client.Client{}
	decoder := NewDecoder()
	err = decoder.Decode(c, req.PostForm)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = "Post form decoding failed"
		return resp, err
	}

	t := template.Must(template.New("Ajung").Parse(client.RequestTmpl))
	var data bytes.Buffer
	err = t.Execute(&data, c)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = "Template generation failed"
		return resp, err
	}

	converter := goods.NewConverter(*c)
	err = converter.Convert(membership)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = "Item converting failed"
		return resp, err
	}

	v := url.Values{}

	encoder := NewEncoder()
	err = encoder.Encode(membership, v)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = "Form value for ERP encoding failed"
		return resp, err
	}

	u := user.NewUser(conf.ID, conf.Password)
	u.SetUrl(map[int]string{
		user.LoginLevel:    conf.LoginURL,
		user.LogoutLevel:   conf.LogoutURL,
		user.RegisterLevel: conf.RegisterURL,
	})

	var status int
	status, err = u.Do(v, data)
	if err != nil {
		resp.StatusCode = status
		resp.Body = "ERP system internal error"
	}

	defer func() {
		membership = goods.EmptyMembership()
	}()

	return resp, nil
}

func init() {
	membership.AdminInformation.Yuchi = conf.ID
	membership.AdminInformation.Jupsu = conf.ID
}

func main() {
	lambda.Start(Handler)
}
