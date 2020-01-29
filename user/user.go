package user

import (
	"fmt"
	"net/http"
	"net/url"
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
	parsedUrl, _ := url.Parse(logout)
	for _, cookie := range u.Client.Jar.Cookies(parsedUrl) {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}

	resp, _ := u.Client.Get(logout)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
}
