package goreq

import "strings"

func (r *Request) Delete() (*Response, error) {
	url := strings.ToLower(r.Url)

	requestInfo := &RequestInfo{
		Url: url,
		Method: "DELETE",
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

