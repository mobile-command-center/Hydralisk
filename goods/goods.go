package goods

//AdminInformation 구조체는 고객등록시 필요한 관리자 ID이다. Cookie로 받아올 수 있다.
type AdminInformation struct {
	Jupsu string `form:"g_jupsu_m_id" json:"jupsu_id"` //접수 ID
	Yuchi string `form:"g_yuchi_m_id" json:"yuchi_id"` //유치 ID
}

//CustomerInformation 구조체는 고객정보탭의 정보를 저장하는 구조체이다.
type CustomerInformation struct {
	Name                       string `form:"c_name" json:"c_name"`                                 //고객명
	CustomerType               string `form:"c_type" json:"c_type"`                                 //고객 종류
	RegistrationNumber1        string `form:"c_jumin1" json:"c_jumin1,omitempty"`                   //주민번호 앞자리
	RegistrationNumber2        string `form:"c_jumin2" json:"c_jumin2,omitempty"`                   //주민번호 뒷자리
	BusinessRegistrationNumber string `form:"c_office_no" json:"c_office_no,omitempty"`             //사업자번호
	BusinessOwenerName         string `form:"c_office_name" json:"c_office_name,omitempty"`         //대표자
	Zipcode                    string `form:"c_zipcode1" json:"c_zipcode1,omitempty"`               //우편번호
	Address                    string `form:"c_address" json:"c_address,omitempty"`                 //주소
	Telephone                  string `form:"c_tel1" json:"c_tel1,omitempty"`                       //전화
	Phone                      string `form:"c_tel2" json:"c_tel2,omitempty"`                       //휴대전화
	Email                      string `form:"c_email" json:"c_email,omitempty"`                     //이메일
	AuthType                   string `form:"c_auth_chk" json:"c_auth_chk,omitempty"`               //고객인증 타입
	AuthData                   string `form:"c_jumin_date" json:"c_jumin_date,omitempty"`           //고객인증 데이터
	RegistrationCourse         string `form:"g_code_course_idx" json:"g_code_course_idx,omitempty"` //접수 경로
}

func (c *CustomerInformation) Validate() error {
	return nil
}

//AccountTransfer 구조체는 계좌이체 정보를 저장하는 구조체이다.
type AccountTransfer struct {
	Bank                string `form:"c_bank_cd" json:"c_bank_cd,omitempty"`         //지급은행
	Account             string `form:"c_bank_no" json:"c_bank_no,omitempty"`         //계좌번호
	OwenerName          string `form:"c_bank_name" json:"c_bank_name,omitempty"`     //예금주
	RegistrationNumber1 string `form:"c_bank_jumin1" json:"c_bank_jumin1,omitempty"` //예금주 주민번호 앞자리
	RegistrationNumber2 string `form:"c_bank_jumin2" json:"c_bank_jumin2,omitempty"` //예금주 주민번호 뒷자리
}

func (a *AccountTransfer) Validate() error {
	return nil
}

//CreditCard 구조체는 신용카드 정보를 저장하는 구조체이다.
type CreditCard struct {
	Company             string `form:"c_card_cd" json:"c_card_cd,omitempty"`         //카드사
	Number              string `form:"g_card_no" json:"g_card_no,omitempty"`         //카드번호
	ExpiredYear         string `form:"g_card_gigan" json:"g_card_gigan,omitempty"`   //카드 유효기간(년)
	ExpiredMonth        string `form:"g_card_gigan1" json:"g_card_gigan1,omitempty"` //카드 유효기간(월)
	OwenerName          string `form:"g_card_name" json:"g_card_name,omitempty"`     //카드주 명
	RegistrationNumber1 string `form:"c_card_jumin1" json:"c_card_jumin1,omitempty"` //카드주 주민번호 앞자리
	RegistrationNumber2 string `form:"c_card_jumin2" json:"c_card_jumin2,omitempty"` //카드주 주민번호 뒷자리
}

func (c *CreditCard) Validate() error {
	return nil
}

//Billing 구조체는 지로납부 정보를 저장하는 구조체이다.
//현재 지로납부에 대해서는 저장해야할 데이터가 없다.
type Billing struct{}

//Summing 구조체는 합산청구 정보를 저장하는 구조체이다.
type Summing struct {
	Billing  string `form:"g_sum_money_chk" json:"g_sum_money_chk,omitempty"` //합산 청구
	Comments string `form:"g_sum_money_txt" json:"g_sum_money_txt,omitempty"` //합산 비고
}

func (s *Summing) Validate() error {
	return nil
}

