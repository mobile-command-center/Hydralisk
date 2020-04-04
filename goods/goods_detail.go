package goods

//Article 구조체는 통신사별 상품 및 세부 정보를 갖는 구조체이다.
type Article struct {
	Article string
	Options map[string]string
	Promise map[string]string
	Sale    map[string]string
	Service map[string]string
}

func newCommpany() map[string]*CommCompany {
	return map[string]*CommCompany{
		"KT":     newKt(),
		"SKB":    newSkb(),
		"SKT":    newSkt(),
		"LG U+":  newLgu(),
		"LG헬로비전": newLgHelloVision(),
		"스카이라이프": newSkylife(),
		"렌탈":     newRental(),
	}
}

//CommCompany 구조체는 통신사에 대한 정보를 갖는 구조체이다.
type CommCompany struct {
	CompanyCode string
	AreaCode    string
	Article     map[string]*Article
}

func newKt() *CommCompany {
	comm := &CommCompany{
		CompanyCode: "724",
		AreaCode:    "22",
		Article:     getKtArticle(),
	}
	return comm
}

func getKtArticle() map[string]*Article {
	return map[string]*Article{
		"인터넷": {
			Article: "730",
			Options: map[string]string{
				"100M": "1668",
				"500M": "1844",
				"1G":   "1845",
			},
			Promise: map[string]string{
				"3년약정": "41",
			},
			Sale: map[string]string{
				"없음":    "275",
				"★패밀리★": "312",
			},
			Service: map[string]string{
				"와이파이신청안함": "41",
				"와이파이신청함":  "93",
			},
		},
		"IPTV": {
			Article: "732",
			Options: map[string]string{
				"슬림(230CH)":  "1677",
				"라이트(240CH)": "1678",
				"에센스(270CH)": "1679",
				"엔터(270CH)":  "1680",
				"베이직(223CH)": "1857",
			},
			Promise: map[string]string{
				"3년약정": "70",
			},
			Sale: map[string]string{
				"없음":      "281",
				"UHD셋탑":   "308",
				"기가지니2셋탑": "309",
				"일반셋탑":    "310",
			},
			Service: map[string]string{
				"없음":    "53",
				"총1대설치": "102",
				"총2대설치": "103",
				"총3대설치": "104",
				"총4대설치": "105",
			},
		},
		"전화": {
			Article: "733",
			Options: map[string]string{
				"번호이동":    "1694",
				"신규가입":    "1695",
				"070번호이동": "1712",
				"070신규":   "1713",
			},
			Promise: map[string]string{
				"3년약정": "79",
			},
			Sale: map[string]string{
				"없음": "341",
			},
			Service: map[string]string{
				"없음": "68",
			},
		},
	}
}

func newSkb() *CommCompany {
	comm := &CommCompany{
		CompanyCode: "78",
		AreaCode:    "1",
		Article:     getSkbArticle(),
	}
	return comm
}

func getSkbArticle() map[string]*Article {
	return map[string]*Article{
		"인터넷": {
			Article: "79",
			Options: map[string]string{
				"100M": "114",
				"500M": "115",
				"1G":   "1705",
			},
			Promise: map[string]string{
				"3년약정": "33",
			},
			Sale: map[string]string{
				"없음":        "259",
				"★패밀리★":     "329",
				"온가족할인(년수)": "331",
				"온프리(1회선)":  "332",
				"기타특이사항확인":  "336",
			},
			Service: map[string]string{
				"와이파이신청안함": "4",
				"와이파이신청함":  "118",
			},
		},
		"전화": {
			Article: "81",
			Options: map[string]string{
				"번호이동":    "117",
				"신규가입":    "1696",
				"070번호이동": "1714",
				"070신규":   "1715",
			},
			Promise: map[string]string{
				"3년약정": "43",
			},
			Sale: map[string]string{
				"없음": "260",
			},
			Service: map[string]string{
				"없음": "15",
			},
		},
		"BTV": {
			Article: "83",
			Options: map[string]string{
				"라이트(212CH)": "127",
				"올(224CH)":   "128",
				"올+(224CH)":  "129",
			},
			Promise: map[string]string{
				"3년약정": "50",
			},
			Sale: map[string]string{
				"없음":     "270",
				"스마트3셋탑": "302",
				"누구2셋탑":  "303",
				"UHD2셋탑": "304",
			},
			Service: map[string]string{
				"총1대설치": "19",
				"총2대설치": "114",
				"총3대설치": "115",
				"총4대설치": "116",
				"없음":    "117",
			},
		},
	}
}

