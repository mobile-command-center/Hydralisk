package goods

import (
	"fmt"

	"github.com/mobile-command-center/Hydralisk/client"
)

func lgParser(i *ItemInformation, c client.Client, company *CommCompany) int {
	fmt.Println("LG PARSER")
	fieldPosition := 3

	if lgInternet(c) != "" {
		convert(i, fieldPosition, newLgInternet(c, company))
		fieldPosition = fieldPosition + 1
	}

	if lgTv(c) != "" {
		convert(i, fieldPosition, newLgTv(c, company))
		fieldPosition = fieldPosition + 1
	}

	if lgPhone(c) != "" {
		convert(i, fieldPosition, newLgPhone(c, company))
		fieldPosition = fieldPosition + 1
	}
	return fieldPosition - 3
}

func lgInternet(c client.Client) string {
	item := map[string]string{
		"와이파이기본_광랜안심(100M)":   "100M+WIFI", //deprecated
		"와이파이기본_기가슬림안심(500M)": "500M+WIFI", //deprecated
		"와이파이기본_기가안심(1G)":     "1G+WIFI",   //deprecated
		"스마트 광랜안심(100M)":      "100M+WIFI",
		"스마트 기가슬림안심(500M)":    "500M+WIFI",
		"스마트 기가안심(1G)":        "1G+WIFI",
	}
	if c.BoardInternet == "" {
		return ""
	}
	return item[c.BoardInternet]
}

func lgTv(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)":           "",
		"TV베이직 - 183채널":           "베이직(185CH)",
		"TV프리미엄 - 223채널":          "프리미엄(224CH)",
		"TV VOD고급형 - 210채널":       "",
		"프리미엄 넷플릭스HD TV - 223채널":  "프리미엄 넷플릭스 HD(224CH)",
		"프리미엄 넷플릭스UHD TV - 223채널": "프리미엄 넷플릭스 UHD(224CH)",
	}
	return item[c.BoardTv]
}

func lgTvAdd(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷단독시)":  "없음",
		"총 1대 설치(추가없음)": "총1대설치",
		"총 2대 설치":       "총2대설치",
		"총 3대 설치":       "총3대설치",
		"총 4대 설치":       "총4대설치",
	}
	return item[c.BoardTvAdd]
}

func lgSettop(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)": "없음",
		"UHD셋탑":         "UHD셋탑",
	}
	return item[c.BoardSettop]
}

func lgPhone(c client.Client) string {
	item := map[string]string{
		"신청안함N":              "",
		"일반(유선)전화 - 신규가입N":   "[CPG]신규가입",
		"일반(유선)전화 - 번호이동Y":   "[CPG]번호이동",
		"WiFi(무선)전화 - 신규가입N": "[WIFI]신규가입",
		"WiFi(무선)전화 - 번호이동Y": "[WIFI]번호이동",
	}
	return item[c.BoardTel]
}

func lgCombination(c client.Client) string {
	item := map[string]string{
		"신청안함":       "없음",
		"인터넷-인터넷 결합": "★패밀리★",
	}
	return item[c.Combination]
}

func lgGiftCardCode(c client.Client) string {
	item := map[string]string{
		"0": "",
		"1": "신세계 상품권",
		"2": "홈플러스 상품권",
	}
	return item[c.SpGiftCardCode]
}

func lgGiftType(c client.Client) string {
	if c.SpGiftCardCode == "0" {
		return "E"
	}
	return "A"
}

func newLgInternet(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["인터넷"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[lgInternet(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[lgCombination(c)],
		Service:         article.Service["없음"],
		LineCount:       "1",
		GiftName:        lgGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: lgGiftType(c),
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newLgTv(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["IPTV"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[lgTv(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[lgSettop(c)],
		Service:         article.Service[lgTvAdd(c)],
		LineCount:       "1",
		GiftName:        lgGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: lgGiftType(c),
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newLgPhone(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["전화"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[lgPhone(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale["없음"],
		Service:         article.Service["없음"],
		LineCount:       "1",
		GiftName:        lgGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: lgGiftType(c),
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}
