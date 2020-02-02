package goods

import (
	"bytes"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

//EncodeKorean 함수는 입력된 문자열을 Korean으로 인코딩하는 함수이다.
func EncodeKorean(b []byte) string {
	var buff bytes.Buffer
	wr := transform.NewWriter(&buff, korean.EUCKR.NewEncoder())
	_, _ = wr.Write(b)
	_ = wr.Close()
	return buff.String()
}

//GetCustomerType 함수는 입력된 고객타입을 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetCustomerType(s string) string {
	customer := map[string]string{
		"개인":  "C",
		"사업자": "O",
		"외국인": "F",
	}
	return customer[s]
}

//GetAuthType 함수는 입력된 고객인증 타입을 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetAuthType(s string) string {
	auth := map[string]string{
		"주민":    "J",
		"운전면허증": "M",
	}
	return auth[s]
}

//GetRegistrationCourse 함수는 입력된 접수 경로 타입을 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetRegistrationCourse(s string) string {
	course := map[string]string{
		"인바운드": "0003",
	}
	return course[s]
}

//GetGiveType 함수는 입력된 납부 정보를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetGiveType(s string) string {
	give := map[string]string{
		"계좌이체": "A",
		"카드납부": "B",
		"지로납부": "C",
		"합산청구": "D",
	}
	return give[s]
}

//GetRelationship 함수는 관계 여부를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetRelationship(s string) string {
	relationship := map[string]string{
		"본인":  "A",
		"배우자": "B",
		"가족":  "C",
	}
	return relationship[s]
}

//GetBilling 함수는 합산 청구를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetBilling(s string) string {
	billing := map[string]string{
		"전화비": "A",
		"인터넷": "B",
	}
	return billing[s]
}

//GetBank 함수는 은행 정보를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetBank(s string) string {
	bank := map[string]string{
		"경남은행":      "23",
		"광주은행":      "14",
		"교보증권":      "42",
		"국민은행":      "5",
		"기업은행":      "3",
		"농협":        "8",
		"대구은행":      "20",
		"대신증권":      "45",
		"대우증권":      "29",
		"도이치뱅크":     "54",
		"동부증권":      "46",
		"동양종합금융증권":  "56",
		"메리츠증권":     "39",
		"미래에셋증권":    "28",
		"미쓰비시도쿄UFJ": "57",
		"뱅크오브아메리카":  "55",
		"부국증권":      "44",
		"부산은행":      "21",
		"산림조합":      "43",
		"산업은행":      "2",
		"삼성증권":      "30",
		"새마을금고":     "15",
		"서울은행":      "12",
		"솔로몬투자증권":   "48",
		"수협중앙회":     "6",
		"스위스 저축은행":  "25",
		"신영증권":      "41",
		"신한금융투자":    "38",
		"신한은행":      "13",
		"신협":        "16",
		"씨티은행":      "22",
		"에스케이증권":    "35",
		"에이치엠씨증권":   "34",
		"외환은행":      "4",
		"우리은행":      "9",
		"우리투자증권":    "32",
		"우체국":       "18",
		"유진투자증권":    "40",
		"이트레이드증권":   "47",
		"전북은행":      "19",
		"제일은행":      "11",
		"제주은행":      "24",
		"조흥은행":      "10",
		"축협":        "26",
		"카카오뱅크":     "58",
		"키움증권":      "51",
		"하나대투증권":    "37",
		"하나은행":      "17",
		"하이투자증권":    "33",
		"한국수출입은행":   "7",
		"한국은행":      "1",
		"한국투자증권":    "31",
		"한화증권":      "36",
		"현대증권":      "27",
		"HSBC":      "52",
		"JP 모간":     "53",
		"K뱅크":       "59",
		"LIG투자증권":   "50",
		"NH투자증권":    "49",
	}
	return bank[s]
}

//GetCardCompany 함수는 카드사 정보를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetCardCompany(s string) string {
	card := map[string]string{
		"광주카드":     "14",
		"농협(BC)카드": "5",
		"롯데카드":     "12",
		"비씨카드":     "2",
		"산업은행카드":   "15",
		"삼성카드":     "6",
		"수협카드":     "16",
		"신한카드":     "11",
		"씨티카드":     "13",
		"외환카드":     "9",
		"우리카드":     "3",
		"전북은행카드":   "17",
		"제주은행카드":   "18",
		"하나(BC)카드": "4",
		"하나카드":     "7",
		"한미카드":     "8",
		"현대카드":     "10",
		"KB카드":     "1",
		"LG카드":     "19",
	}
	return card[s]
}

//GetCommCompany 함수는 통신사를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetCommCompany(s string) string {
	comm := map[string]string{
		"KT":     "724",
		"SKB":    "78",
		"SKT":    "735",
		"LG U+":  "80",
		"LG헬로비전": "740",
		"스카이라이프": "739",
		"렌탈":     "743",
	}
	return comm[s]
}

//GetGoodsCount 함수는 상품 갯수를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetGoodsCount(s string) string {
	good := map[string]string{
		"단독":  "1",
		"DPS": "2",
		"TPS": "3",
	}
	return good[s]
}

//GetGiftPaymentType 함수는 상품타입을 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetGiftPaymentType(s string) string {
	gift := map[string]string{
		"사은품 없음":     "E",
		"유치자 지급":     "A",
		"고객센터 지급":    "B",
		"본사상품권 + 현금": "F",
	}
	return gift[s]
}

//GetMoveCommCompany 함수는 번호이동 통신사를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func GetMoveCommCompany(s string) string {
	comm := map[string]string{
		"SK": "SK",
		"LG": "LG",
		"KT": "KT",
	}
	return comm[s]
}
