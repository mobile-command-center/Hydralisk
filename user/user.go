package user

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type User struct {
	Id       string       //관리자 ID
	Password string       //관리자 Password
	Client   *http.Client //ERP 통신 클라이언트
}

//Login 함수는 관리자 로그인 함수이다.
//로그인 성공시 세션정보를 획득한다.
func (u *User) Login(login string) {
	resp, _ := u.Client.PostForm(login, url.Values{
		"m_id":     {u.Id},
		"m_passwd": {u.Password},
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

	resp, _ := u.Client.Get(logout)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}

//setCookie 함수는 고객 등록시 필요한 쿠키정보를 설정하는 함수이다.
func (u *User) setCookie(rawUrl string, req *http.Request) {
	parsedUrl, _ := url.Parse(rawUrl)
	for _, cookie := range u.Client.Jar.Cookies(parsedUrl) {
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
	resp, _ := u.Client.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}