//PaymentInformation 구조체는 납부 정보를 저장하는 구조체이다.
type PaymentInformation struct {
	PaymentType     string          `form:"g_give_type" json:"g_give_type,omitempty"` //납부 정보
	Relationship    string          `form:"g_give_chk" json:"g_give_chk,omitempty"`   //관계 여부
	AccountTransfer AccountTransfer `json:"account_transfer,omitempty"`               //계좌이체
	CreditCard      CreditCard      `json:"credit,omitempty"`                         //카드납부
	Billing         Billing         `json:"billing,omitempty"`                        //지로납부
	Summing         Summing         `json:"summing,omitempty"`                        // 합산청구
}

func (p *PaymentInformation) Validate() error {
	return nil
}

//GiftAccount 구조체는 사은품 계좌 정보를 저장하는 구조체이다.
type GiftAccount struct {
	Bank                string `form:"g_sp_bank_code" json:"g_sp_bank_code,omitempty"`       //지급은행
	Account             string `form:"g_sp_bank_account" json:"g_sp_bank_account,omitempty"` //계좌번호
	OwenerName          string `form:"g_sp_bank_name" json:"g_sp_bank_name,omitempty"`       //예금주
	RegistrationNumber1 string `form:"g_sp_bank_jumin1" json:"g_sp_bank_jumin1,omitempty"`   //예금주 주민번호 앞자리
	RegistrationNumber2 string `form:"g_sp_bank_jumin2" json:"g_sp_bank_jumin2,omitempty"`   //예금주 주민번호 뒷자리
}

//FirstItem 구조체는 첫번째 사은품 정보를 저장하는 구조체이다.
type FirstItem struct {
	Item            string `form:"g_article_idx2" json:"g_article_idx2,omitempty"`   //상품 선택
	Option          string `form:"g_option_idx" json:"g_option_idx,omitempty"`       //옵션 선택
	Promise         string `form:"g_code_promise1" json:"g_code_promise1,omitempty"` //약정 선택
	Sale            string `form:"g_code_sale" json:"g_code_sale,omitempty"`         //할인탭 선택
	Service         string `form:"g_code_service1" json:"g_code_service1,omitempty"` //부가서비스
	LineCount       string `form:"g_article_cnt1" json:"g_article_cnt1,omitempty"`   //설치 회선
	GiftName        string `form:"g_sp_name" json:"g_sp_name,omitempty"`             //사은품명
	GiftPrice       string `form:"g_sp_price" json:"g_sp_price,omitempty"`           //사은품 가격
	GiftPaymentDay  string `form:"g_sp_date_start" json:"g_sp_date_start,omitempty"` //사은품 지급 날짜
	GiftPaymentType string `form:"g_sp_give_type" json:"g_sp_give_type,omitempty"`   //사은품 지급자 선택
	ReviewPrice     string `form:"g_sp_price_add1" json:"g_sp_price_add1,omitempty"` //후기금액
	TopGiftName     string `form:"g_sp_name_top1" json:"g_sp_name_top1,omitempty"`   //본사 사은품명
	TopGiftPrice    string `form:"g_sp_price_top1" json:"g_sp_price_top1,omitempty"` //본사 사은품 가격
}

//SecondItem 구조체는 번째 사은품 정보를 저장하는 구조체이다.
type SecondItem struct {
	Item            string `form:"g_article_idx2_1" json:"g_article_idx2_1,omitempty"` //상품 선택
	Option          string `form:"g_option_idx_1" json:"g_option_idx_1,omitempty"`     //옵션 선택
	Promise         string `form:"g_code_promise2" json:"g_code_promise2,omitempty"`   //약정 선택
	Sale            string `form:"g_code_sale1" json:"g_code_sale1,omitempty"`         //할인탭 선택
	Service         string `form:"g_code_service2" json:"g_code_service2,omitempty"`   //부가서비스
	LineCount       string `form:"g_article_cnt2" json:"g_article_cnt2,omitempty"`     //설치 회선
	GiftName        string `form:"g_sp_name1" json:"g_sp_name1,omitempty"`             //사은품명
	GiftPrice       string `form:"g_sp_price1" json:"g_sp_price1,omitempty"`           //사은품 가격
	GiftPaymentDay  string `form:"g_sp_date_start1" json:"g_sp_date_start1,omitempty"` //사은품 지급 날짜
	GiftPaymentType string `form:"g_sp_give_type1" json:"g_sp_give_type1,omitempty"`   //사은품 지급자 선택
	ReviewPrice     string `form:"g_sp_price_add2" json:"g_sp_price_add2,omitempty"`   //후기금액
	TopGiftName     string `form:"g_sp_name_top2" json:"g_sp_name_top2,omitempty"`     //본사 사은품명
	TopGiftPrice    string `form:"g_sp_price_top2" json:"g_sp_price_top2,omitempty"`   //본사 사은품 가격
}