func newSkt() *CommCompany {
	comm := &CommCompany{
		CompanyCode: "735",
		AreaCode:    "4",
		Article:     getSktArticle(),
	}
	return comm
}

func getSktArticle() map[string]*Article {
	return map[string]*Article{
		"인터넷": {
			Article: "736",
			Options: map[string]string{
				"100M": "1700",
				"500M": "1701",
				"1G":   "1716",
			},
			Promise: map[string]string{
				"3년약정": "84",
			},
			Sale: map[string]string{
				"없음":        "301",
				"★패밀리★":     "330",
				"온플랜(2회선↑)": "333",
				"온가족할인(년수)": "334",
				"기타특이사항확인":  "335",
			},
			Service: map[string]string{
				"와이파이신청안함": "120",
				"와이파이신청함":  "119",
			},
		},
		"전화": {
			Article: "737",
			Options: map[string]string{
				"신규가입":    "1702",
				"번호이동":    "1703",
				"070신규":   "1720",
				"070번호이동": "1721",
			},
			Promise: map[string]string{
				"3년약정": "88",
			},
			Sale: map[string]string{
				"없음": "340",
			},
			Service: map[string]string{
				"없음": "79",
			},
		},
		"BTV": {
			Article: "738",
			Options: map[string]string{
				"라이트(212CH)": "1704",
				"올(224CH)":   "1722",
				"올+(224CH)":  "1723",
			},
			Promise: map[string]string{
				"3년약정": "92",
			},
			Sale: map[string]string{
				"없음":     "300",
				"스마트3셋탑": "305",
				"누구2셋탑":  "306",
				"UHD2셋탑": "307",
			},
			Service: map[string]string{
				"없음":    "85",
				"총1대설치": "110",
				"총2대설치": "111",
				"총3대설치": "112",
				"총4대설치": "113",
			},
		},
	}
}

func newLgu() *CommCompany {
	comm := &CommCompany{
		CompanyCode: "80",
		AreaCode:    "16",
		Article:     getLguArticle(),
	}
	return comm
}

func getLguArticle() map[string]*Article {
	return map[string]*Article{
		"인터넷": {
			Article: "84",
			Options: map[string]string{
				"100M+WIFI": "1725",
				"500M+WIFI": "1726",
				"1G+WIFI":   "1727",
			},
			Promise: map[string]string{
				"3년약정": "58",
			},
			Sale: map[string]string{
				"없음":    "272",
				"★패밀리★": "328",
			},
			Service: map[string]string{
				"없음":      "30",
				"휴대폰직접결합": "92",
			},
		},
		"전화": {
			Article: "85",
			Options: map[string]string{
				"[WIFI]번호이동": "133",
				"[WIFI]신규가입": "135",
				"[CPG]번호이동":  "137",
				"[CPG]신규가입":  "139",
			},
			Promise: map[string]string{
				"3년약정": "61",
			},
			Sale: map[string]string{
				"없음": "287",
			},
			Service: map[string]string{
				"없음": "76",
			},
		},
		"IPTV": {
			Article: "86",
			Options: map[string]string{
				"베이직(185CH)":           "141",
				"프리미엄(224CH)":          "142",
				"VOD고급(210CH)":         "1856",
				"프리미엄 넷플릭스 HD(224CH)":  "1728",
				"프리미엄 넷플릭스 UHD(224CH)": "1854",
			},
			Promise: map[string]string{
				"3년약정": "64",
			},
			Sale: map[string]string{
				"없음":    "274",
				"UHD셋탑": "311",
				"구글셋탑":  "327",
			},
			Service: map[string]string{
				"없음":    "36",
				"총1대설치": "106",
				"총2대설치": "107",
				"총3대설치": "108",
				"총4대설치": "109",
			},
		},
	}
}

