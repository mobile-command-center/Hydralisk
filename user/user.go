package user

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	LoginLevel int = 1 + iota
	LogoutLevel
	RegisterLevel
)

type UserData struct {
	Data     bytes.Buffer
	Filename string
}

func NewUserData(data bytes.Buffer, filename string) *UserData {
	return &UserData{
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

func (u *User) Do(v url.Values, data UserData) (status int, err error) {
	status = http.StatusOK
	err = nil

	status, err = u.Login(u.url[LoginLevel])
	if err != nil {
		return
	}

	status, err = u.Register(u.url[RegisterLevel], v, data)
	if err != nil {
		return
	}

	status, err = u.Logout(u.url[LogoutLevel])
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

func (u *User) SetUrl(url map[int]string) {
	u.url = url
}

//Login 함수는 관리자 로그인 함수이다.
//로그인 성공시 세션정보를 획득한다.
func (u *User) Login(loginURL string) (int, error) {
	value := url.Values{
		"m_id":     {u.id},
		"m_passwd": {u.password},
	}

	resp, err := u.client.PostForm(loginURL, value)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	return http.StatusOK, nil
}

//Logout 함수는 관리자 로그아웃 함수이다.
func (u *User) Logout(logoutURL string) (int, error) {
	req, err := http.NewRequest("GET", logoutURL, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	u.setCookie(logoutURL, req)

	resp, err := u.client.Get(logoutURL)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	return http.StatusOK, nil
}

//setCookie 함수는 고객 등록시 필요한 쿠키정보를 설정하는 함수이다.
func (u *User) setCookie(rawURL string, req *http.Request) {
	parsedURL, _ := url.Parse(rawURL)
	for _, cookie := range u.client.Jar.Cookies(parsedURL) {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}
}

//Register 함수는 고객정보를 등록하는 함수이다.
//ERP시스템에서는 multipart 타입으로 POST 요청해야 한다.
func (u *User) Register(registerURL string, v url.Values, user UserData) (int, error) {
	b, c := makeMultiPart(v, user)
	req, err := http.NewRequest(http.MethodPost, registerURL, strings.NewReader(b))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req.Header.Add("Content-Type", c)
	u.setCookie(registerURL, req)
	resp, err := u.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	return http.StatusOK, nil
}