//ThirdItem 구조체는 세번째 사은품 정보를 저장하는 구조체이다.
type ThirdItem struct {
	Item            string `form:"g_article_idx2_2" json:"g_article_idx2_2,omitempty"` //상품 선택
	Option          string `form:"g_option_idx_2" json:"g_option_idx_2,omitempty"`     //옵션 선택
	Promise         string `form:"g_code_promise3" json:"g_code_promise3,omitempty"`   //약정 선택
	Sale            string `form:"g_code_sale2" json:"g_code_sale2,omitempty"`         //할인탭 선택
	Service         string `form:"g_code_service3" json:"g_code_service3,omitempty"`   //부가서비스
	LineCount       string `form:"g_article_cnt3" json:"g_article_cnt3,omitempty"`     //설치 회선
	GiftName        string `form:"g_sp_name2" json:"g_sp_name2,omitempty"`             //사은품명
	GiftPrice       string `form:"g_sp_price2" json:"g_sp_price2,omitempty"`           //사은품 가격
	GiftPaymentDay  string `form:"g_sp_date_start2" json:"g_sp_date_start2,omitempty"` //사은품 지급 날짜
	GiftPaymentType string `form:"g_sp_give_type2" json:"g_sp_give_type2,omitempty"`   //사은품 지급자 선택
	ReviewPrice     string `form:"g_sp_price_add3" json:"g_sp_price_add3,omitempty"`   //후기금액
	TopGiftName     string `form:"g_sp_name_top3" json:"g_sp_name_top3,omitempty"`     //본사 사은품명
	TopGiftPrice    string `form:"g_sp_price_top3" json:"g_sp_price_top3,omitempty"`   //본사 사은품 가격
}

//ItemInformation 구조체는 상품 정보를 저장하는 구조체이다.
type ItemInformation struct {
	Company    string     `form:"g_article_idx1" json:"g_article_idx1,omitempty"` // 회사 선택
	Location   string     `form:"g_code_area" json:"g_code_area,omitempty"`       //지역 선택
	GoodsCount string     `form:"goods_cnt" json:"goods_cnt,omitempty"`           //상품 갯수
	FirstItem  FirstItem  `json:"first_item,omitempty"`
	SecondItem SecondItem `json:"second_item,omitempty"`
	ThirdItem  ThirdItem  `json:"third_item,omitempty"`
}

//NumberMoving 구조체는 번호이동 정보를 저장하는 구조체이다.
type NumberMoving struct {
	Telephone           string `form:"g_move_tel1" json:"g_move_tel1,omitempty"`       //전화번호
	Name                string `form:"g_move_name" json:"g_move_name,omitempty"`       //명의자
	RegistrationNumber1 string `form:"g_move_jumin1" json:"g_move_jumin1,omitempty"`   //주민번호 앞자리
	RegistrationNumber2 string `form:"g_move_jumin2" json:"g_move_jumin2,omitempty"`   //주민번호 뒷자리
	Phone               string `form:"g_move_tel2" json:"g_move_tel2,omitempty"`       //연락처
	Date                string `form:"g_move_date" json:"g_move_date,omitempty"`       //발급 일자
	PreviousCompany     string `form:"g_move_company" json:"g_move_company,omitempty"` //기존통신사
	Memo                string `form:"g_move_memo" json:"g_move_memo,omitempty"`       // 메모
}

//Comments 구조체는 특이사항을 저장하는 구조체이다.
type Comments struct {
	Comments string `form:"g_bigo" json:"g_bigo,omitempty"` // 기타 내용
}

//Attachments 구조체는 첨부파일 정보를 저장하는 구조체이다.
type Attachments struct {
	File1 string `form:"g_file1" json:"g_file1,omitempty"` //첨부파일 1
	File2 string `form:"g_file2" json:"g_file2,omitempty"` //첨부파일 2
	File3 string `form:"g_file3" json:"g_file3,omitempty"` //첨부파일 3
	File4 string `form:"g_file4" json:"g_file4,omitempty"` //첨부파일 4
}

//Membership 구조체는 가입신청 정보를 관리하는 구조체이다.
type Membership struct {
	AdminInformation    AdminInformation    `json:"admin_info,omitempty"`
	CustomerInformation CustomerInformation `json:"customer_info,omitempty"`
	PaymentInformation  PaymentInformation  `json:"payment_info,omitempty"`
	GiftAccount         GiftAccount         `json:"gift_account,omitempty"`
	ItemInformation     ItemInformation     `json:"item_info,omitempty"`
	NumberMoving        NumberMoving        `json:"number_moving,omitempty"`
	Comments            Comments            `json:"comments,omitempty"`
	Attachments         Attachments         `json:"attachments,omitempty"`
}

func (m *Membership) Validate() error {
	return nil
}
