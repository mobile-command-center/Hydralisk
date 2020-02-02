package user

import (
	"bytes"
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
			_, _ = io.Copy(fw, strings.NewReader(value))
		}
	}
	w.Close()

	return b.String(), w.FormDataContentType()
}
