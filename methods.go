package goreq

import "strings"

func (r *Request) Get() (*Response, error) {
	url := strings.ToLower(r.Url)

	requestInfo := &RequestInfo{
		Url: url,
		Method: "GET",
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


func (r *Request) Options() (*Response, error) {
	url := strings.ToLower(r.Url)

	requestInfo := &RequestInfo{
		Url: url,
		Method: "OPTIONS",
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


func (r *Request) Put() (*Response, error) {
	url := strings.ToLower(r.Url)

	requestInfo := &RequestInfo{
		Url: url,
		Method: "PUT",
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
