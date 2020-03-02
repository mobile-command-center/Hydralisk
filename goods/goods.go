package goods

import (
	"fmt"
	"github.com/mobile-command-center/Hydralisk/client"
	"strconv"
)

//AdminInformation 구조체는 고객등록시 필요한 관리자 ID이다. Cookie로 받아올 수 있다.
type AdminInformation struct {
	Jupsu string `form:"g_jupsu_m_id"` //접수 ID
	Yuchi string `form:"g_yuchi_m_id"` //유치 ID
}

//CustomerInformation 구조체는 고객정보탭의 정보를 저장하는 구조체이다.
type CustomerInformation struct {
	Name                       string `form:"c_name"`            //고객명
	CustomerType               string `form:"c_type"`            //고객 종류
	RegistrationNumber1        string `form:"c_jumin1"`          //주민번호 앞자리
	RegistrationNumber2        string `form:"c_jumin2"`          //주민번호 뒷자리
	BusinessRegistrationNumber string `form:"c_office_no"`       //사업자번호
	BusinessOwenerName         string `form:"c_office_name"`     //대표자
	Zipcode                    string `form:"c_zipcode1"`        //우편번호
	Address                    string `form:"c_address"`         //주소
	Telephone                  string `form:"c_tel1"`            //전화
	Phone                      string `form:"c_tel2"`            //휴대전화
	Email                      string `form:"c_email"`           //이메일
	AuthType                   string `form:"c_auth_chk"`        //고객인증 타입
	AuthData                   string `form:"c_jumin_date"`      //고객인증 데이터
	RegistrationCourse         string `form:"g_code_course_idx"` //접수 경로
}

//1. 가입자명 -> 고객명
//2. 휴대폰 번호 -> 핸드폰
//3. 통신사 -> 전화 내용에 삽입(삽입 불가)
//4. 본인여부 -> 고객인증탭의 빈칸에 작성
//5. 상담받은 연락처 -> 사업자번호칸에 작성
//6. 비상연락처 -> 대표자 칸에 작성
//7. 이메일 -> 이메일
//8. 주소 -> 주소
//9. 주민등록번호 -> 주민등록번호칸 (외국인 체크 무시)
//10. 외국인 체크 -> 종류(외국인)
func (customer *CustomerInformation) Convert(c client.Client) {
	customer.Name = c.Name
	customer.CustomerType = getCustomerType("개인")
	customer.Phone = c.Tel21 + "-" + c.Tel22 + "-" + c.Tel23
	customer.AuthData = c.Auth
	customer.BusinessRegistrationNumber = c.Tel31 + "-" + c.Tel32 + "-" + c.Tel33
	customer.BusinessOwenerName = c.Tel11 + "-" + c.Tel12 + "-" + c.Tel13
	customer.Email = c.Email1 + "@" + c.Email2
	customer.Zipcode = c.ZipCode1 + c.ZipCode2
	customer.Address = c.Address1 + " " + c.Address2
	customer.RegistrationNumber1 = c.Jumin1
	customer.RegistrationNumber2 = c.Jumin2
	customer.RegistrationCourse = getRegistrationCourse("없음")
}

//AccountTransfer 구조체는 계좌이체 정보를 저장하는 구조체이다.
type AccountTransfer struct {
	Bank                string `form:"c_bank_cd"`     //지급은행
	Account             string `form:"c_bank_no"`     //계좌번호
	OwenerName          string `form:"c_bank_name"`   //예금주
	RegistrationNumber1 string `form:"c_bank_jumin1"` //예금주 주민번호 앞자리
	RegistrationNumber2 string `form:"c_bank_jumin2"` //예금주 주민번호 뒷자리
}

func (a *AccountTransfer) Convert(c client.Client) {
	if c.BankCd == "0" {
		a.Bank = c.BankCd
		return
	}

	a.Bank = c.BankCd
	a.Account = c.BankNo
	a.OwenerName = c.BankHolder
}

