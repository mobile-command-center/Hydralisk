package goods

import (
	"fmt"
	"log"

	"github.com/mobile-command-center/Hydralisk/client"
)

func skbParser(i *ItemInformation, c client.Client, company *CommCompany) int {
	fmt.Println("SKB PARSER")

	log.Printf("%+v\n", c)
	fieldPosition := 3

	if skbInternet(c) != "" {
		convert(i, fieldPosition, newSkbInternet(c, company))
		fieldPosition = fieldPosition + 1
	}

	if skbTv(c) != "" {
		convert(i, fieldPosition, newSkbTv(c, company))
		fieldPosition = fieldPosition + 1
	}

	if skbPhone(c) != "" {
		convert(i, fieldPosition, newSkbPhone(c, company))
		fieldPosition = fieldPosition + 1
	}
	return fieldPosition - 3
}

func skbInternet(c client.Client) string {
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

func skbWifi(c client.Client) string {
	item := map[string]string{
		"신청안함(기존 공유기 사용)": "와이파이신청안함",
		"신청":              "와이파이신청함",
	}
	return item[c.BoardWifi]
}

func skbTv(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)":     "",
		"B TV Lite - 211채널": "라이트(212CH)",
		"B TV All - 234채널":  "올(224CH)",
	}
	return item[c.BoardTv]
}

func skbTvAdd(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷단독시)":  "없음",
		"총 1대 설치(추가없음)": "총1대설치",
		"총 2대 설치":       "총2대설치",
		"총 3대 설치":       "총3대설치",
		"총 4대 설치":       "총4대설치",
	}
	return item[c.BoardTvAdd]
}

func skbSettop(c client.Client) string {
	item := map[string]string{
		"신청안함(인터넷 단독시)": "없음",
		"스마트티비셋탑":       "스마트3셋탑",
		"NUGU셋탑":        "누구2셋탑",
	}
	return item[c.BoardSettop]
}

func skbPhone(c client.Client) string {
	item := map[string]string{
		"신청안함N":              "",
		"일반(유선)전화 - 신규가입N":   "신규가입",
		"일반(유선)전화 - 번호이동Y":   "번호이동",
		"WiFi(무선)전화 - 신규가입N": "070신규",
		"WiFi(무선)전화 - 번호이동Y": "070번호이동",
	}
	return item[c.BoardTel]
}

func skbCombination(c client.Client) string {
	item := map[string]string{
		"결합없음":          "없음",
		"휴대폰 2대이상(온플랜)": "온프리(1회선)",
		"2,30년 이상(온할인)": "온가족할인(년수)",
		"SKT인터넷-인터넷 결합": "★패밀리★",
		"기타(요청란에 기입)":   "기타특이사항확인",
	}
	return item[c.Combination]
}

func skbGiftCardCode(c client.Client) string {
	item := map[string]string{
		"0": "",
		"1": "신세계 상품권",
		"3": "OK 캐시백",
		"2": "홈플러스 상품권 (프리)",
		"4": "SK 상품권 (플랜)",
	}
	return item[c.SpGiftCardCode]
}

func newSkbInternet(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["인터넷"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[skbInternet(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[skbCombination(c)],
		Service:         article.Service[skbWifi(c)],
		LineCount:       "1",
		GiftName:        skbGiftCardCode(c),
		GiftPrice:       "",
		GiftPaymentDay:  "",
		GiftPaymentType: "E",
		ReviewPrice:     "",
		TopGiftName:     "",
		TopGiftPrice:    "",
	}
}

func newSkbTv(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["BTV"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[skbTv(c)],
		Promise:         article.Promise["3년약정"],
		Sale:            article.Sale[skbSettop(c)],
		Service:         article.Service[skbTvAdd(c)],
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

func newSkbPhone(c client.Client, company *CommCompany) *DummyItem {
	article := company.Article["전화"]

	return &DummyItem{
		Item:            article.Article,
		Option:          article.Options[skbPhone(c)],
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
