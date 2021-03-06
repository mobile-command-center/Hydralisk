package user

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	//LoginLevel 은 Login URL을 지정한다.
	LoginLevel = 1 + iota
	//LogoutLevel 은 Logout URL을 지정한다.
	LogoutLevel = 2
	//RegisterLevel 은 Register URL을 지정한다.
	RegisterLevel = 3
)

//Data 구조체는 RawData 정보를 갖는 구조체이다.
type Data struct {
	Data     bytes.Buffer
	Filename string
}

//NewData 함수는 UserData 객체를 생성하여 리턴한다.
func NewData(data bytes.Buffer, filename string) *Data {
	return &Data{
		Data:     data,
		Filename: filename,
	}
}

//NewUser 함수는 User 계정 정보를 가지고 User 객체를 만들어서 반환하는 함수이다.
func NewUser(id string, password string) *User {
	return &User{
		id:       id,
		password: password,
		client: &http.Client{
			Jar: func() *cookiejar.Jar {
				c, _ := cookiejar.New(nil)
				return c
			}(),
		},
	}
}

//Do 함수는 아정통신 CMS 시스템에 데이터를 등록하는 로직을 수행한다.
func (u *User) Do(v url.Values, data Data) (status int, err error) {
	status = http.StatusOK
	err = nil

	status, err = u.login(u.url[LoginLevel])
	if err != nil {
		return
	}

	status, err = u.register(u.url[RegisterLevel], v, data)
	if err != nil {
		return
	}

	status, err = u.logout(u.url[LogoutLevel])
	if err != nil {
		return
	}
	return
}

//User 구조체는 Ajung 시스템 접속 정보를 갖는 구조체이다.
type User struct {
	id       string         //관리자 ID
	password string         //관리자 Password
	client   *http.Client   //ERP 통신 클라이언트
	url      map[int]string //ERP url
}

//SetURL 함수는 아정통신 CMS URL을 설정하는 함수이다
func (u *User) SetURL(url map[int]string) {
	u.url = url
}

//Login 함수는 관리자 로그인 함수이다.
//로그인 성공시 세션정보를 획득한다.
func (u *User) login(loginURL string) (int, error) {
	value := url.Values{
		"m_id":     {u.id},
		"m_passwd": {u.password},
	}

	req, err := http.NewRequest(http.MethodPost, loginURL, bytes.NewBufferString(value.Encode()))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	parsedURL, _ := url.Parse(loginURL)
	log.Printf("%v\n", u.client.Jar.Cookies(parsedURL))

	resp, err := u.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	debugResponseBody(resp.StatusCode, body)

	return http.StatusOK, nil
}

//Logout 함수는 관리자 로그아웃 함수이다.
func (u *User) logout(logoutURL string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, logoutURL, nil)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	parsedURL, _ := url.Parse(logoutURL)
	log.Printf("%v\n", u.client.Jar.Cookies(parsedURL))

	resp, err := u.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	debugResponseBody(resp.StatusCode, body)

	return http.StatusOK, nil
}

//Register 함수는 고객정보를 등록하는 함수이다.
//ERP시스템에서는 multipart 타입으로 POST 요청해야 한다.
func (u *User) register(registerURL string, v url.Values, d Data) (int, error) {
	b, c := makeMultiPart(v, d)
	req, err := http.NewRequest(http.MethodPost, registerURL, strings.NewReader(b))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req.Header.Add("Content-Type", c)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6")

	log.Printf("data to cms\n%+v\n", v)

	parsedURL, _ := url.Parse(registerURL)
	log.Printf("%v\n", u.client.Jar.Cookies(parsedURL))

	resp, err := u.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	debugResponseBody(resp.StatusCode, body)

	return http.StatusOK, nil
}
