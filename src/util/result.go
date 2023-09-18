package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"webhook/src/global/enum"
	"webhook/src/logger"
)

type Result struct {
	Writer http.ResponseWriter
}

func (p *Result) Error(status int, message string) {
	p.rpc(status, message)
	str := fmt.Sprintf("{\"status\":%d,\"msg\":%s}", status, message)
	_, _ = io.WriteString(p.Writer, str)
	logger.Errorf(str)
}

func (p *Result) Success(object any) {
	p.rpc(enum.STATUS_OK, "success")
	jsonByte, _ := json.Marshal(object)
	str := fmt.Sprintf("{\"status\":%d,\"msg\":%s}", enum.STATUS_OK, string(jsonByte))
	_, _ = io.WriteString(p.Writer, str)
	logger.Infof(str)
}

func (p *Result) Pack(data []byte) []byte {
	l := len([]rune(string(data)))
	l1 := []byte(strconv.Itoa(l))
	var l2 []byte
	for i := 0; i < 5; i++ {
		if i < len(l1) {
			l2 = append(l2, l1[i])
		} else {
			l2 = append(l2, 0)
		}
	}
	return append(l2, data...)
}

func (p *Result) Unpack(data []byte) ([]byte, error) {
	if len(data) < 5 {
		return nil, errors.New("仅支持RPC协议")
	}
	return data[5:], nil
}

func (p *Result) rpc(status int, message string) {
	p.Writer.Header().Add("content-type", "application/json; charset=UTF-8")
	p.Writer.Header().Add("x-status", strconv.Itoa(status))
	p.Writer.Header().Add("x-message", message)
	p.Writer.Header().Add("server", runtime.Version())
}