//CreditCard 구조체는 신용카드 정보를 저장하는 구조체이다.
type CreditCard struct {
	Company             string `form:"c_card_cd"`     //카드사
	Number              string `form:"g_card_no"`     //카드번호
	ExpiredYear         string `form:"g_card_gigan"`  //카드 유효기간(년)
	ExpiredMonth        string `form:"g_card_gigan1"` //카드 유효기간(월)
	OwenerName          string `form:"g_card_name"`   //카드주 명
	RegistrationNumber1 string `form:"c_card_jumin1"` //카드주 주민번호 앞자리
	RegistrationNumber2 string `form:"c_card_jumin2"` //카드주 주민번호 뒷자리
}

func (credit *CreditCard) Convert(c client.Client) {
	if c.CardCd == "0" {
		credit.Company = c.CardCd
		return
	}

	credit.Company = c.CardCd
	credit.Number = c.CardNo
	credit.ExpiredMonth = c.CardGigan1
	credit.ExpiredYear = c.CardGigan2
	credit.OwenerName = c.CardHolder
}

//Billing 구조체는 지로납부 정보를 저장하는 구조체이다.
//현재 지로납부에 대해서는 저장해야할 데이터가 없다.
type Billing struct{}

//Summing 구조체는 합산청구 정보를 저장하는 구조체이다.
type Summing struct {
	Billing  string `form:"g_sum_money_chk"` //합산 청구
	Comments string `form:"g_sum_money_txt"` //합산 비고
}

//PaymentInformation 구조체는 납부 정보를 저장하는 구조체이다.
type PaymentInformation struct {
	PaymentType     string          `form:"g_give_type"` //납부 정보
	Relationship    string          `form:"g_give_chk"`  //관계 여부
	AccountTransfer AccountTransfer //계좌이체
	CreditCard      CreditCard      //카드납부
	Billing         Billing         //지로납부
	Summing         Summing         // 합산청구
}

func (p *PaymentInformation) Convert(c client.Client) {
	p.Relationship = getRelationship("본인")
	if c.BankCd != "0" {
		p.PaymentType = getGiveType("계좌이체")
		p.AccountTransfer.Convert(c)
	}
	if c.CardCd != "0" {
		p.PaymentType = getGiveType("카드납부")
		p.CreditCard.Convert(c)
	}
}

//GiftAccount 구조체는 사은품 계좌 정보를 저장하는 구조체이다.
type GiftAccount struct {
	Bank                string `form:"g_sp_bank_code"`   //지급은행
	Account             string `form:"g_sp_bank_acount"` //계좌번호
	OwenerName          string `form:"g_sp_bank_name"`   //예금주
	RegistrationNumber1 string `form:"g_sp_bank_jumin1"` //예금주 주민번호 앞자리
	RegistrationNumber2 string `form:"g_sp_bank_jumin2"` //예금주 주민번호 뒷자리
}

func (g *GiftAccount) Convert(c client.Client) {
	if c.SpBankCode == "0" {
		g.Bank = "0"
		return
	}

	g.Bank = c.SpBankCode
	g.Account = c.SpBankAccount
	g.OwenerName = c.SpBankHolder
}

type DummyItem struct {
	Item            string
	Option          string
	Promise         string
	Sale            string
	Service         string
	LineCount       string
	GiftName        string
	GiftPrice       string
	GiftPaymentDay  string
	GiftPaymentType string
	ReviewPrice     string
	TopGiftName     string
	TopGiftPrice    string
}

//FirstItem 구조체는 첫번째 사은품 정보를 저장하는 구조체이다.
type FirstItem struct {
	Item            string `form:"g_article_idx2"`  //상품 선택
	Option          string `form:"g_option_idx"`    //옵션 선택
	Promise         string `form:"g_code_promise1"` //약정 선택
	Sale            string `form:"g_code_sale"`     //할인탭 선택
	Service         string `form:"g_code_service1"` //부가서비스
	LineCount       string `form:"g_article_cnt1"`  //설치 회선
	GiftName        string `form:"g_sp_name"`       //사은품명
	GiftPrice       string `form:"g_sp_price"`      //사은품 가격
	GiftPaymentDay  string `form:"g_sp_date_start"` //사은품 지급 날짜
	GiftPaymentType string `form:"g_sp_give_type"`  //사은품 지급자 선택
	ReviewPrice     string `form:"g_sp_price_add1"` //후기금액
	TopGiftName     string `form:"g_sp_name_top1"`  //본사 사은품명
	TopGiftPrice    string `form:"g_sp_price_top1"` //본사 사은품 가격
}

