package goods

import (
	"fmt"

	"github.com/mobile-command-center/Hydralisk/client"
)

func rentalParser(i *ItemInformation, c client.Client, company *CommCompany) int {
	fmt.Println("RENTAL PARSER")

	fieldPosition := 3
	convert(i, fieldPosition, newRentalEmpty(c, company))

	return 1
}

func newRentalEmpty(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["미기입"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options["미기입"],
		Promise:         article.Promise["미기입"],
		Sale:            article.Sale["미기입"],
		Service:         article.Service["미기입"],
		LineCount:       "1",
		GiftName:        "현금",
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "A",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}
