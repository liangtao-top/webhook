package util

import (
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
