package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/mail"
	"github.com/mobile-command-center/Hydralisk/user"
)

//Config 구조체는 Ajung 툴 접속 정보를 갖는 구조체
type Config struct {
	LoginURL    string `json:"login"`     //Login URL
	LogoutURL   string `json:"logout"`    //Logout URL
	RegisterURL string `json:"register"`  //Register URL
	ID          string `json:"id"`        //Admin id
	Password    string `json:"password"`  //Admin password
	Recipient   string `json:"recipient"` //Email recipient
}

var (
	conf = &Config{
		LoginURL:    os.Getenv("LOGIN"),
		LogoutURL:   os.Getenv("LOGOUT"),
		RegisterURL: os.Getenv("REGISTRATION"),
		ID:          os.Getenv("ID"),
		Password:    os.Getenv("PASSWORD"),
		Recipient:   os.Getenv("RECIPIENT"),
	}
	membership = goods.EmptyMembership()
	userClient = user.NewUser(conf.ID, conf.Password)
)

const (
	BugReport          = "아정통신 버그리포트"
	RegistrationReport = "아정통신 가입신청서"
	Contect            = "관리자에게 연락바랍니다.\n"
)

//NewEncoder 함수는 웹페이지에서 전송되는 데이터를 구조체로
//변환하는 엔코더를 리턴한다.
func NewEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	return encoder
}

//NewDecoder 함수는 웹페이지에서 전송된 데이터 구조체를 FormValue로 변환하는
//디코더를 리턴한다.
func NewDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.SetAliasTag("form")
	decoder.IgnoreUnknownKeys(true)
	return decoder
}

//MakeResponse 함수는 APIGatewayProxyResponse 구조체를 리턴한다.
func MakeResponse() events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Content-Type"] = "text/plain; charset=utf-8"
	resp.StatusCode = http.StatusOK
	resp.IsBase64Encoded = false
	resp.Body = "가입 신청서가 성공적으로 전송되었습니다."

	return resp
}

//MakeRequest 함수는 APIGatewayProxyRequest 구조체를 리턴한다.
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
		resp.Body = mail.RequestBodyError + Contect
		m := mail.NewMail(BugReport, conf.Recipient)
		m.SetBody(mail.RequestBodyError + mail.CommonInfo + req.PostForm.Encode())
		mail.Send(m)
		return resp, err
	}

	maxMemory := int64(len(request.Body))
	err = req.ParseMultipartForm(maxMemory)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = mail.MultiPartParsingError + Contect
		m := mail.NewMail(BugReport, conf.Recipient)
		m.SetBody(mail.MultiPartParsingError + mail.CommonInfo + req.PostForm.Encode())
		mail.Send(m)
		return resp, err
	}

	log.Printf("Data from web\n%+v\n", req.Body)

	c := &client.Client{}
	decoder := NewDecoder()
	err = decoder.Decode(c, req.PostForm)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = mail.PostFormDecodingError + Contect
		m := mail.NewMail(BugReport, conf.Recipient)
		m.SetBody(mail.PostFormDecodingError + mail.CommonInfo + req.PostForm.Encode())
		mail.Send(m)
		return resp, err
	}

	t := template.Must(template.New("Ajung").Parse(client.RequestTmpl))
	var data bytes.Buffer
	err = t.Execute(&data, c)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = mail.TemplateGenerationError + Contect
		m := mail.NewMail(BugReport, conf.Recipient)
		m.SetBody(mail.TemplateGenerationError + mail.CommonInfo + req.PostForm.Encode())
		mail.Send(m)
		return resp, err
	}

	converter := goods.NewConverter(*c)
	err = converter.Convert(membership)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = mail.ItemConvertingError + Contect
		m := mail.NewMail(BugReport, conf.Recipient)
		m.SetBody(mail.ItemConvertingError + mail.CommonInfo + req.PostForm.Encode())
		mail.Send(m)
		return resp, err
	}

	v := url.Values{}

	encoder := NewEncoder()
	err = encoder.Encode(membership, v)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = mail.ErpDataEncodingError + Contect
		m := mail.NewMail(BugReport, conf.Recipient)
		m.SetBody(mail.ErpDataEncodingError + mail.CommonInfo + req.PostForm.Encode())
		mail.Send(m)
		return resp, err
	}

	userClient.SetURL(map[int]string{
		user.LoginLevel:    conf.LoginURL,
		user.LogoutLevel:   conf.LogoutURL,
		user.RegisterLevel: conf.RegisterURL,
	})

	var status int
	userData := user.NewUserData(data, membership.CustomerInformation.Name)
	status, err = userClient.Do(v, *userData)
	if err != nil {
		resp.StatusCode = status
		resp.Body = mail.ErpSystemError + Contect
		m := mail.NewMail(BugReport, conf.Recipient)
		m.SetBody(mail.ErpSystemError + mail.CommonInfo + req.PostForm.Encode())
		mail.Send(m)
		return resp, err
	}

	defer func() {
		membership = goods.EmptyMembership()
	}()

	m := mail.NewMail(RegistrationReport+" "+c.Name+" 고객님", conf.Recipient)
	m.SetBody(data.String())
	mail.Send(m)
	return resp, nil
}

func init() {
	membership.AdminInformation.Yuchi = conf.ID
	membership.AdminInformation.Jupsu = conf.ID
}

func main() {
	lambda.Start(Handler)
}
