package toutf8

import (
	"bytes"
	"fmt"
	"io"

	"github.com/mattn/go-encoding"
	"github.com/saintfish/chardet"
	"golang.org/x/text/transform"
)

func ToUTF8(body []byte) ([]byte, error) {
	det := chardet.NewTextDetector()
	res, err := det.DetectBest(body)
	if err != nil {
		return nil, fmt.Errorf("toUTF8 error: %w", err)
	}
	enc := encoding.GetEncoding(res.Charset)
	if enc == nil {
		return nil, fmt.Errorf("fail to get encoding")
	}
	trans := enc.NewDecoder()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, transform.NewReader(bytes.NewReader(body), trans))
	if err != nil {
		return nil, fmt.Errorf("toUTF8 error: %w", err)
	}
	return buf.Bytes(), nil
}
