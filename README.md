# Hydralisk
ERP  사이트 API v1

Usage
```go
echo '
{
  "admin_info": {
    "jupsu_id": "<id>",
    "yuchi_id": "<id>"
  },
  "customer_info": {
    "c_name": "테스트1",
    "c_type": "C",
    "c_tel2": "010-000-0000",
    "g_code_course_idx": "0",
    "c_email": "test@test.net"
  },
  "payment_info": {
    "g_give_type": "A",
    "g_give_chk": "A",
    "account_transfer": {
      "c_bank_cd": "0"
    },
    "credit": {
      "c_card_cd": "0"
    }
  },
  "gift_account": {
    "g_sp_bank_code": "0"
  },
  "item_info": {
    "g_article_idx1": "735",
    "g_code_area": "4",
    "goods_cnt": "1",
    "first_item": {
      "g_article_idx2": "736",
      "g_option_idx": "1716",
      "g_code_promise1": "84",
      "g_code_sale": "301",
      "g_code_service1": "78",
      "g_article_cnt1": "1",
      "g_sp_give_type": "A"
    },
    "second_item": {
      "g_article_idx2_1": "0",
      "g_option_idx_1": "0",
      "g_code_promise2": "0",
      "g_code_sale1": "0",
      "g_code_service2": "0",
      "g_article_cnt2": "1"
    },
    "third_item": {
      "g_article_idx2_2": "0",
      "g_option_idx_2": "0",
      "g_code_promise3": "0",
      "g_code_sale2": "0",
      "g_code_service3": "0",
      "g_article_cnt3": "1"
    }
  }
}
' >> data.json

b, err := ioutil.ReadFile("./data.json")
if err != nil {
    panic(err)
}
res, err := http.NewRequest(url, "application/json", bytes.NewReader(b))
if err != nil {
    log.Println(err)
}

defer res.Body.Close()

var b []byte
res.Body.Read(b)
fmt.Println(string(b))
```

