package goreq

import (
	"strings"
	"fmt"
	"errors"
	"bytes"
	"strconv"
)

// Helper functions mainly for parsing and formatting

func parseRawResponse(response *Response) error {
    statusAndHeaders, body := splitResponseAndBody(response.Raw)
    statusAndHeadersLines := strings.Split(statusAndHeaders, "\r\n")
    statusLine := statusAndHeadersLines[0]
    statusParts := strings.SplitN(statusLine, " ", 3)
    if len(statusParts) != 3 {
        return fmt.Errorf("invalid status line in raw response")
    }
    response.StatusCode, _ = strconv.Atoi(statusParts[1])
	response.Headers = statusAndHeaders
	response.Text = body
	return nil
}


func splitResponseAndBody(rawResponse string) (string, string) {
    parts := strings.SplitN(rawResponse, "\r\n\r\n", 2)
    if len(parts) == 2 {
        return parts[0], parts[1]
    }
    return parts[0], ""
}


func checkChunkedEnding(buffer []byte) bool {
	if bytes.HasSuffix(buffer, []byte("\r\n0\r\n\r\n")) {
		return true
	} else {
		return false
	}
}

func PrefixBaseAndPathOfUrl(r *RequestInfo) error {
	urlPrefix, baseUrlPath, err := splitUrl(r.Url)
	if err != nil {
		return err
	}
	baseUrl, path := splitBaseUrlAndPath(baseUrlPath)
	r.UrlPrefix = urlPrefix
	r.BaseUrl = baseUrl
	r.UrlPath = path
	return nil
}


func HostFromUrl(r *RequestInfo) {
	if StringContainsChar(r.BaseUrl, ':') {
		// no need for port, if baseUrl contains ':' it is most likely "127.0.0.1:5000" or a server ip and port
		r.Host = r.BaseUrl
	} else {
		// assign port depending on prefix
		if strings.HasPrefix(r.UrlPrefix, "https") {
			r.Host = r.BaseUrl + ":443"
		} else {
			r.Host = r.BaseUrl + ":80"
		}
	}
	return
}


func splitUrl(url string) (string, string, error) {
	splitList := strings.Split(url, "://")
	if len(splitList) != 2 {
		return "", "", errors.New("error parsing URL")
	}
	urlPrefix := splitList[0]
	baseUrlPath := splitList[1]
	return urlPrefix, baseUrlPath, nil
}


func splitBaseUrlAndPath(baseUrlPath string) (baseUrl, path string) {
	index := strings.Index(baseUrlPath, "/")
	if index != -1 {
		baseUrl = baseUrlPath[:index]
		path = baseUrlPath[index:]
	} else {
		baseUrl = baseUrlPath
		path = "/"
	}
	return baseUrl, path
}


func StringContainsChar(s string, char byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == char {
			return true
		}
	}
	return false
}
