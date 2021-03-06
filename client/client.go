package client

//Client 구조체는 Ajung 가입페이지에서 입력된 폼 데이터를 관리하는 구조체이다.
type Client struct {
	Vendor                  string   `json:"vendor" form:"vendor"`                                         //통신사
	TelephoneCarrierMoveChk string   `json:"telephone_carrier_move_chk" form:"telephone_carrier_move_chk"` // 번호이동?
	Name                    string   `json:"c_name" form:"c_name"`                                         //가입자명
	Tel2Type                string   `json:"c_tel2_type" form:"c_tel2_type"`                               //가입자 휴대폰 통신사
	Tel21                   string   `json:"c_tel21" form:"c_tel21"`                                       //가입자 휴대폰 앞자리
	Tel22                   string   `json:"c_tel22" form:"c_tel22"`                                       //가입자 휴대폰 중간자리
	Tel23                   string   `json:"c_tel23" form:"c_tel23"`                                       //가입자 휴대폰 뒷자리
	Auth                    string   `json:"g_auth" form:"g_auth"`                                         //가입자 휴대폰 본인여부
	Tel31                   string   `json:"c_tel31" form:"c_tel31"`                                       //상담 받은 연락처 앞자리
	Tel32                   string   `json:"c_tel32" form:"c_tel32"`                                       //상담 받은 연락처 중간자리
	Tel33                   string   `json:"c_tel33" form:"c_tel33"`                                       //상담 받은 연락처 끝자리
	Tel11                   string   `json:"c_tel11" form:"c_tel11"`                                       //비상 연락처 앞자리
	Tel12                   string   `json:"c_tel12" form:"c_tel12"`                                       //비상 연락처 중간자리
	Tel13                   string   `json:"c_tel13" form:"c_tel13"`                                       //비상 연락처 뒷자리
	Email1                  string   `json:"c_email1" form:"c_email1"`                                     //이메일 주소 @앞자리
	Email2                  string   `json:"c_email2" form:"c_email2"`                                     //이메일 주소 @뒤 도메인
	ZipCode1                string   `json:"c_zipcode1" form:"c_zipcode1"`                                 //우편번호 앞자리
	ZipCode2                string   `json:"c_zipcode2" form:"c_zipcode2"`                                 //우편번호 뒷자리
	Address1                string   `json:"c_address" form:"c_address"`                                   //주소
	Address2                string   `json:"c_address2" form:"c_address2"`                                 //상세주소
	Jumin1                  string   `json:"c_jumin1" form:"c_jumin1"`                                     //주민번호 앞자리
	Jumin2                  string   `json:"c_jumin2" form:"c_jumin2"`                                     //주민번호 뒷자리
	Gender                  string   `json:"c_gender" form:"c_gender"`                                     //성별 (렌탈페이지)
	PaymentMethod           string   `json:"g_payment_method" form:"g_payment_method"`                     //납부 정보
	BankCd                  string   `json:"g_bank_cd" form:"g_bank_cd"`                                   //은행번호
	BankNo                  string   `json:"g_bank_no" form:"g_bank_no"`                                   //계좌번호
	BankHolder              string   `json:"g_bank_holder" form:"g_bank_holder"`                           //예금주
	CardCd                  string   `json:"g_card_cd" form:"g_card_cd"`                                   //카드 정보
	CardNo                  string   `json:"g_card_no" form:"g_card_no"`                                   //카드 번호
	CardGigan1              string   `json:"g_card_gigan" form:"g_card_gigan"`                             //카드 유효기간 (년)
	CardGigan2              string   `json:"g_card_gigan1" form:"g_card_gigan1"`                           //카드 유효기간 (월)
	CardHolder              string   `json:"g_card_holder" form:"g_card_holder"`                           //카드주 명
	SpBankCode              string   `json:"g_sp_bank_code" form:"g_sp_bank_code"`                         //사은품(현금) 지급은행 번호
	SpBankAccount           string   `json:"g_sp_bank_acount" form:"g_sp_bank_acount"`                     //사은품(현금) 지급은행 계좌번호
	SpBankHolder            string   `json:"g_sp_bank_holder" form:"g_sp_bank_holder"`                     //사은품(현금) 지급은행 예금주
	SpGiftCardCode          string   `json:"g_sp_giftcard_code" form:"g_sp_giftcard_code"`                 //사은품(상품권) 번호
	Bigo                    string   `json:"g_bigo" form:"g_bigo"`                                         // 요청사항
	BoardInternet           string   `json:"board_internet" form:"board_internet"`                         //상품정보(인터넷)
	BoardTv                 string   `json:"board_tv" form:"board_tv"`                                     //상품정보(TV)
	BoardTvAdd              string   `json:"board_tv_add" form:"board_tv_add"`                             //상품정보(TV추가설치)
	BoardTel                string   `json:"board_tel" form:"board_tel"`                                   //상품정보(전화)
	BoardSettop             string   `json:"board_settop" form:"board_settop"`                             //상품정보(셋탑박스)
	BoardWifi               string   `json:"board_wifi" form:"board_wifi"`                                 //상품정보(와이파이)
	CodePromise             string   `json:"g_code_promise" form:"g_code_promise"`                         //상품정보(약정 선택)
	MoveCompany             string   `json:"g_move_company" form:"g_move_company"`                         //기존 통신사
	MoveTel1                string   `json:"g_move_tel1" form:"g_move_tel1"`                               //기존 통신사 전화번호
	MoveAuth                string   `json:"g_move_auth" form:"g_move_auth"`                               //기존 통신사
	MoveNo                  string   `json:"g_move_no" form:"g_move_no"`                                   //기존 통신사
	Combination             string   `json:"p_combiation" form:"p_combiation"`                             //결합
	RentalVendor            []string `json:"p_vendor" form:"p_vendor"`                                     //Rental 회사
	RentalProduct           []string `json:"p_product" form:"p_product"`                                   //Rental 제품
	RentalProductName       string   `json:"p_product_name" form:"p_product_name"`                         //Rental 제품 이름
	RentalProductColor      string   `json:"p_product_color" form:"p_product_color"`                       //Rental 제품 색상
	RentalPromise           string   `json:"p_promise" form:"p_promise"`                                   //Rental 약정
	KtUser                  string   `json:"p_kt_user" form:"p_kt_user"`                                   //Skylife 옵션 (기존KT사용)
}

