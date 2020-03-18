package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"

	"github.com/gorilla/schema"
	"github.com/mobile-command-center/Hydralisk/client"
	"github.com/mobile-command-center/Hydralisk/goods"
	"github.com/mobile-command-center/Hydralisk/user"
	"github.com/sirupsen/logrus"
)

//Config 함수는 ERP 서버 정보를 갖는 구조체이다.
type Config struct {
	Endpoint    string `json:"endpoint"` //ERP URL endpoint
	LoginURL    string `json:"login"`    //Login URL
	LogoutURL   string `json:"logout"`   //Logout URL
	RegisterURL string `json:"register"` //Register URL
	ID          string `json:"id"`       //Admin id
	Password    string `json:"password"` //Admin password
}

var (
	conf       = &Config{}
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
	log = logrus.New()
)

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Debug(string(reqBody))
	decoder := json.NewDecoder(bytes.NewBuffer(reqBody))

	c := &client.Client{}
	if err := decoder.Decode(c); err != nil {
		fmt.Fprintln(w, http.StatusBadRequest)
	}

	converter := goods.NewConverter(*c)

	if err := converter.Convert(membership); err != nil {
		fmt.Fprintln(w, http.StatusInternalServerError)
	}

	u := user.NewUser(conf.ID, conf.Password)
	formValue := url.Values{}

	encoder := schema.NewEncoder()
	encoder.SetAliasTag("form")
	if err := encoder.Encode(membership, formValue); err != nil {
		fmt.Fprintln(w, http.StatusInternalServerError)
	}

	t := template.Must(template.New("Ajung").Parse(client.RequestTmpl))
	var rawData bytes.Buffer
	err = t.Execute(&rawData, c)

	status, err := u.Login(conf.LoginURL)
	if err != nil {
		fmt.Fprintln(w, status)
	}
	defer u.Logout(conf.LogoutURL)

	log.Debug(formValue.Encode())
	status, err = u.Register(conf.RegisterURL, formValue, rawData)
	if err != nil {
		fmt.Fprintln(w, status)
	}
	fmt.Fprintln(w, http.StatusOK)
}

func init() {
	log.Info("Hydralisk ERP server starting...")
	log.Info("Loading access information...")
	confFile, err := ioutil.ReadFile("info.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(confFile, conf); err != nil {
		panic(err)
	}

	log.SetLevel(logrus.DebugLevel)
	membership.AdminInformation.Yuchi = conf.ID
	membership.AdminInformation.Jupsu = conf.ID

	log.Info("Initialization success...")
}

func main() {
	http.HandleFunc("/register", registrationHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
