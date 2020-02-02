package user

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
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

func makeMultiPart(v url.Values) (string, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.SetBoundary("ajnetbot")

	for key, r := range v {
		var fw io.Writer
		fw, _ = w.CreateFormField(key)
		for _, value := range r {
			_, _ = io.Copy(fw, strings.NewReader(value))
		}
	}
	w.Close()

	return b.String(), w.FormDataContentType()
}

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
