package mail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"log"
)

const (
	From    = "support@ajungweb.co.kr"
	CharSet = "UTF-8"
)

//EmailContent는 이메일 정보를 담는 구조체이다
type EmailContent struct {
	Recipient string //받는 사람
	Subject   string //제목
	Body      string //내용
}

func (e *EmailContent) SetRecipient(recipient string) {
	e.Recipient = recipient
}

func (e *EmailContent) SetSubject(subject string) {
	e.Subject = subject
}

func (e *EmailContent) SetBody(body string) {
	e.Body = body
}

//NewMail 함수는 EmailContent 구조체를 리턴한다
//에러발생시에만 이메일을 전송한다.
func NewMail(subject string, recipient string) *EmailContent {
	return &EmailContent{
		Subject:   subject,
		Recipient: recipient,
	}
}

func Send(e *EmailContent) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(e.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(e.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(e.Subject),
			},
		},
		Source: aws.String(From),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	log.Println("Email send to address: " + e.Recipient)
	log.Println(result)
}

const (
	CommonInfo              = `아래 요청된 원본데이터를 확인하시기 바랍니다.\n 지속적으로 에러 발생시 개발담당자에게 연락 부탁드립니다.\n`
	RequestBodyError        = `요청하신 데이터가 일치하지 않습니다.\n`
	MultiPartParsingError   = `고객이 입력한 데이터 로딩에 실패하였습니다.\n`
	PostFormDecodingError   = `데이터 변환에 실패하였습니다.\n`
	TemplateGenerationError = `원본 데이터 파일 생성에 실패하였습니다.\n`
	ItemConvertingError     = `아정통신 CMS 툴에 호환되는 데이터 생성에 실패하였습니다.\n`
	ErpDataEncodingError    = `아정통신 CMS 데이터 인코딩에 실패하였습니다.\n`
	ErpSystemError          = `아정통신 CMS 데이터 통신에 실패하였습니다.\n`
)
