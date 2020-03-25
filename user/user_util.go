package user

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"strings"
	"time"

	"golang.org/x/text/encoding/korean"
)

//makeMultiPart 함수는 POST 요청시 multipart 요청 데이터를 만드는 함수이다.
func makeMultiPart(v url.Values, user UserData) (string, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.SetBoundary("ajnetbot")

	for key, r := range v {
		var fw io.Writer
		switch key {
		case "g_file1":
			h := make(textproto.MIMEHeader)
			loc, _ := time.LoadLocation("Asia/Seoul")
			now := time.Now().In(loc)
			euckr := korean.EUCKR.NewEncoder()
			temp := fmt.Sprintf("%s_%d%02d%02d_%02d%02d%02d.txt", user.Filename, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
			filename, _ := euckr.String(temp)
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, key, filename))
			h.Set("Content-Type", "text/plain")
			part, _ := w.CreatePart(h)
			part.Write(user.Data.Bytes())
		default:
			fw, _ = w.CreateFormField(key)
			for _, value := range r {
				//value type string. to make value as EUCKR
				euckr := korean.EUCKR.NewEncoder()
				s, _ := euckr.String(value)
				_, _ = io.Copy(fw, strings.NewReader(s))
			}
		}
	}
	w.Close()
	return b.String(), w.FormDataContentType()
}
