package plantuml

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

const plantuml_map = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

var plantuml_encoding *base64.Encoding

func init() {
	plantuml_encoding = base64.NewEncoding(plantuml_map)
}

func Encode(text string) string {
	b := new(bytes.Buffer)
	w, _ := flate.NewWriter(b, flate.BestCompression)
	w.Write([]byte(text))
	w.Close()
	return plantuml_encoding.EncodeToString(b.Bytes())
}

func Decode(encoded string) (string, error) {
	compressed, err := plantuml_encoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	r := flate.NewReader(bytes.NewReader(compressed))
	b := bytes.NewBuffer(nil)
	if _, err := io.Copy(b, r); err != nil {
		return "", err
	}
	return b.String(), nil
}
