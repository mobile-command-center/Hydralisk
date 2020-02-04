package user

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

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

type User struct {
	id       string       //관리자 ID
	password string       //관리자 Password
	client   *http.Client //ERP 통신 클라이언트
}

//Login 함수는 관리자 로그인 함수이다.
//로그인 성공시 세션정보를 획득한다.
func (u *User) Login(login string) {
	resp, _ := u.client.PostForm(login, url.Values{
		"m_id":     {u.id},
		"m_passwd": {u.password},
	})
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}

//Logout 함수는 관리자 로그아웃 함수이다.
func (u *User) Logout(logout string) {
	req, _ := http.NewRequest("GET", logout, nil)

	u.setCookie(logout, req)

	resp, _ := u.client.Get(logout)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}

//setCookie 함수는 고객 등록시 필요한 쿠키정보를 설정하는 함수이다.
func (u *User) setCookie(rawUrl string, req *http.Request) {
	parsedUrl, _ := url.Parse(rawUrl)
	for _, cookie := range u.client.Jar.Cookies(parsedUrl) {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}
}

//Register 함수는 고객정보를 등록하는 함수이다.
//ERP시스템에서는 multipart 타입으로 POST 요청해야 한다.
func (u *User) Register(register string, v url.Values) {
	b, c := makeMultiPart(v)
	req, _ := http.NewRequest(http.MethodPost, register, strings.NewReader(b))
	req.Header.Add("Content-Type", c)
	u.setCookie(register, req)
	resp, _ := u.client.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}
