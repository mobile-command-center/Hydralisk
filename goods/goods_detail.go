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
		"인터넷": &Article{
			Article: "730",
			Options: map[string]string{
				"100M": "1668",
				"500M": "1669",
				"1G":   "1670",
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
		"IPTV": &Article{
			Article: "732",
			Options: map[string]string{
				"슬림(230CH)":  "1677",
				"라이트(240CH)": "1678",
				"에센스(270CH)": "1679",
				"엔터(270CH)":  "1680",
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
		"전화": &Article{
			Article: "733",
			Options: map[string]string{
				"번호이동":    "1694",
				"신규가입":    "1695",
				"070번호이동": "1712",
				"070신규":   "1713",
			},
			Sale: map[string]string{
				"": "",
			},
			Promise: map[string]string{
				"3년약정": "79",
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
		"인터넷": &Article{
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
		"전화": &Article{
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
		"BTV": &Article{
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
		"인터넷": &Article{
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
		"전화": &Article{
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
				"": "",
			},
			Service: map[string]string{
				"없음": "79",
			},
		},
		"BTV": &Article{
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
		"인터넷": &Article{
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
		"전화": &Article{
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
				"없음":    "287",
				"★패밀리★": "328",
			},
			Service: map[string]string{
				"없음": "76",
			},
		},
		"TV": &Article{
			Article: "86",
			Options: map[string]string{
				"베이직(185CH)":           "141",
				"프리미엄(224CH)":          "142",
				"프리미엄 넷플릭스 HD(224CH)":  "1728",
				"프리미엄 넷플릭스 UHD(224CH)": "1842",
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
		"인터넷": &Article{
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
		"TV": &Article{
			Article: "745",
			Options: map[string]string{
				"이코노미(105CH)": "1732",
				"베이직(185CH)":  "1733",
				"프리미엄(213CH)": "1734",
			},
			Promise: map[string]string{
				"3년약정": "94",
			},
			Sale: map[string]string{
				"UHD셋탑": "314",
				"HD셋탑":  "316",
			},
			Service: map[string]string{
				"베이직HD 1대더추가설치": "129",
				"이코노미HD1대더추가설치": "130",
				"추가설치없음":        "131",
			},
		},
		"전화": &Article{
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
		"인터넷": &Article{
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
		"TV": &Article{
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
		"기존KT": &Article{
			Article: "749",
			Options: map[string]string{
				"기존KT사용": "1856",
				"해당없음":   "1857",
			},
			Promise: map[string]string{
				"3년약정": "98",
			},
			Sale: map[string]string{
				"없음": "323",
			},
			Service: map[string]string{
				"없음": "100",
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
		"SK매직": &Article{
			Article: "750",
			Options: map[string]string{
				"냉온정수기":  "1747",
				"얼음정수기":  "1748",
				"공기청정기":  "1749",
				"비데":     "1750",
				"주방가전용품": "1751",
				"전기레인지":  "1752",
				"의류건조기":  "1753",
				"안마의자":   "1754",
				"제빙기":    "1825",
			},
			Promise: map[string]string{
				"3년약정": "99",
				"4년약정": "100",
				"5년약정": "101",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"청호나이스": &Article{
			Article: "751",
			Options: map[string]string{
				"커피정수기":  "1755",
				"얼음정수기":  "1756",
				"냉온정수기":  "1757",
				"공기청정기":  "1758",
				"비데":     "1759",
				"주방가전용품": "1760",
				"연수기":    "1761",
				"제빙기":    "1829",
			},
			Promise: map[string]string{
				"3년약정": "102",
				"4년약정": "103",
				"5년약정": "104",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"LG전자": &Article{
			Article: "752",
			Options: map[string]string{
				"냉온정수기":      "1762",
				"공기청정기":      "1763",
				"주방가전용품":     "1764",
				"의류건조기":      "1765",
				"안마의자":       "1766",
				"전기레인지":      "1767",
				"의류관리기":      "1768",
				"홈브루(맥주제조기)": "1826",
				"식기세척기":      "1827",
			},
			Promise: map[string]string{
				"3년약정": "105",
				"4년약정": "106",
				"5년약정": "107",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"쿠쿠": &Article{
			Article: "753",
			Options: map[string]string{
				"펫용품":      "1769",
				"얼음정수기":    "1770",
				"냉온정수기":    "1771",
				"공기청정기":    "1772",
				"비데":       "1773",
				"연수기":      "1774",
				"안마의자":     "1775",
				"전기레인지":    "1776",
				"매트리스&프레임": "1777",
			},
			Promise: map[string]string{
				"3년약정": "108",
				"4년약정": "109",
				"5년약정": "110",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"현대렌탈케어": &Article{
			Article: "754",
			Options: map[string]string{
				"얼음정수기": "1778",
				"냉온정수기": "1779",
				"공기청정기": "1780",
				"비데":    "1781",
			},
			Promise: map[string]string{
				"3년약정": "111",
				"4년약정": "112",
				"5년약정": "113",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"바디프랜드": &Article{
			Article: "755",
			Options: map[string]string{
				"냉온정수기":    "1782",
				"매트리스&프레임": "1783",
			},
			Promise: map[string]string{
				"3년약정": "114",
				"4년약정": "115",
				"5년약정": "116",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"교원웰스": &Article{
			Article: "756",
			Options: map[string]string{
				"냉온정수기":  "1784",
				"공기청정기":  "1785",
				"비데":     "1786",
				"연수기":    "1787",
				"주방가전용품": "1788",
				"의류건조기":  "1789",
				"안마의자":   "1790",
				"전기레인지":  "1791",
				"의류관리기":  "1792",
			},
			Promise: map[string]string{
				"3년약정": "117",
				"4년약정": "118",
				"5년약정": "119",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"코웨이": &Article{
			Article: "757",
			Options: map[string]string{
				"얼음정수기":    "1793",
				"냉온정수기":    "1794",
				"공기청정기":    "1795",
				"비데":       "1796",
				"연수기":      "1797",
				"전기레인지":    "1798",
				"의류관리기":    "1799",
				"매트리스&프레임": "1800",
			},
			Promise: map[string]string{
				"3년약정": "120",
				"4년약정": "121",
				"5년약정": "122",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"현대렌탈서비스": &Article{
			Article: "758",
			Options: map[string]string{
				"펫용품":      "1801",
				"냉온정수기":    "1806",
				"공기청정기":    "1807",
				"비데":       "1808",
				"연수기":      "1809",
				"주방가전용품":   "1810",
				"의류건조기":    "1811",
				"안마의자":     "1812",
				"전기레인지":    "1813",
				"생활가전용품":   "1814",
				"의류관리기":    "1815",
				"매트리스&프레임": "1816",
				"가구":       "1817",
			},
			Promise: map[string]string{
				"3년약정": "123",
				"4년약정": "124",
				"5년약정": "125",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"루헨스": &Article{
			Article: "759",
			Options: map[string]string{
				"냉온정수기": "1802",
				"공기청정기": "1804",
				"비데":    "1805",
			},
			Promise: map[string]string{
				"3년약정": "126",
				"4년약정": "127",
				"5년약정": "128",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"기타(렌탈)": &Article{
			Article: "760",
			Options: map[string]string{
				"커피머신":   "1818",
				"웰스팜":    "1819",
				"건조기":    "1820",
				"세탁기":    "1821",
				"에어드레서":  "1822",
				"탑퍼(토퍼)": "1823",
				"LED마스크": "1824",
				"에어컨":    "1828",
				"청소기":    "1830",
				"전기자전거":  "1831",
				"테블릿PC":  "1832",
				"복합기":    "1833",
				"전기장판":   "1834",
				"프로젝트빔":  "1835",
				"제습기":    "1836",
			},
			Promise: map[string]string{
				"3년약정": "129",
				"4년약정": "130",
				"5년약정": "131",
			},
			Sale: map[string]string{
				"": "",
			},
			Service: map[string]string{
				"": "",
			},
		},
		"미기입": &Article{
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
