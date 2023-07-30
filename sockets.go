package goreq

import (
	"crypto/tls"
	"strings"
)

func initTls(r *RequestInfo) error {
	tlsConfig := &tls.Config{
		ServerName: r.BaseUrl,
	}

	tlsConn := tls.Client(r.Conn, tlsConfig)
	if err := tlsConn.Handshake(); err != nil {
		return err
	}
	r.TlsConn = tlsConn
	return nil
}


func tlsRequest(rinfo *RequestInfo) (int, error) {
    _, err := rinfo.TlsConn.Write([]byte(rinfo.RequestString))
    if err != nil {
        return 0, err
    }

	totalRead, err := rinfo.TlsConn.Read(rinfo.ResponseBuffer)
	if err != nil {
		return totalRead, err
	}
	chunkCheck := strings.Contains(strings.ToLower(string(rinfo.ResponseBuffer[:totalRead])), "encoding: chunked")
	if chunkCheck {
		rinfo.Chunked = true
		if !checkChunkedEnding(rinfo.ResponseBuffer[:totalRead]) {
			for {
				n, err := rinfo.TlsConn.Read(rinfo.ResponseBuffer[totalRead:])
				if err != nil {
					return totalRead, err
				}
				totalRead += n
				if checkChunkedEnding(rinfo.ResponseBuffer[:totalRead]) {
					break
				}
			}
		}
	}
    return totalRead, nil
}


func nonTlsRequest(rinfo *RequestInfo) (int, error) {
    _, err := rinfo.Conn.Write([]byte(rinfo.RequestString))
    if err != nil {
        return 0, err
    }

	totalRead, err := rinfo.Conn.Read(rinfo.ResponseBuffer)
	if err != nil {
		return totalRead, err
	}
	chunkCheck := strings.Contains(strings.ToLower(string(rinfo.ResponseBuffer[:totalRead])), "encoding: chunked")
	if chunkCheck {
		rinfo.Chunked = true
		if !checkChunkedEnding(rinfo.ResponseBuffer[:totalRead]) {
			for {
				n, err := rinfo.Conn.Read(rinfo.ResponseBuffer[totalRead:])
				if err != nil {
					return totalRead, err
				}
				totalRead += n
				if checkChunkedEnding(rinfo.ResponseBuffer[:totalRead]) {
					break
				}
			}
		}
	}
	return totalRead, nil
}
