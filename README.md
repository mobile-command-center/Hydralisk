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
  + [aws](https://aws.amazon.com/ko/) 이동후 로그인
  + Service > Lambda > 함수 > hydralisk-dev-registration 이동
  + 구성 탭의 함수코드 내 업로드 > main.zip 업로드

## 개발 가이드
+ API 가이드  
[API 문서](https://godoc.org/github.com/mobile-command-center/Hydralisk) 확인  
[AWS 람다 가이드](https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/golang-handler.html) 확인
