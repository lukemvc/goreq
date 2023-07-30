package goreq

import (
	"fmt"
	"strings"
	"errors"
	"net"
	"encoding/json"
)

// Functions that are used by all request methods
/////////////////////////////////////////////////

func makeRequest(requestInfo *RequestInfo) (*Response, error) {
	var n int
	var err error
	switch requestInfo.UrlPrefix{
		case "https":
			if err = initTls(requestInfo); err != nil {
				return nil, err
			}
			defer requestInfo.TlsConn.Close()

			n, err = tlsRequest(requestInfo)
			if err != nil {
				return nil, err
			}

		case "http":
			n, err = nonTlsRequest(requestInfo)
			if err != nil {
				return nil, err
			}
	
		default:
			return nil, errors.New("invalid url prefix, must be http/https")
	}
	
	response := &Response{
		Size: n,
		Raw: string(requestInfo.ResponseBuffer[:n]),
	}
	
	if err := parseRawResponse(response); err != nil {
		return nil, err
	}
	if requestInfo.Chunked {
		response.Text = response.Text[:len(response.Text) - 7]
	}
	return response, nil
}


// Set headers, used by all methods
func setupHeaders(r *Request, requestString string) string {
	for key, value := range r.Headers {
		requestString += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	return requestString
}


// Setting body for post and put 
func handlePostPut(r *Request, requestString string) (string, error) {
	requestString = setupHeaders(r, requestString)

	if r.Json != nil {
		jsonBytes, err := json.Marshal(r.Json)
		if err != nil {
			return "", err
		}
		requestString += "Content-Type: application/json\r\n"
		requestString += fmt.Sprintf("Content-Length: %d\r\n", len(jsonBytes))
		requestString += "\r\n"
		requestString += string(jsonBytes)
	} else if strings.TrimSpace(r.Data) != "" {
		requestString += "Content-Type: application/x-www-form-urlencoded\r\n"
		requestString += fmt.Sprintf("Content-Length: %d\r\n", len(r.Data))
		requestString += "\r\n"
		requestString += r.Data
	} else {
		requestString += "\r\n"
	}

	return requestString, nil
}

// Formatting the request based on method, preparing requestInfo for makeRequest(), settign up TCP connection
func (requestInfo *RequestInfo) Init(r *Request) error {
	if err := PrefixBaseAndPathOfUrl(requestInfo); err != nil {
		return err
	}
	var requestString string
	var err error
	switch requestInfo.Method {
		case "GET":
			requestString = fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\n", requestInfo.UrlPath, requestInfo.BaseUrl)
			requestString = setupHeaders(r, requestString)
			requestString += "\r\n"

		case "OPTIONS":
			requestString = fmt.Sprintf("OPTIONS %s HTTP/1.1\r\nHost: %s\r\n", requestInfo.UrlPath, requestInfo.BaseUrl)
			requestString = setupHeaders(r, requestString)
			requestString += "\r\n"

		case "DELETE":
			requestString = fmt.Sprintf("DELETE %s HTTP/1.1\r\nHost: %s\r\n", requestInfo.UrlPath, requestInfo.BaseUrl)
			requestString = setupHeaders(r, requestString)
			requestString += "\r\n"

		case "POST":
			requestString = fmt.Sprintf("POST %s HTTP/1.1\r\nHost: %s\r\n", requestInfo.UrlPath, requestInfo.BaseUrl)
			requestString, err = handlePostPut(r, requestString)
			if err != nil {
				return err
			}
		case "PUT":
			requestString = fmt.Sprintf("PUT %s HTTP/1.1\r\nHost: %s\r\n", requestInfo.UrlPath, requestInfo.BaseUrl)
			requestString, err = handlePostPut(r, requestString)
			if err != nil {
				return err
			}

		default:
			return errors.New("invalid request method")
	}

	requestInfo.RequestString = requestString
	requestInfo.ResponseBuffer = make([]byte, 100000)

	HostFromUrl(requestInfo)
	conn, err := net.Dial("tcp", requestInfo.Host)
	if err != nil {
		return err
    }
	requestInfo.Conn = conn
	return nil
}


