package util

import (
	"encoding/json"
	"errors"
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
	_, _ = io.WriteString(p.Writer, message)
	logger.Errorf(message)
}

func (p *Result) Success(object any) {
	p.rpc(enum.STATUS_OK, "success")
	js, _ := json.Marshal(object)
	_, _ = io.WriteString(p.Writer, string(js))
	logger.Infof(string(js))
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
