package goods

import (
	"fmt"
	"github.com/mobile-command-center/Hydralisk/client"
)

func cjParser(i *ItemInformation, c client.Client, company *CommCompany) int {
	fmt.Println("CJ PARSER")
	fieldPosition := 3

	if cjInternet(c) != "" {
		convert(i, fieldPosition, newCjInternet(c, company))
		fieldPosition = fieldPosition + 1
	}

	if cjTv(c) != "" {
		convert(i, fieldPosition, newCjTv(c, company))
		fieldPosition = fieldPosition + 1
	}
	return fieldPosition - 3
}

func cjInternet(c client.Client) string {
	item := map[string]string{
		"광랜라이트+(100MB)": "광랜라이트(100M)",
		"광랜(160MB)":     "광랜(160M)",
		"기가라이트(500MB)":  "기가라이트(500M)",
		"플래티넘기가(1G)":    "플래티넘기가(1G)",
	}
	if c.BoardInternet == "" {
		return ""
	}
	return item[c.BoardInternet]
}

func cjWifi(c client.Client) string {
	item := map[string]string{
		"신청안함": "와이파이신청안함",
		"신청":   "와이파이신청",
	}
	return item[c.BoardWifi]
}

func cjTv(c client.Client) string {
	item := map[string]string{
		"베이직 UHD SMT":  "베이직(185CH)",
		"프리미엄 UHD SMT": "프리미엄(213CH)",
	}
	return item[c.BoardTv]
}

func cjTvAdd(c client.Client) string {
	item := map[string]string{
		"신청안함":    "추가설치없음",
		"베이직 HD":  "베이직HD 1대더추가설치",
		"이코노미 HD": "이코노미HD1대더추가설치",
	}
	return item[c.BoardTvAdd]
}

func cjCombination(c client.Client) string {
	item := map[string]string{
		"": "없음",
	}
	return item[c.Combination]
}

func cjGiftCardCode(c client.Client) string {
	item := map[string]string{
		"0": "",
		"1": "신세계 상품권",
		"2": "홈플러스 상품권",
	}
	return item[c.SpGiftCardCode]
}

func newCjInternet(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["인터넷"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[cjInternet(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[cjCombination(c)],
		Service:         article.Service[cjWifi(c)],
		LineCount:       "1",
		GiftName:        cjGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newCjTv(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["TV"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[cjTv(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale["UHD셋탑"],
		Service:         article.Service[cjTvAdd(c)],
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
