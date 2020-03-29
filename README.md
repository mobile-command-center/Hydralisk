# 아정통신 CMS 시스템

[![Build Status](https://travis-ci.org/mobile-command-center/Hydralisk.svg?branch=master)](https://travis-ci.org/github/mobile-command-center/Hydralisk)
[![codecov](https://codecov.io/gh/mobile-command-center/Hydralisk/branch/master/graph/badge.svg)](https://codecov.io/gh/mobile-command-center/Hydralisk)
[![Go Report Card](https://goreportcard.com/badge/github.com/mobile-command-center/Hydralisk)](https://goreportcard.com/report/github.com/mobile-command-center/Hydralisk)
[![Documentation](https://godoc.org/github.com/mobile-command-center/Hydralisk?status.svg)](https://godoc.org/github.com/mobile-command-center/Hydralisk)  

## 개발환경 구축
+ golang download  
[다운로드](https://golang.org/dl)페이지로 이동하여 sdk 패키지 다운로드(OS환경에 맞추어 다운로드 한다.)
+ 설치 완료 후 프로젝트 다운로드
```sh
$ git clone https://github.com/mobile-command-center/Hydralisk.git
```
또는  
```sh
$ go get github.com/mobile-command-center/Hydralisk
$ cd $GOPATH/src/github.com/mobile-command-center/Hydralisk
```
+ 프로젝트 빌드
```sh
$ cd <project root>
$ make build
```
빌드 완료 후, 프로젝트 top에 main.zip 파일 생성 확인
+ aws 업로드 (web)
  + [aws](https://aws.amazon.com/ko/) 이동 후 로그인
  + Service > Lambda > 함수 > hydralisk-dev-registration 이동
  + 구성 탭의 함수코드 내 업로드 > main.zip 업로드
  
## 개발 가이드
+ API 가이드  
[API 문서](https://godoc.org/github.com/mobile-command-center/Hydralisk) 확인  
[AWS 람다 가이드](https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/golang-handler.html) 확인  
  
## 트러블슈팅
### 가입신청서 요청받을 시, 이메일로만 전송되고 CMS에 등록되지 않는 경우  
  + [aws](https://aws.amazon.com/ko/) 이동 후 로그인
  + Service > CloudWatch > 로그 > 로그 그룹 > /aws/lambda/hydralisk-dev-registration 이동
  + 시간대별 Log Stream 확인하여 CMS response body 확인
### 에러코드  

| 에러 | 설명 |
| --- | ---- |  
| RequestBodyError | Proxy request에서 HttpRequest 로 변환하는 과정에서 발생하는 에러 |  
| MultiPartParsingError | ajungweb.co.kr 가입신청서 요청 데이터가 Client 구조체로 변환하는 과정에서 발생하는 에러 |  
| TemplateGenerationError | 이메일로 전송될 내용, CMS 첨부파일, 생성에 실패하는 경우 발생하는 에러 |  
| ItemConvertingError | Client 구조체 데이터가 Membership구조체로 변환되는 과정에서 실패하는 경우 발생하는 에러 |  
| ErpDataEncodingError | Membership 구조체가 CMS에 호환되는 데이터로 변환되는 과정에서 실패하는 경우 발생하는 에러 |  
| ErpSystemError | CMS로 데이터 전송시 CMS 서버단에서 에러처리 될 경우 발생하는 에러 |  
+ 에러코드 설명  
  + <b>RequestBodyError</b>  
  Proxy Request Header의 Content-Type 체크 필요(구현부 파싱 실패로 인한 에러이기 때문에 관련 부분 디버깅 필요)
  + <b>MultiPartParsingError</b>  
  가입신청서에 입력된 정보가 비정상 적일 경우 파싱하는 과정에서 인식되지 않아서 에러 발생시킴 (가입신청서 입력데이터 확인 필요)
  + <b>TemplateGenerationError</b>  
  메일 양식 생성시 내부 라이브러리에서 에러 발생함. Template 양식 확인 필요 (client 패키지에 있음)
  + <b>ItemConvertingError</b>  
  상품 파싱시 신규로 추가되거나 상품 이름이 변경되는 경우 정상적인 파싱이 되지 않아서 에러 발생함. 보통은 람다함수 panic걸리기 때문에
  최우선적으로 수정 필요함. CMS 상품 이름, 가입신청서의 상품 이름, 코드등 체크해야함
  + <b>ErpDataEncodingError</b>  
  상품코드로 변환된 데이터들을 CMS에서 인식하는 데이터로 변환하는 과정에서 발생하는 에러. 변환되기 전의 데이터 확인 필요
  + <b>ErpSystemError</b>  
  CMS 서버에러이므로, CMS 서버에 정상적으로 데이터 전송되었는지 확인 필요. 미지원 코드일 경우 에러 발생하는 경우도 있음  
  
### DB에러  
접수ID, 유치ID 누락하게되면 다음과 같이 response body를 CMS 서버단에서 받게 됨. 상품 등록이 거부됨. 
```html
<font face="Arial" size=2>
<p>ADODB.Field</font> <font face="Arial" size=2>error '800a0bcd'</font>
<p>
<font face="Arial" size=2>Either BOF or EOF is True, or the current record has been deleted. Requested operation requires a current record.</font>
<p>
<font face="Arial" size=2>/customer/p_custom_regist_top_ok.asp</font><font face="Arial" size=2>, line 191</font> 
```
## TODO
### 브랜치 전략
+ master : release 브랜치로 안정화 코드만 관리되도록 한다.
+ dev : 개발 브랜치로 관리. 유지보수시 dev 브랜치에서 개발 후 master로 PR 한다.
+ feature : 신규 기능 구현시 dev 브랜치에서 feature 브랜치 생성하 후 dev 브랜치로 PR 한다.
