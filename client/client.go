package client

type Client struct {
	Vendor                  string `json:"vendor" form:"vendor"`                                         //통신사
	TelephoneCarrierMoveChk string `json:"telephone_carrier_move_chk" form:"telephone_carrier_move_chk"` // 번호이동?
	Name                    string `json:"c_name" form:"c_name"`                                         //가입자명
	Tel2Type                string `json:"c_tel2_type" form:"c_tel2_type"`                               //가입자 휴대폰 통신사
	Tel21                   string `json:"c_tel21" form:"c_tel21"`                                       //가입자 휴대폰 앞자리
	Tel22                   string `json:"c_tel22" form:"c_tel22"`                                       //가입자 휴대폰 중간자리
	Tel23                   string `json:"c_tel23" form:"c_tel23"`                                       //가입자 휴대폰 뒷자리
	Auth                    string `json:"g_auth" form:"g_auth"`                                         //가입자 휴대폰 본인여부
	Tel31                   string `json:"c_tel31" form:"c_tel31"`                                       //상담 받은 연락처 앞자리
	Tel32                   string `json:"c_tel32" form:"c_tel32"`                                       //상담 받은 연락처 중간자리
	Tel33                   string `json:"c_tel33" form:"c_tel33"`                                       //상담 받은 연락처 끝자리
	Tel11                   string `json:"c_tel11" form:"c_tel11"`                                       //비상 연락처 앞자리
	Tel12                   string `json:"c_tel12" form:"c_tel12"`                                       //비상 연락처 중간자리
	Tel13                   string `json:"c_tel13" form:"c_tel13"`                                       //비상 연락처 뒷자리
	Email1                  string `json:"c_email1" form:"c_email1"`                                     //이메일 주소 @앞자리
	Email2                  string `json:"c_email2" form:"c_email2"`                                     //이메일 주소 @뒤 도메인
	ZipCode1                string `json:"c_zipcode1" form:"c_zipcode1"`                                 //우편번호 앞자리
	ZipCode2                string `json:"c_zipcode2" form:"c_zipcode2"`                                 //우편번호 뒷자리
	Address1                string `json:"c_address" form:"c_address"`                                   //주소
	Address2                string `json:"c_address2" form:"c_address2"`                                 //상세주소
	Jumin1                  string `json:"c_jumin1" form:"c_jumin1"`                                     //주민번호 앞자리
	Jumin2                  string `json:"c_jumin2" form:"c_jumin2"`                                     //주민번호 뒷자리
	PaymentMethod           string `json:"g_payment_method" form:"g_payment_method"`                     //납부 정보
	BankCd                  string `json:"g_bank_cd" form:"g_bank_cd"`                                   //은행번호
	BankNo                  string `json:"g_bank_no" form:"g_bank_no"`                                   //계좌번호
	BankHolder              string `json:"g_bank_holder" form:"g_bank_holder"`                           //예금주
	CardCd                  string `json:"g_card_cd" form:"g_card_cd"`                                   //카드 정보
	CardNo                  string `json:"g_card_no" form:"g_card_no"`                                   //카드 번호
	CardGigan1              string `json:"g_card_gigan" form:"g_card_gigan"`                             //카드 유효기간 (년)
	CardGigan2              string `json:"g_card_gigan1" form:"g_card_gigan1"`                           //카드 유효기간 (월)
	CardHolder              string `json:"g_card_holder" form:"g_card_holder"`                           //카드주 명
	SpBankCode              string `json:"g_sp_bank_code" form:"g_sp_bank_code"`                         //사은품(현금) 지급은행 번호
	SpBankAccount           string `json:"g_sp_bank_acount" form:"g_sp_bank_acount"`                     //사은품(현금) 지급은행 계좌번호
	SpBankHolder            string `json:"g_sp_bank_holder" form:"g_sp_bank_holder"`                     //사은품(현금) 지급은행 예금주
	SpGiftCardCode          string `json:"g_sp_giftcard_code" form:"g_sp_giftcard_code"`                 //사은품(상품권) 번호
	Bigo                    string `json:"g_bigo" form:"g_bigo"`                                         // 요청사항
	BoardInternet           string `json:"board_internet" form:"board_internet"`                         //상품정보(인터넷)
	BoardTv                 string `json:"board_tv" form:"board_tv"`                                     //상품정보(TV)
	BoardTvAdd              string `json:"board_tv_add" form:"board_tv_add"`                             //상품정보(TV추가설치)
	BoardTel                string `json:"board_tel" form:"board_tel"`                                   //상품정보(전화)
	BoardSettop             string `json:"board_settop" form:"board_settop"`                             //상품정보(셋탑박스)
	BoardWifi               string `json:"board_wifi" form:"board_wifi"`                                 //상품정보(와이파이)
	CodePromise             string `json:"g_code_promise" form:"g_code_promise"`                         //상품정보(약정 선택)
	MoveCompany             string `json:"g_move_company" form:"g_move_company"`                         //기존 통신사
	MoveTel1                string `json:"g_move_tel1" form:"g_move_tel1"`                               //기존 통신사 전화번호
	MoveAuth                string `json:"g_move_auth" form:"g_move_auth"`                               //기존 통신사
	MoveNo                  string `json:"g_move_no" form:"g_move_no"`                                   //기존 통신사
	Combination             string `json:"p_combiation" form:"p_combiation"`                             //결합
}
