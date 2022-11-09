package util

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"

	strip "github.com/grokify/html-strip-tags-go"
)

func Stringify(data interface{}) string {
	dataByte, _ := json.MarshalIndent(data, "", " ")
	return string(dataByte)
}

func DecodeBase64String(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	byOfStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	r := []rune(str)
	identifier := string(r[0:2])

	if identifier == "/g" || identifier == "QQ" {
		b := bytes.Buffer{}
		b.Write(byOfStr)

		enconder := gob.NewDecoder(&b)
		err = enconder.Decode(&byOfStr)
		if err != nil {
			return "", err
		}
	}

	return string(byOfStr), nil
}

func RemoveHtmlTag(s string) string {

	s = strip.StripTags(s)

	return s
}