//RequestTmpl 는 CMS 시스템이나 EMAIL에 전송할 템플릿이다.
const RequestTmpl = `
아정통신 관리 시스템 요청 데이타 원본
통신사 : {{if .Vendor}} {{.Vendor}} {{else}} {{end}}
가입자명 : {{if .Name}} {{.Name}} {{else}} {{end}}
가입자 휴대폰 통신사 : {{if .Tel2Type}} {{.Tel2Type}} {{else}} {{end}}
가입자 휴대폰 : {{if and .Tel21 .Tel22 .Tel23}} {{.Tel21}}-{{.Tel22}}-{{.Tel23}} {{else}} {{end}}
가입자 휴대폰 본인 여부 : {{if .Auth}} {{.Auth}} {{else}} {{end}}
상담 받은 연락처 : {{if and .Tel31 .Tel32 .Tel33}} {{.Tel31}}-{{.Tel32}}-{{.Tel33}} {{else}} {{end}}
비상 연락처 : {{if and .Tel11 .Tel12 .Tel13}} {{.Tel11}}-{{.Tel12}}-{{.Tel13}} {{else}} {{end}}
이메일 :  {{if and .Email1 .Email2}} {{.Email1}}@{{.Email2}} {{else}} {{end}}
우편번호 : {{if and .ZipCode1 .ZipCode2}} {{.ZipCode1}}-{{.ZipCode2}} {{else}} {{end}}
주소 : {{if and .Address1 .Address2}} {{.Address1}} {{.Address2}} {{else}} {{end}}
주민번호 : {{if .Jumin1}} {{if .Jumin2}} {{.Jumin1}}-{{.Jumin2}} {{else}} {{.Jumin1}} {{end}} {{else}} {{end}}
성별 : {{if .Gender}} {{.Gender}} {{else}} {{end}}
납부 정보 : {{if .PaymentMethod}} {{.PaymentMethod}} {{else}} {{end}}
은행번호 : {{if .BankCd}} {{.BankCd}} {{else}} {{end}}
계좌번호 : {{if .BankNo}} {{.BankNo}} {{else}} {{end}}
예금주 : {{if .BankHolder}} {{.BankHolder}} {{else}} {{end}}
카드 정보 : {{if .CardCd}} {{.CardCd}} {{else}} {{end}}
카드 번호 : {{if .CardNo}} {{.CardNo}} {{else}} {{end}}
카드 유효 기간(MM/YY) : {{if and .CardGigan2 .CardGigan1}} {{.CardGigan2}}/{{.CardGigan1}} {{else}} {{end}}
카드주명 : {{if .CardHolder}} {{.CardHolder}} {{else}} {{end}}
사은품(현금) 지급은행 번호 : {{if .SpBankCode}} {{.SpBankCode}} {{else}} {{end}}
사은품(현금) 지급은행 계좌번호 : {{if .SpBankAccount}} {{.SpBankAccount}} {{else}} {{end}}
사은품(현금) 지급은행 예금주 : {{if .SpBankHolder}} {{.SpBankHolder}} {{else}} {{end}}
사은품(상품권) 번호 : {{if .SpGiftCardCode}} {{.SpGiftCardCode}} {{else}} {{end}}
요청사항 : {{if .Bigo}} {{.Bigo}} {{else}} {{end}}
상품정보(인터넷) : {{if .BoardInternet}} {{.BoardInternet}} {{else}} {{end}}
상품정보(TV) : {{if .BoardTv}} {{.BoardTv}} {{else}} {{end}}
상품정보(TV추가설치) : {{if .BoardTvAdd}} {{.BoardTvAdd}} {{else}} {{end}}
상품정보(전화) : {{if .BoardTel}} {{.BoardTel}} {{else}} {{end}}
상품정보(셋탑박스) : {{if .BoardSettop}} {{.BoardSettop}} {{else}} {{end}}
상품정보(와이파이) : {{if .BoardWifi}} {{.BoardWifi}} {{else}} {{end}}
상품정보(약정선택) : {{if .CodePromise}} {{.CodePromise}} {{else}} {{end}}
기존 통신사 : {{if .MoveCompany}} {{.MoveCompany}} {{else}} {{end}}
기존 통신사 전화번호 : {{if .MoveTel1}} {{.MoveTel1}} {{else}} {{end}}
번호이동 인증 : {{if .MoveAuth}} {{.MoveAuth}} {{else}} {{end}}
번호이동 인증번호 : {{if .MoveNo}} {{.MoveNo}} {{else}} {{end}}
결합 : {{if .Combination}} {{.Combination}} {{else}} {{end}}
렌탈회사 정보
{{if .RentalVendor}}
{{range .RentalVendor}}
	{{.}}
{{end}}
{{end}}
렌탈제품
{{if .RentalProduct}}
{{range .RentalProduct}}
	{{.}}
{{end}}
{{end}}
렌탈제품 이름 : {{if .RentalProductName}} {{.RentalProductName}} {{else}} {{end}}
렌탈제품 색상 : {{if .RentalProductColor}} {{.RentalProductColor}} {{else}} {{end}}
렌탈약정 : {{if .RentalPromise}} {{.RentalPromise}} {{else}} {{end}}
Skylife 옵션 (기존KT사용) : {{if .KtUser}} {{.KtUser}} {{else}} {{end}}

은행 코드
경남은행 : 23
광주은행 : 14
교보증권 : 42
국민은행 : 5
기업은행 : 3
농협 : 8
대구은행 : 20
대신증권 : 45
대우증권 : 29
도이치뱅크 : 54
동부증권 : 46
동양종합금융증권 : 56
메리츠증권 : 39
미래에셋증권 : 28
미쓰비시도쿄UFJ : 57
뱅크오브아메리카 : 55
부국증권 : 44
부산은행 : 21
산림조합 : 43
산업은행 : 2
삼성증권 : 30
새마을금고 : 15
서울은행 : 12
솔로몬투자증권 : 48
수협중앙회 : 6
스위스 저축은행 : 25
신영증권 : 41
신한금융투자 : 38
신한은행 : 13
신협 : 16
씨티은행 : 22
에스케이증권 : 35
에이치엠씨증권 : 34
외환은행 : 4
우리은행 : 9
우리투자증권 : 32
우체국 : 18
유진투자증권 : 40
이트레이드증권 : 47
전북은행 : 19
제일은행 : 11
제주은행 : 24
조흥은행 : 10
축협 : 26
카카오뱅크 : 58
키움증권 : 51
하나대투증권 : 37
하나은행 :17
하이투자증권 : 33
한국수출입은행 : 7
한국은행 : 1
한국투자증권 : 31
한화증권 : 36
현대증권 : 27
HSBC : 52
JP 모간 : 53
K뱅크 : 59
LIG투자증권 : 50
NH투자증권 : 49

카드 코드
광주카드 : 14
농협(BC)카드 : 5
롯데카드 : 12
비씨카드 : 2
산업은행카드 : 15
삼성카드 : 6
수협카드 : 16
신한카드 : 11
씨티카드 : 13
외환카드 : 9
우리카드 : 3
전북은행카드 : 17
제주은행카드 : 18
하나(BC)카드 : 4
하나카드 : 7
한미카드 : 8
현대카드 : 10
KB카드 : 1
LG카드 : 19
`
