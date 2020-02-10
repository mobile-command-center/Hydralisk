package user

import (
	"bytes"
	"golang.org/x/text/encoding/korean"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
)

//makeMultiPart 함수는 POST 요청시 multipart 요청 데이터를 만드는 함수이다.
func makeMultiPart(v url.Values) (string, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.SetBoundary("ajnetbot")

	for key, r := range v {
		var fw io.Writer
		fw, _ = w.CreateFormField(key)
		for _, value := range r {
			//value type string. to make value as EUCKR
			euckr := korean.EUCKR.NewEncoder()
			s, _ := euckr.String(value)
			_, _ = io.Copy(fw, strings.NewReader(s))
		}
	}
	w.Close()

	return b.String(), w.FormDataContentType()
}
