package goods

import (
	"reflect"

	"github.com/mobile-command-center/Hydralisk/client"
)

//NewConverter 구조체는 FormData를 Membership로 변환하는 Converter을 리턴한다.
func NewConverter(client client.Client) *Converter {
	return &Converter{c: client}
}

//Converter 구조체는 FormData 정보를 Membership로 변환하는 구조체이다.
type Converter struct {
	c client.Client
}

//Convert 함수는 Membership구조체로 Client 데이터를 변환한다.
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
		"SKB":     "SKB",
		"SKT":     "SKT",
		"KT":      "KT",
		"LG":      "LG U+",
		"skylife": "스카이라이프",
		"Hello":   "LG헬로비전",
		"rental":  "렌탈",
	}
	return comm[s]
}

func newServiceParser() map[string]func(*ItemInformation, client.Client, *CommCompany) int {
	return map[string]func(*ItemInformation, client.Client, *CommCompany) int{
		"SKT":     skParser,
		"KT":      ktParser,
		"LG":      lgParser,
		"skylife": skylifeParser,
		"Hello":   cjParser,
		"SKB":     skbParser,
		"rental":  rentalParser,
	}
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

//EmptyMembership 함수는 빈 Membership 구조체를 리턴한다.
func EmptyMembership() *Membership {
	return &Membership{
		PaymentInformation: PaymentInformation{
			AccountTransfer: AccountTransfer{
				Bank: "0",
			},
			CreditCard: CreditCard{
				Company: "0",
			},
		},
		ItemInformation: ItemInformation{
			FirstItem: &FirstItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
			SecondItem: &SecondItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
			ThirdItem: &ThirdItem{
				Item:      "0",
				Option:    "0",
				Promise:   "0",
				Sale:      "0",
				Service:   "0",
				LineCount: "1",
			},
		},
	}
}
