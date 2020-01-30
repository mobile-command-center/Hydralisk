package user

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type User struct {
	Id       string
	Password string
	Client   *http.Client
}

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

func (u *User) Logout(logout string) {
	req, _ := http.NewRequest("GET", logout, nil)

	u.setCookie(logout, req)

	resp, _ := u.Client.Get(logout)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}

func (u *User) setCookie(rawUrl string, req *http.Request) {
	parsedUrl, _ := url.Parse(rawUrl)
	for _, cookie := range u.Client.Jar.Cookies(parsedUrl) {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}
}

func (u *User) Register(register string, values url.Values) {
	req, _ := http.NewRequest("POST", register, strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	u.setCookie(register, req)
	resp, _ := u.Client.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}
