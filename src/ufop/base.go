package ufop

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"regexp"
	"strings"
)

type UfopRequest struct {
	Cmd string         `json:"cmd"`
	Src UfopRequestSrc `json:"src"`
}

type UfopRequestSrc struct {
	Url      string `json:"url"`
	MimeType string `json:"mimetype"`
	Fsize    int64  `json:"fsize"`
}

type UfopError struct {
	Request UfopRequest
	Error   string
}

func getParam(fromStr, pattern, key string) (value string) {
	keyRegx := regexp.MustCompile(pattern)
	matchStr := keyRegx.FindString(fromStr)
	value = strings.Replace(matchStr, key+"/", "", -1)
	return
}

func getParamDecoded(fromStr, pattern, key string) (value string, err error) {
	strToDecode := getParam(fromStr, pattern, key)
	decodedBytes, decodeErr := base64.URLEncoding.DecodeString(strToDecode)
	if decodeErr != nil {
		err = decodeErr
		return
	}
	value = string(decodedBytes)
	return
}

func md5Hex(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