func (f *FirstItem) Convert(i interface{}) {
	d := i.(*DummyItem)

	f.Item = d.Item
	f.Option = d.Option
	f.Promise = d.Promise
	f.Sale = d.Sale
	f.Service = d.Service
	f.LineCount = d.LineCount
	f.GiftName = d.GiftName
	f.GiftPrice = d.GiftPrice
	f.GiftPaymentDay = d.GiftPaymentDay
	f.GiftPaymentType = d.GiftPaymentType
	f.ReviewPrice = d.ReviewPrice
	f.TopGiftName = d.TopGiftName
	f.TopGiftPrice = d.TopGiftPrice
}

//SecondItem 구조체는 번째 사은품 정보를 저장하는 구조체이다.
type SecondItem struct {
	Item            string `form:"g_article_idx2_1"` //상품 선택
	Option          string `form:"g_option_idx_1"`   //옵션 선택
	Promise         string `form:"g_code_promise2"`  //약정 선택
	Sale            string `form:"g_code_sale1"`     //할인탭 선택
	Service         string `form:"g_code_service2"`  //부가서비스
	LineCount       string `form:"g_article_cnt2"`   //설치 회선
	GiftName        string `form:"g_sp_name1"`       //사은품명
	GiftPrice       string `form:"g_sp_price1"`      //사은품 가격
	GiftPaymentDay  string `form:"g_sp_date_start1"` //사은품 지급 날짜
	GiftPaymentType string `form:"g_sp_give_type1"`  //사은품 지급자 선택
	ReviewPrice     string `form:"g_sp_price_add2"`  //후기금액
	TopGiftName     string `form:"g_sp_name_top2"`   //본사 사은품명
	TopGiftPrice    string `form:"g_sp_price_top2"`  //본사 사은품 가격
}

func (s *SecondItem) Convert(i interface{}) {
	d := i.(*DummyItem)

	s.Item = d.Item
	s.Option = d.Option
	s.Promise = d.Promise
	s.Sale = d.Sale
	s.Service = d.Service
	s.LineCount = d.LineCount
	s.GiftName = d.GiftName
	s.GiftPrice = d.GiftPrice
	s.GiftPaymentDay = d.GiftPaymentDay
	s.GiftPaymentType = d.GiftPaymentType
	s.ReviewPrice = d.ReviewPrice
	s.TopGiftName = d.TopGiftName
	s.TopGiftPrice = d.TopGiftPrice
}

//ThirdItem 구조체는 세번째 사은품 정보를 저장하는 구조체이다.
type ThirdItem struct {
	Item            string `form:"g_article_idx2_2"` //상품 선택
	Option          string `form:"g_option_idx_2"`   //옵션 선택
	Promise         string `form:"g_code_promise3"`  //약정 선택
	Sale            string `form:"g_code_sale2"`     //할인탭 선택
	Service         string `form:"g_code_service3"`  //부가서비스
	LineCount       string `form:"g_article_cnt3"`   //설치 회선
	GiftName        string `form:"g_sp_name2"`       //사은품명
	GiftPrice       string `form:"g_sp_price2"`      //사은품 가격
	GiftPaymentDay  string `form:"g_sp_date_start2"` //사은품 지급 날짜
	GiftPaymentType string `form:"g_sp_give_type2"`  //사은품 지급자 선택
	ReviewPrice     string `form:"g_sp_price_add3"`  //후기금액
	TopGiftName     string `form:"g_sp_name_top3"`   //본사 사은품명
	TopGiftPrice    string `form:"g_sp_price_top3"`  //본사 사은품 가격
}

