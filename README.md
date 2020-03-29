# 아정통신 CMS 시스템

[![Build Status](https://travis-ci.org/mobile-command-center/Hydralisk.svg?branch=master)](https://travis-ci.org/github/mobile-command-center/Hydralisk)
[![codecov](https://codecov.io/gh/mobile-command-center/Hydralisk/branch/master/graph/badge.svg)](https://codecov.io/gh/mobile-command-center/Hydralisk)
[![Go Report Card](https://goreportcard.com/badge/github.com/mobile-command-center/Hydralisk)](https://goreportcard.com/report/github.com/mobile-command-center/Hydralisk)
[![Documentation](https://godoc.org/github.com/mobile-command-center/Hydralisk?status.svg)](https://godoc.org/github.com/mobile-command-center/Hydralisk)  
## 공통데이터  
### 공통사항  
| 가입신청서     | CMS    | 비고         |
| --------- | ------ | ---------- |
| 가입자명      | 고객명    |            |
| 휴대폰       | 핸드폰    |            |
| 통신사       |        | 통신사 입력란 업음 |
| 본인여부      | 고객인증 탭 |            |
| 상담 받은 연락처 | 사업자번호  |            |
| 비상연락처     | 대표자    |            |
| 이메일       | 이메일    |            |
| 주소        | 주소     |            |
| 주민등록번호    | 주민등록번호 |            |
### 납부자정보-계좌이체  
| 가입신청서 | CMS  | 비고           |
| ----- | ---- | ------------ |
| 계좌이체  | 계좌이체 | 납부 정보 > 계좌이체 |
| 은행명   | 지급은행 | 납부 정보 > 계좌이체 |
| 계좌번호  | 계좌번호 | 납부 정보 > 계좌이체 |
| 예금주   | 예금주  | 납부 정보 > 계좌이체 |
### 납부자정보-카드납부  
| 가입신청서 | CMS  | 비고           |
| ----- | ---- | ------------ |
| 카드납부  | 카드납부 | 납부 정보 > 카드납부 |
| 카드사   | 카드사  | 납부 정보 > 카드납부 |
| 유효기간  | 유효기간 | 납부 정보 > 카드납부 |
| 카드주명  | 카드주명 | 납부 정보 > 카드납부 |
### 사은품정보  
| 가입신청서  | CMS        | 비고                                              |
| ------ | ---------- | ----------------------------------------------- |
| 상품권 종류 | 본사 사은품     | 상품 정보 1에 입력. 상품 갯수가 DPS/TPS일경우에도 항상 상품 정보 1에 입력 |
| 지급은행   | 상품권계좌-지급은행 |                                                 |
| 계좌번호   | 상품권계좌-계좌번호 |                                                 |
| 예금주    | 상품권계좌-예금주  |                                                 |
## KT  
<b>회사 코드 : 724</b>  
<b>지역 코드 : 22</b>  
  
### 인터넷  
<b>상품코드 : 730</b>  
+ 상품 옵션  
  
| 가입신청서           | CMS  | 코드   |
| --------------- | ---- | ---- |
| 올레인터넷(100M)     | 100M | 1668 |
| 기가인터넷 콤팩트(500M) | 500M | 1844 |
| 기가인터넷(1G)       | 1G   | 1845 |
  
+ 약정 선택  
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 41  |
  
+ 할인탭 선택  
  
| 가입신청서      | CMS   | 코드  |
| ---------- | ----- | --- |
| 신청안함       | 없음    | 275 |
| 인터넷-인터넷 결합 | ★패밀리★ | 312 |
  
+ 부가서비스 선택  
  
| 가입신청서           | CMS      | 코드  |
| --------------- | -------- | --- |
| 신청안함(기존 공유기 사용) | 와이파이신청안함 | 41  |
| 신청              | 와이파이신청함  | 93  |
  
### IPTV  
<b>상품코드 : 732</b>  
+ 상품 옵션  
  
| 가입신청서                   | CMS        | 코드   |
| ----------------------- | ---------- | ---- |
| OTV슬림(구 OTV10) - 228채널  | 슬림(230CH)  | 1677 |
| OTV라이트(구 OTV12) - 238채널 | 라이트(240CH) | 1678 |
| OTV에센스(구 OTV15) - 261채널 | 에센스(270CH) | 1679 |
  
+ 약정 선택  
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 70  |
  
+ 할인탭 선택  
  
| 가입신청서         | CMS     | 코드  |
| ------------- | ------- | --- |
| 신청안함(인터넷 단독시) | 없음      | 281 |
| 일반셋탑          | UHD셋탑   | 308 |
| UHD셋탑         | 기가지니2셋탑 | 309 |
| 기가지니2셋탑       | 일반셋탑    | 310 |
  
+ 부가서비스 선택  
  
| 가입신청서         | CMS   | 코드  |
| ------------- | ----- | --- |
| 신청안함(인터넷단독시)  | 없음    | 53  |
| 총 1대 설치(추가없음) | 총1대설치 | 102 |
| 총 2대 설치       | 총2대설치 | 103 |
| 총 3대 설치       | 총3대설치 | 104 |
| 총 4대 설치       | 총4대설치 | 105 |
  
### 전화
<b>상품코드 : 733</b>  
+ 상품 옵션  
  
| 가입신청서              | CMS     | 코드   |
| ------------------ | ------- | ---- |
| 일반(유선)전화 - 신규가입N   | 번호이동    | 1694 |
| 일반(유선)전화 - 번호이동Y   | 신규가입    | 1695 |
| WiFi(무선)전화 - 신규가입N | 070번호이동 | 1712 |
| WiFi(무선)전화 - 번호이동Y | 070신규   | 1713 |
  
+ 약정 선택  
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 70  |
  
+ 할인탭 선택  
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 341 |
  
+ 부가서비스 선택  
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 68  |

### 사은품
  
| 가입신청서 | CMS           |
| ----- | ------------- |
| 1     | 신세계 상품권       |
| 2     | 홈플러스 상품권 (프리) |
  
## SKB
<b>회사 코드 : 78</b>  
<b>지역 코드 : 1</b>  

### 인터넷
<b>상품코드 : 79</b>  
+ 상품 옵션
  
| 가입신청서           | CMS  | 코드   |
| --------------- | ---- | ---- |
| 스마트 다이렉트(100M)  | 100M | 114  |
| 기가인터넷 라이트(500M) | 500M | 115  |
| 기가인터넷(1G)       | 1G   | 1705 |
  
+ 약정 선택
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 33  |
  
+ 할인탭 선택
  
| 가입신청서         | CMS       | 코드  |
| ------------- | --------- | --- |
| 결합없음          | 없음        | 259 |
| SKT인터넷-인터넷 결합 | ★패밀리★     | 329 |
| 2,30년 이상(온할인) | 온가족할인(년수) | 331 |
| 휴대폰 2대이상(온플랜) | 온프리(1회선)  | 332 |
| 기타(요청란에 기입)   | 기타특이사항확인  | 336 |
  
+ 부가서비스 선택  
  
| 가입신청서           | CMS      | 코드  |
| --------------- | -------- | --- |
| 신청안함(기존 공유기 사용) | 와이파이신청안함 | 4   |
| 신청              | 와이파이신청함  | 118 |
  
### 전화
<b>상품코드 : 81</b>  
+ 상품 옵션  
  
| 가입신청서              | CMS     | 코드   |
| ------------------ | ------- | ---- |
| 일반(유선)전화 - 신규가입N   | 신규가입    | 1696 |
| 일반(유선)전화 - 번호이동Y   | 번호이동    | 117  |
| WiFi(무선)전화 - 신규가입N | 070신규   | 1715 |
| WiFi(무선)전화 - 번호이동Y | 070번호이동 | 1714 |
  
+ 약정 선택  
    
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 43  |
  
+ 할인탭 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 260 |

+ 부가서비스 선택  
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 15  |
  
### BTV
<b>상품코드 : 83</b>  
+ 상품 옵션
  
| 가입신청서             | CMS        | 코드  |
| ----------------- | ---------- | --- |
| B TV Lite - 211채널 | 라이트(212CH) | 127 |
| B TV All - 234채널  | 올(224CH)   | 128 |
  
+ 약정 선택
    
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 50  |
  
+ 할인탭 선택
  
| 가입신청서         | CMS    | 코드  |
| ------------- | ------ | --- |
| 신청안함(인터넷 단독시) | 없음     | 270 |
| 스마트티비셋탑       | 스마트3셋탑 | 302 |
| NUGU셋탑        | 누구2셋탑  | 303 |
  
+ 부가서비스 선택
  
| 가입신청서         | CMS   | 코드  |
| ------------- | ----- | --- |
| 신청안함(인터넷단독시)  | 없음    | 117 |
| 총 1대 설치(추가없음) | 총1대설치 | 19  |
| 총 2대 설치       | 총2대설치 | 114 |
| 총 3대 설치       | 총3대설치 | 115 |
| 총 4대 설치       | 총4대설치 | 116 |
  
### 사은품
  
| 가입신청서 | CMS           |
| ----- | ------------- |
| 1     | 신세계 상품권       |
| 2     | 홈플러스 상품권 (프리) |
| 3     | OK 캐시백        |
| 4     | SK 상품권 (플랜)   |
    
## SKT
<b>회사 코드 : 735</b>  
<b>지역 코드 : 4</b>  

### 인터넷
<b>상품코드 : 736</b>  
+ 상품 옵션  
  
| 가입신청서           | CMS  | 코드   |
| --------------- | ---- | ---- |
| 스마트 다이렉트(100M)  | 100M | 1700 |
| 기가인터넷 라이트(500M) | 500M | 1701 |
| 기가인터넷(1G)       | 1G   | 1716 |
  
+ 약정 선택
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 84  |
  
+ 할인탭 선택
  
| 가입신청서         | CMS       | 코드  |
| ------------- | --------- | --- |
| 결합없음          | 없음        | 301 |
| SKT인터넷-인터넷 결합 | ★패밀리★     | 330 |
| 2,30년 이상(온할인) | 온가족할인(년수) | 334 |
| 휴대폰 2대이상(온플랜) | 온플랜(2회선↑) | 333 |
| 기타(요청란에 기입)   | 기타특이사항확인  | 335 |
  
+ 부가서비스 선택  
  
| 가입신청서           | CMS      | 코드  |
| --------------- | -------- | --- |
| 신청안함(기존 공유기 사용) | 와이파이신청안함 | 120 |
| 신청              | 와이파이신청함  | 119 |
  
### 전화
<b>상품코드 : 737</b>  
+ 상품 옵션  
  
| 가입신청서              | CMS     | 코드   |
| ------------------ | ------- | ---- |
| 일반(유선)전화 - 신규가입N   | 신규가입    | 1702 |
| 일반(유선)전화 - 번호이동Y   | 번호이동    | 1703 |
| WiFi(무선)전화 - 신규가입N | 070신규   | 1720 |
| WiFi(무선)전화 - 번호이동Y | 070번호이동 | 1721 |
  
+ 약정 선택  
    
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 88  |
  
+ 할인탭 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 340 |

+ 부가서비스 선택  
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 79  |
  
### BTV
<b>상품코드 : 738</b>  
+ 상품 옵션

| 가입신청서             | CMS        | 코드   |
| ----------------- | ---------- | ---- |
| B TV Lite - 211채널 | 라이트(212CH) | 1704 |
| B TV All - 234채널  | 올(224CH)   | 1722 |
  
+ 약정 선택
    
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 92  |
  
+ 할인탭 선택
  
| 가입신청서         | CMS    | 코드  |
| ------------- | ------ | --- |
| 신청안함(인터넷 단독시) | 없음     | 330 |
| 스마트티비셋탑       | 스마트3셋탑 | 305 |
| NUGU셋탑        | 누구2셋탑  | 306 |
  
+ 부가서비스 선택
  
| 가입신청서         | CMS   | 코드  |
| ------------- | ----- | --- |
| 신청안함(인터넷단독시)  | 없음    | 85  |
| 총 1대 설치(추가없음) | 총1대설치 | 110 |
| 총 2대 설치       | 총2대설치 | 111 |
| 총 3대 설치       | 총3대설치 | 112 |
| 총 4대 설치       | 총4대설치 | 113 |

### 사은품
  
| 가입신청서 | CMS           |
| ----- | ------------- |
| 1     | 신세계 상품권       |
| 2     | 홈플러스 상품권 (프리) |
| 3     | OK 캐시백        |
| 4     | SK 상품권 (플랜)   |
  
## LG U+
<b>회사 코드 : 80</b>  
<b>지역 코드 : 16</b>  

### 인터넷
<b>상품코드 : 84</b>  
+ 상품 옵션  
  
| 가입신청서               | CMS       | 코드   |
| ------------------- | --------- | ---- |
| 와이파이기본_광랜안심(100M)   | 100M+WIFI | 1725 |
| 와이파이기본_기가슬림안심(500M) | 500M+WIFI | 1726 |
| 와이파이기본_기가안심(1G)     | 1G+WIFI   | 1727 |
  
+ 약정 선택  
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 58  |
  
+ 할인탭 선택  
  
| 가입신청서      | CMS   | 코드  |
| ---------- | ----- | --- |
| 신청안함       | 없음    | 272 |
| 인터넷-인터넷 결합 | ★패밀리★ | 328 |
  
+ 부가서비스 선택  
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 30  |

### 전화
<b>상품코드 : 85</b>  
+ 상품 옵션
  
| 가입신청서              | CMS        | 코드  |
| ------------------ | ---------- | --- |
| 일반(유선)전화 - 신규가입N   | [CPG]신규가입  | 135 |
| 일반(유선)전화 - 번호이동Y   | [CPG]번호이동  | 133 |
| WiFi(무선)전화 - 신규가입N | [WIFI]신규가입 | 137 |
| WiFi(무선)전화 - 번호이동Y | [WIFI]번호이동 | 139 |
  
+ 약정 선택
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 61  |
  
+ 할인탭 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 287 |
  
+ 부가서비스 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 76  |
  
### IPTV
<b>상품코드 : 86</b>  
+ 상품 옵션
  
| 가입신청서                    | CMS                  | 코드   |
| ------------------------ | -------------------- | ---- |
| TV베이직 - 183채널            | 베이직(185CH)           | 141  |
| TV프리미엄 - 223채널           | 프리미엄(224CH)          | 142  |
| TV VOD고급형 - 210채널        | VOD고급(210CH)         | 1856 |
| 프리미엄 넷플릭스HD TV - 223채널   | 프리미엄 넷플릭스 HD(224CH)  | 1728 |
| 프리미엄 넷플릭스ㅕHD TV - 223채널 | 프리미엄 넷플릭스 UHD(224CH) | 1854 |
  
+ 약정 선택
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 64  |
  
+ 할인탭 선택
  
| 가입신청서         | CMS   | 코드  |
| ------------- | ----- | --- |
| 신청안함(인터넷 단독시) | 없음    | 274 |
| UHD셋탑         | UHD셋탑 | 311 |
  
+ 부가서비스 선택
  
| 가입신청서         | CMS   | 코드  |
| ------------- | ----- | --- |
| 신청안함(인터넷단독시)  | 없음    | 36  |
| 총 1대 설치(추가없음) | 총1대설치 | 106 |
| 총 2대 설치       | 총2대설치 | 107 |
| 총 3대 설치       | 총3대설치 | 108 |
| 총 4대 설치       | 총4대설치 | 109 |

### 사은품
  
| 가입신청서 | CMS           |
| ----- | ------------- |
| 1     | 신세계 상품권       |
| 2     | 홈플러스 상품권 (프리) |
  
## LG헬로비전
<b>회사 코드 : 740</b>  
<b>지역 코드 : 23</b>  

### 인터넷
<b>상품코드 : 744</b>  
+ 상품 옵션
  
| 가입신청서         | CMS          | 코드   |
| ------------- | ------------ | ---- |
| 광랜라이트+(100MB) | 광랜라이트(100MB) | 1848 |
| 광랜(160MB)     | 광랜(160MB)    | 1849 |
| 기가라이트(500MB)  | 기가라이트(500MB) | 1850 |
| 플래티넘기가(1G)    | 플래티넘기가(1G)   | 1851 |
  
+ 약정 선택  
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 93  |
  
+ 할인탭 선택  
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 313 |
  
+ 부가서비스 선택  
  
| 가입신청서 | CMS      | 코드  |
| ----- | -------- | --- |
| 신청안함  | 와이파이신청안함 | 96  |
| 신청    | 와이파이신청함  | 95  |
  
### TV
<b>상품코드 : 745</b>  
+ 상품 옵션
  
| 가입신청서        | CMS      | 코드   |
| ------------ | -------- | ---- |
| 베이직 UHD SMT  | 베이직 UHD  | 1852 |
| 프리미엄 UHD SMT | 프리미엄 UHD | 1853 |
  
+ 약정 선택
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 94  |
  
+ 할인탭 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 314 |
  
+ 부가서비스 선택
  
| 가입신청서   | CMS           | 코드  |
| ------- | ------------- | --- |
| 신청안함    | 추가설치없음        | 131 |
| 베이직 HD  | 베이직HD 1대더추가설치 | 129 |
| 이코노미 HD | 이코노미HD1대더추가설치 | 130 |
  
### 사은품
  
| 가입신청서 | CMS           |
| ----- | ------------- |
| 1     | 신세계 상품권       |
| 2     | 홈플러스 상품권 (프리) |
  
## 스카이라이프
<b>회사 코드 : 739</b>  
<b>지역 코드 : 24</b>

### 인터넷
<b>상품코드 : 747</b>  
+ 상품 옵션
  
| 가입신청서        | CMS  | 코드   |
| ------------ | ---- | ---- |
| 올레인터넷(100M)  | 100M | 1739 |
| 기가(200M)     | 200M | 1740 |
| 기가 콤팩트(500M) | 500M | 1846 |
| 기가(1G)       | 1G   | 1847 |
  
+ 약정 선택
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 96  |
  
+ 할인탭 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
| 신청안함  | 없음  | 319 |
| 홈결합   | 홈결합 | 337 |
  
+ 부가서비스 선택
  
| 가입신청서 | CMS      | 코드  |
| ----- | -------- | --- |
|       | 일반와이파이신청 | 126 |
|       | 기가와이파이신청 | 127 |
|       | 신청안함     | 128 |
  
### TV
<b>상품코드 : 748</b>  
+ 상품 옵션
  
| 가입신청서        | CMS             | 코드   |
| ------------ | --------------- | ---- |
| SKY GReen A+ | Green A+(195CH) | 1745 |
| SKY Blue A+  | Blue A+(218CH)  | 1744 |
  
+ 약정 선택
  
| 가입신청서 | CMS  | 코드  |
| ----- | ---- | --- |
| 3년약정  | 3년약정 | 97  |
  
+ 할인탭 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 없음  | 322 |
  
+ 부가서비스 선택
  
| 가입신청서         | CMS   | 코드  |
| ------------- | ----- | --- |
| 신청안함(인터넷단독시)  | 없음    | 101  |
| 총 1대 설치(추가없음) | 총1대설치 | 122 |
| 총 2대 설치       | 총2대설치 | 123 |
| 총 3대 설치       | 총3대설치 | 124 |
| 총 4대 설치       | 총4대설치 | 125 |
  
### 사은품
  
| 가입신청서 | CMS           |
| ----- | ------------- |
| 1     | 신세계 상품권       |
| 2     | 홈플러스 상품권 (프리) |
  
## 렌탈
<b>회사 코드 : 743</b>  
<b>지역 코드 : 25</b>  
### 미기입
<b>상품코드 : 761</b>  
+ 상품 옵션
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 미기입  | 1855 |
  
+ 약정 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 미기입  | 132 |
  
+ 할인탭 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 미기입  | 339 |
  
+ 부가서비스 선택
  
| 가입신청서 | CMS | 코드  |
| ----- | --- | --- |
|       | 미기입  | 132 |
  



