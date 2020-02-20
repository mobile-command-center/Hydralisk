package goods

import (
	"github.com/mobile-command-center/Hydralisk/client"
	"reflect"
	"strconv"
)

func NewConverter(client client.Client) *Converter {
	return &Converter{c: client}
}

type Converter struct {
	c client.Client
}

func (c Converter) Convert(m *Membership) error {
	return m.Convert(c.c)
}

//GetCustomerType 함수는 입력된 고객타입을 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func getCustomerType(s string) string {
	customer := map[string]string{
		"개인":  "C",
		"사업자": "O",
		"외국인": "F",
	}
	return customer[s]
}

//GetRegistrationCourse 함수는 입력된 접수 경로 타입을 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func getRegistrationCourse(s string) string {
	course := map[string]string{
		"없음":   "0",
		"인바운드": "0003",
	}
	return course[s]
}

//GetGiveType 함수는 입력된 납부 정보를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func getGiveType(s string) string {
	give := map[string]string{
		"계좌이체": "A",
		"카드납부": "B",
		"지로납부": "C",
		"합산청구": "D",
	}
	return give[s]
}

//GetRelationship 함수는 관계 여부를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func getRelationship(s string) string {
	relationship := map[string]string{
		"본인":  "A",
		"배우자": "B",
		"가족":  "C",
	}
	return relationship[s]
}

func replaceCompanyName(s string) string {
	comm := map[string]string{
		"SK":      "SKT",
		"KT":      "KT",
		"LG":      "LG U+",
		"skylife": "스카이라이프",
		"CJ":      "LG헬로비전",
	}
	return comm[s]
}

func newServiceParser() map[string]func(*ItemInformation, client.Client, *CommCompany) {
	return map[string]func(*ItemInformation, client.Client, *CommCompany){
		"SK":      skParser,
		"KT":      ktParser,
		"LG":      lgParser,
		"skylife": skylifeParser,
		"CJ":      cjParser,
	}
}

func getGoodsCount(c client.Client) string {
	var goodCount int
	if c.BoardInternet != "" {
		goodCount++
	}

	if c.BoardTv != "" {
		goodCount++
	}

	if c.BoardTel != "" {
		goodCount++
	}

	return strconv.Itoa(goodCount)
}

//GetMoveCommCompany 함수는 번호이동 통신사를 ERP시스템에 맞는 값으로 치환해주는 함수이다.
func getMoveCommCompany(s string) string {
	comm := map[string]string{
		"SK":  "SK",
		"LG":  "LG",
		"KT":  "KT",
		"알뜰폰": "기타",
	}
	return comm[s]
}

func convert(i *ItemInformation, pos int, d interface{}) {
	r := reflect.Indirect(reflect.ValueOf(i))
	child := r.Field(pos)
	method := child.MethodByName("Convert")
	in := []reflect.Value{reflect.ValueOf(d)}
	method.Call(in)
}