func (t *ThirdItem) Convert(i interface{}) {
	d := i.(*DummyItem)

	t.Item = d.Item
	t.Option = d.Option
	t.Promise = d.Promise
	t.Sale = d.Sale
	t.Service = d.Service
	t.LineCount = d.LineCount
	t.GiftName = d.GiftName
	t.GiftPrice = d.GiftPrice
	t.GiftPaymentDay = d.GiftPaymentDay
	t.GiftPaymentType = d.GiftPaymentType
	t.ReviewPrice = d.ReviewPrice
	t.TopGiftName = d.TopGiftName
	t.TopGiftPrice = d.TopGiftPrice
}

//ItemInformation 구조체는 상품 정보를 저장하는 구조체이다.
type ItemInformation struct {
	Company    string `form:"g_article_idx1"` // 회사 선택
	Location   string `form:"g_code_area"`    //지역 선택
	GoodsCount string `form:"goods_cnt"`      //상품 갯수
	FirstItem  *FirstItem
	SecondItem *SecondItem
	ThirdItem  *ThirdItem
}

func (i *ItemInformation) Convert(c client.Client) {
	company := newCommpany()
	commCompany := company[replaceCompanyName(c.Vendor)]

	i.Company = commCompany.CompanyCode
	i.Location = commCompany.AreaCode

	parser := newServiceParser()
	count := parser[c.Vendor](i, c, commCompany)

	i.GoodsCount = strconv.Itoa(count)
}

//NumberMoving 구조체는 번호이동 정보를 저장하는 구조체이다.
type NumberMoving struct {
	Telephone           string `form:"g_move_tel1"`    //전화번호
	Name                string `form:"g_move_name"`    //명의자
	RegistrationNumber1 string `form:"g_move_jumin1"`  //주민번호 앞자리
	RegistrationNumber2 string `form:"g_move_jumin2"`  //주민번호 뒷자리
	Phone               string `form:"g_move_tel2"`    //연락처
	Date                string `form:"g_move_date"`    //발급 일자
	PreviousCompany     string `form:"g_move_company"` //기존통신사
	Memo                string `form:"g_move_memo"`    // 메모
}

func (n *NumberMoving) Convert(c client.Client) {
	if c.TelephoneCarrierMoveChk == "true" {
		n.Telephone = c.MoveTel1
		company := getMoveCommCompany(c.MoveCompany)
		if company == "기타" {
			n.Memo = "알뜰폰"
		} else {
			n.PreviousCompany = company
		}
	}
}

//Comments 구조체는 특이사항을 저장하는 구조체이다.
type Comments struct {
	Comments string `form:"g_bigo"` // 기타 내용
}

func (comment *Comments) Convert(c client.Client) {
	var s string
	if c.Vendor == "rental" {
		fmt.Println("RENTAL")
		s = c.RentalVendor + " / " + c.RentalProduct + " / " + c.RentalProductName + " / " + c.RentalProductColor + " / " + c.RentalPromise + "\n"

	}
	if s != "" {
		comment.Comments = s + c.Bigo
		return
	}
	comment.Comments = c.Bigo
}

//Attachments 구조체는 첨부파일 정보를 저장하는 구조체이다.
type Attachments struct {
	File1 string `form:"g_file1"` //첨부파일 1
	File2 string `form:"g_file2"` //첨부파일 2
	File3 string `form:"g_file3"` //첨부파일 3
	File4 string `form:"g_file4"` //첨부파일 4
}

//Membership 구조체는 가입신청 정보를 관리하는 구조체이다.
type Membership struct {
	AdminInformation    AdminInformation
	CustomerInformation CustomerInformation
	PaymentInformation  PaymentInformation
	GiftAccount         GiftAccount
	ItemInformation     ItemInformation
	NumberMoving        NumberMoving
	Comments            Comments
	Attachments         Attachments
}

func (m *Membership) Convert(c client.Client) error {
	m.CustomerInformation.Convert(c)
	m.PaymentInformation.Convert(c)
	m.GiftAccount.Convert(c)
	m.ItemInformation.Convert(c)
	m.NumberMoving.Convert(c)
	m.Comments.Convert(c)

	return nil
}
