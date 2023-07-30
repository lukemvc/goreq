package goreq

import (
	"net"
	"crypto/tls"
)

type Request struct {
	Url string
	Headers map[string]string
	Data string
	Json interface{}
}

type RequestInfo struct {
	Url string
	UrlPrefix string
	BaseUrl string
	UrlPath string
	Method string
	Host string
	RequestString string
	Chunked bool
	ResponseBuffer []byte
	Conn net.Conn
	TlsConn *tls.Conn
}

type Response struct {
	Text string
	Raw string
	StatusCode int
	Headers string
	Size int
}




