package goods

import (
	"fmt"
	"github.com/mobile-command-center/Hydralisk/client"
)

func skParser(i *ItemInformation, c client.Client, company *CommCompany) {
	fmt.Println("SK PARSER")
	fieldPosition := 3

	if skInternet(c) != "" {
		convert(i, fieldPosition, newSkInternet(c, company))
		fieldPosition = fieldPosition + 1
	}

	if skTv(c) != "" {
		convert(i, fieldPosition, newSkTv(c, company))
		fieldPosition = fieldPosition + 1
	}

	if skPhone(c) != "" {
		convert(i, fieldPosition, newSkPhone(c, company))
	}
}

func skInternet(c client.Client) string {
	item := map[string]string{
		"스마트 다이렉트(100M)":  "100M",
		"기가인터넷 라이트(500M)": "500M",
		"기가인터넷(1G)":       "1G",
	}
	if c.BoardInternet == "" {
		return ""
	}
	return item[c.BoardInternet]
}

func skWifi(c client.Client) string {
	item := map[string]string{
		"신청안함(기존 공유기 사용)": "",
		"신청":              "+WIFI",
	}
	return item[c.BoardWifi]
}

func skTv(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)":     "",
		"B TV Lite - 211채널": "라이트(212CH)",
		"B TV All - 234채널":  "올(224CH)",
	}
	return item[c.BoardTv]
}

func skTvAdd(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷단독시)":  "",
		"총 1대 설치(추가없음)": "1",
		"총 2대 설치":       "2",
		"총 3대 설치":       "3",
		"총 4대 설치":       "4",
	}
	return item[c.BoardTvAdd]
}

func skSettop(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)": "없음",
		"스마트티비셋탑":       "스마트3셋탑",
		"NUGU셋탑":        "누구2셋탑",
	}
	return item[c.BoardSettop]
}

func skPhone(c client.Client) string {
	item := map[string]string{
		"신청안함N":              "",
		"일반(유선)전화 - 신규가입N":   "신규가입",
		"일반(유선)전화 - 번호이동Y":   "번호이동",
		"WiFi(무선)전화 - 신규가입N": "070신규",
		"WiFi(무선)전화 - 번호이동Y": "070번호이동",
	}
	return item[c.BoardTel]
}

func skCombination(c client.Client) string {
	item := map[string]string{
		"신청안함":       "미결합",
		"1회선 결합":     "온프리(1회선)",
		"2회선 결합":     "온가족할인(년수)",
		"3회선 결합":     "온가족할인(년수)",
		"4회선 결합":     "온가족할인(년수)",
		"인터넷-인터넷 결합": "휴대폰직접결합",
	}
	return item[c.Combination]
}

func skGiftCardCode(c client.Client) string {
	item := map[string]string{
		"0": "",
		"1": "신세계 상품권",
		"3": "OK 캐시백",
		"2": "홈플러스 상품권 (프리)",
		"4": "SK 상품권 (플랜)",
	}
	return item[c.SpGiftCardCode]
}

func newSkInternet(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["인터넷"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[skInternet(c)+skWifi(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[skCombination(c)],
		Service:         article.Service[skWifi(c)],
		LineCount:       "1",
		GiftName:        skGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newSkTv(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["BTV"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[skTv(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[skSettop(c)],
		Service:         article.Service["없음"],
		LineCount:       skTvAdd(c),
		GiftName:        "",
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newSkPhone(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["전화"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[skPhone(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            "0",
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
