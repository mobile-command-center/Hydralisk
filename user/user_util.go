package user

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"strings"

	"golang.org/x/text/encoding/korean"
)

//makeMultiPart 함수는 POST 요청시 multipart 요청 데이터를 만드는 함수이다.
func makeMultiPart(v url.Values, rawData bytes.Buffer) (string, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.SetBoundary("ajnetbot")

	for key, r := range v {
		var fw io.Writer
		if key == "g_file1" {
			part, _ := w.CreateFormFile(key, "origin.txt")
			/*
				h := make(textproto.MIMEHeader)
				h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,	key, "origin.txt"))
				h.Set("Content-Type", "text/plain")
				part, _ := w.CreatePart(h)
			*/
			part.Write(rawData.Bytes())
		} else {
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
