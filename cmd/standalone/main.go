package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"io/ioutil"
	"log"
	"net/http"
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

type MyEvent struct {
	Membership goods.Membership `json:"membership"`
}

var (
	c          = &Config{}
	membership = &goods.Membership{
		PaymentInformation: goods.PaymentInformation{
			AccountTransfer: goods.AccountTransfer{
				Bank: "0",
			},
			CreditCard: goods.CreditCard{
				Company: "0",
			},
		},
		ItemInformation: goods.ItemInformation{
			FirstItem: &goods.FirstItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
			SecondItem: &goods.SecondItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
			ThirdItem: &goods.ThirdItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
		},
	}
)

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(reqBody))

	client := &client.Client{}
	if err := decoder.Decode(client); err != nil {
		fmt.Fprintln(w, http.StatusBadRequest)
	}

	fmt.Printf("Client\n%+v\n", client)

	converter := goods.NewConverter(*client)

	if err := converter.Convert(membership); err != nil {
		fmt.Fprintln(w, http.StatusInternalServerError)
	}

	u := user.NewUser(c.Id, c.Password)
	formValue := url.Values{}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	encoder.Encode(membership, formValue)

	fmt.Printf("formValue\n%+v\n", formValue)

	u.Login(c.LoginUrl)
	defer u.Logout(c.LogoutUrl)
	u.Register(c.RegisterUrl, formValue)

	fmt.Fprintln(w, "%d", http.StatusOK)
}

func init() {
	log.Println("Hydralisk ERP server starting...")
	log.Println("Loading access information...")
	conf, err := ioutil.ReadFile("info.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(conf, c); err != nil {
		panic(err)
	}

	membership.AdminInformation.Yuchi = c.Id
	membership.AdminInformation.Jupsu = c.Id

	log.Println("Initialization success...")
}

func main() {
	http.HandleFunc("/register", registrationHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
