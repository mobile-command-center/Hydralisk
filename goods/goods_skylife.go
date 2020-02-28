package goods

import (
	"fmt"
	"github.com/mobile-command-center/Hydralisk/client"
)

func skylifeParser(i *ItemInformation, c client.Client, company *CommCompany) int {
	fmt.Println("SKYLIFE PARSER")
	fieldPosition := 3

	if skylifeInternet(c) != "" {
		convert(i, fieldPosition, newSkylifeInternet(c, company))
		fieldPosition = fieldPosition + 1
	}

	if skylifeTv(c) != "" {
		convert(i, fieldPosition, newSkylifeTv(c, company))
		fieldPosition = fieldPosition + 1
	}
	return fieldPosition - 3
}

func skylifeInternet(c client.Client) string {
	item := map[string]string{
		"올레인터넷(100M)":  "100M",
		"기가(200M)":     "200M",
		"기가 콤팩트(500M)": "1G",
		"기가(1G)":       "1G",
	}
	if c.BoardInternet == "" {
		return ""
	}
	return item[c.BoardInternet]
}

func skylifeTv(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)": "",
		"SKY GReen A+":  "Green A+(195CH)",
		"SKY Blue A+":   "Blue A+(218CH)",
	}
	return item[c.BoardTv]
}

func skylifeCombination(c client.Client) string {
	item := map[string]string{
		"신청안함": "없음",
		"홈결합":  "없음",
	}
	return item[c.Combination]
}

func skylifeGiftCardCode(c client.Client) string {
	item := map[string]string{
		"0": "",
		"1": "신세계 상품권",
		"2": "홈플러스 상품권",
	}
	return item[c.SpGiftCardCode]
}

func newSkylifeInternet(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["인터넷"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[skylifeInternet(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[skylifeCombination(c)],
		Service:         article.Service["없음"],
		LineCount:       "1",
		GiftName:        skylifeGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newSkylifeTv(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["IPTV"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[ktTv(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale["없음"],
		Service:         article.Service["없음"],
		LineCount:       "1",
		GiftName:        "",
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}
