package goods

import (
	"fmt"

	"github.com/mobile-command-center/Hydralisk/client"
)

func ktParser(i *ItemInformation, c client.Client, company *CommCompany) int {
	fmt.Println("KT PARSER")
	fieldPosition := 3

	if ktInternet(c) != "" {
		convert(i, fieldPosition, newKtInternet(c, company))
		fieldPosition = fieldPosition + 1
	}

	if ktTv(c) != "" {
		convert(i, fieldPosition, newKtTv(c, company))
		fieldPosition = fieldPosition + 1
	}

	if ktPhone(c) != "" {
		convert(i, fieldPosition, newKtPhone(c, company))
		fieldPosition = fieldPosition + 1
	}
	return fieldPosition - 3
}

func ktInternet(c client.Client) string {
	item := map[string]string{
		"올레인터넷(100M)":     "100M",
		"기가인터넷 콤팩트(500M)": "500M",
		"기가인터넷(1G)":       "1G",
	}
	if c.BoardInternet == "" {
		return ""
	}
	return item[c.BoardInternet]
}

func ktWifi(c client.Client) string {
	item := map[string]string{
		"신청안함(기존 공유기 사용)": "와이파이신청안함",
		"신청":              "와이파이신청함",
	}
	return item[c.BoardWifi]
}

func ktTv(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)":           "",
		"OTV슬림(구 OTV10) - 228채널":  "슬림(230CH)",
		"OTV라이트(구 OTV12) - 238채널": "라이트(240CH)",
		"OTV에센스(구 OTV15) - 261채널": "에센스(270CH)",
	}
	return item[c.BoardTv]
}

func ktTvAdd(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷단독시)":  "없음",
		"총 1대 설치(추가없음)": "총1대설치",
		"총 2대 설치":       "총2대설치",
		"총 3대 설치":       "총3대설치",
		"총 4대 설치":       "총4대설치",
	}
	return item[c.BoardTvAdd]
}

func ktSettop(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)": "없음",
		"일반셋탑":          "일반셋탑",
		"UHD셋탑":         "UHD셋탑",
		"기가지니2셋탑":       "기가지니2셋탑",
	}
	return item[c.BoardSettop]
}

func ktPhone(c client.Client) string {
	item := map[string]string{
		"신청안함N":              "",
		"일반(유선)전화 - 신규가입N":   "신규가입",
		"일반(유선)전화 - 번호이동Y":   "번호이동",
		"WiFi(무선)전화 - 신규가입N": "070신규",
		"WiFi(무선)전화 - 번호이동Y": "070번호이동",
	}
	return item[c.BoardTel]
}

func ktCombination(c client.Client) string {
	item := map[string]string{
		"신청안함":       "없음",
		"인터넷-인터넷 결합": "★패밀리★",
	}
	return item[c.Combination]
}

func ktGiftCardCode(c client.Client) string {
	item := map[string]string{
		"0": "",
		"1": "신세계 상품권",
		"2": "홈플러스 상품권",
	}
	return item[c.SpGiftCardCode]
}

func newKtInternet(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["인터넷"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[ktInternet(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[ktCombination(c)],
		Service:         article.Service[ktWifi(c)],
		LineCount:       "1",
		GiftName:        ktGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "E",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newKtTv(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["IPTV"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[ktTv(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[ktSettop(c)],
		Service:         article.Service[ktTvAdd(c)],
		LineCount:       "1",
		GiftName:        "",
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "E",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newKtPhone(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["전화"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[ktPhone(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            "0",
		Service:         article.Service["없음"],
		LineCount:       "1",
		GiftName:        "",
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "E",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}
