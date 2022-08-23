package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/util"
)

func ToJsonString(object any, format ...bool) string {
	b := true
	if len(format) > 0 {
		b = format[0]
	}
	if b {
		js, _ := json.MarshalIndent(object, "", "\t")
		return string(js)
	} else {
		js, _ := json.Marshal(object)
		return string(js)
	}
}

func LocalIP() string {
	return util.LocalIP()
}

func GenHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	//sha := hex.EncodeToString(h.Sum(nil))
	//fmt.Printf("sha:%s\n", sha)
	return Base64UrlSafeEncode(h.Sum(nil))
}

func Base64UrlSafeEncode(source []byte) string {
	byteArr := base64.StdEncoding.EncodeToString(source)
	//safeUrl := strings.Replace(string(byteArr), "/", "_", -1)
	//safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	//safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return string(byteArr)
}