func newLgHelloVision() *CommCompany {
	comm := &CommCompany{
		CompanyCode: "740",
		AreaCode:    "23",
		Article:     getLgHelloVisionArticle(),
	}
	return comm
}

func getLgHelloVisionArticle() map[string]*Article {
	return map[string]*Article{
		"인터넷": {
			Article: "744",
			Options: map[string]string{
				"광랜라이트(100M)": "1848",
				"광랜(160M)":    "1849",
				"기가라이트(500M)": "1850",
				"플래티넘기가(1G)":  "1851",
			},
			Promise: map[string]string{
				"3년약정": "93",
			},
			Sale: map[string]string{
				"없음": "313",
			},
			Service: map[string]string{
				"와이파이신청":   "95",
				"와이파이신청안함": "96",
			},
		},
		"TV": {
			Article: "745",
			Options: map[string]string{
				"베이직 UHD":  "1852",
				"프리미엄 UHD": "1853",
			},
			Promise: map[string]string{
				"3년약정": "94",
			},
			Sale: map[string]string{
				"없음": "314",
			},
			Service: map[string]string{
				"베이직HD 1대더추가설치": "129",
				"이코노미HD1대더추가설치": "130",
				"추가설치없음":        "131",
			},
		},
		"전화": {
			Article: "746",
			Options: map[string]string{
				"표준+(1,100원)": "1735",
			},
			Promise: map[string]string{
				"3년약정": "95",
			},
			Sale: map[string]string{
				"없음": "315",
			},
			Service: map[string]string{
				"없음": "99",
			},
		},
	}
}

func newSkylife() *CommCompany {
	comm := &CommCompany{
		CompanyCode: "739",
		AreaCode:    "24",
		Article:     getSkylifeArticle(),
	}
	return comm
}

func getSkylifeArticle() map[string]*Article {
	return map[string]*Article{
		"인터넷": {
			Article: "747",
			Options: map[string]string{
				"100M": "1739",
				"200M": "1740",
				"500M": "1846",
				"1G":   "1847",
			},
			Promise: map[string]string{
				"3년약정": "96",
			},
			Sale: map[string]string{
				"없음":  "319",
				"홈결합": "337",
			},
			Service: map[string]string{
				"일반와이파이신청": "126",
				"기가와이파이신청": "127",
				"신청안함":     "128",
			},
		},
		"TV": {
			Article: "748",
			Options: map[string]string{
				"Blue A+(218CH)":       "1744",
				"Green A+(195CH)":      "1745",
				"Green A+(30%)(195CH)": "1746",
			},
			Promise: map[string]string{
				"3년약정": "97",
			},
			Sale: map[string]string{
				"없음":    "322",
				"UHD셋탑": "326",
			},
			Service: map[string]string{
				"없음":    "101",
				"총1대설치": "122",
				"총2대설치": "123",
				"총3대설치": "124",
				"총4대설치": "125",
			},
		},
	}
}

func newRental() *CommCompany {
	comm := &CommCompany{
		CompanyCode: "743",
		AreaCode:    "25",
		Article:     getRentalArticle(),
	}
	return comm
}

func getRentalArticle() map[string]*Article {
	return map[string]*Article{
		"미기입": {
			Article: "761",
			Options: map[string]string{
				"미기입": "1855",
			},
			Promise: map[string]string{
				"미기입": "132",
			},
			Sale: map[string]string{
				"미기입": "339",
			},
			Service: map[string]string{
				"미기입": "132",
			},
		},
	}
}
