package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

//Config 함수는 ERP 서버 정보를 갖는 구조체이다.
type Config struct {
	Endpoint    string `json:"endpoint"` //ERP URL endpoint
	LoginUrl    string `json:"login"`    //Login URL
	LogoutUrl   string `json:"logout"`   //Logout URL
	RegisterUrl string `json:"register"` //Register URL
	Id          string `json:"id"`       //Admin id
	Password    string `json:"password"` //Admin password
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
			Name:               goods.EncodeKorean([]byte("테스트1")),
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
