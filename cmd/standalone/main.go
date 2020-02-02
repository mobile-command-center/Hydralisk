package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Config struct {
	Endpoint    string `json:"endpoint"`
	LoginUrl    string `json:"login"`
	LogoutUrl   string `json:"logout"`
	RegisterUrl string `json:"register"`
	Id          string `json:"id"`
	Password    string `json:"password"`
}

func main() {
	conf, err := ioutil.ReadFile("info.json")
	if err != nil {
		panic(err)
	}

	c := &Config{}
	if err := json.Unmarshal(conf, c); err != nil {
		panic(err)
	}

	u := &user.User{
		Id:       c.Id,
		Password: c.Password,
		Client: &http.Client{
			Jar: func() *cookiejar.Jar {
				c, _ := cookiejar.New(nil)
				return c
			}(),
		},
	}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")

	member := goods.Membership{
		AdminInformation: goods.AdminInformation{
			Jupsu: c.Id,
			Yuchi: c.Id,
		},
		CustomerInformation: goods.CustomerInformation{
			Name: func() string {
				var buff bytes.Buffer
				wr := transform.NewWriter(&buff, korean.EUCKR.NewEncoder())
				_, _ = wr.Write([]byte("테스트1"))
				_ = wr.Close()
				return buff.String()
			}(),
			Phone:              "010-000-0000",
			CustomerType:       "C",
			RegistrationCourse: "0",
		},
		ItemInformation: goods.ItemInformation{
			Company:    "735",
			Location:   "4",
			GoodsCount: "1",
			FirstItem: goods.FirstItem{
				Item:            "736",
				Option:          "1716",
				Promise:         "84",
				Sale:            "301",
				Service:         "78",
				LineCount:       "1",
				GiftName:        "",
				GiftPrice:       "",
				GiftPaymentDay:  "",
				GiftPaymentType: "A",
				ReviewPrice:     "",
				TopGiftName:     "",
				TopGiftPrice:    "",
			},
			SecondItem: goods.SecondItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
			ThirdItem: goods.ThirdItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
		},
		PaymentInformation: goods.PaymentInformation{
			PaymentType:  "A",
			Relationship: "A",
			AccountTransfer: goods.AccountTransfer{
				Bank: "0",
			},
			CreditCard: goods.CreditCard{
				Company: "0",
			},
		},
		GiftAccount: goods.GiftAccount{
			Bank: "0",
		},
	}

	formValue := url.Values{}
	_ = encoder.Encode(&member, formValue)

	u.Login(c.LoginUrl)
	u.Register(c.RegisterUrl, formValue)
	defer u.Logout(c.LogoutUrl)
}
