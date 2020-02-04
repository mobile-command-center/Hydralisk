# Hydralisk
ERP  사이트 API v1

```go
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
```
