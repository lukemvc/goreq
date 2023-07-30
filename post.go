package goreq

import "strings"


func (r *Request) Post() (*Response, error) {
	url := strings.ToLower(r.Url)

	requestInfo := &RequestInfo{
		Url: url,
		Method: "POST",
	}
	
	err := requestInfo.Init(r)
	if err != nil {
		return nil, err
	}
	defer requestInfo.Conn.Close()	

	response, err := makeRequest(requestInfo)
	if err != nil {
		return nil, err
	}

	return response, nil
} 
