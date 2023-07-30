# Goreq

A Golang package for http/https requests, built using sockets ("net")

---

## GET, OPTIONS, DELETE
- Headers are optional
```golang

func main() {

    url := "https://example.com"
    headers := map[string]string{
        "origin": "https://example.com",
        "referer": "https://example.com",
    }

    // Construct the request
    request := &Request{
        Url: url,
        Headers: headers,
    }

    // Perform the request and check for errors
    response, err := request.get() // OR request.Options() OR request.Delete()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Access different values from the response
    fmt.Printf("Status Code: %d | Response Body: %s | Raw Response: %q\n", response.StatusCode, response.Text, response.Raw)
}
```
<br>

## POST, PUT
- Support for url-encoded and application/json payloads
- Headers are optional, Content-Type and Content-Length will be set by default depending on the payload

### - application/url-encoded

```golang
func main() {
    
    url := "https://example.com"
    headers := map[string]string{
        "origin": "https://example.com",
        "referer": "https://example.com",
    }
    
    // URL ENCODED PAYLOAD
    data := "email=user@example.com&password=Password123!"

    // Construct the request
    request := &Request{
        Url: url,
        Headers: headers,
        Data: data,
    }

    // Perform the request and check for errors
    response, err := request.Post() // OR request.Put()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Access different values from the response
    fmt.Printf("Status Code: %d | Response Body: %s | Raw Response: %q\n", response.StatusCode, response.Text, response.Raw)
}
```

### - application/json

```golang

func main() {

    url := "https://example.com"
    headers := map[string]string{
        "origin": "https://example.com",
        "referer": "https://example.com",
    }

    // APPLICATION JSON PAYLOAD
    jsonPayload := map[string]interface{}{
        "UserID": 10,
        "UserName": "Tommy",
        "Orders": map[int]string{
            1: "ND2L3234LN",
            2: "LADKNS234F",
            3: "KJNADL123K",
        },
    }

    // Construct request with Json payload
    jsonRequest := &Request{
        Url: url,
        Headers: headers,
        Json: jsonPayload,
    }

    // Perform request and check for errors
    jsonResponse, err := jsonRequest.Post() // OR jsonRequest.Put()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Access different values from the response
    fmt.Printf("Status Code: %d | Response Body: %s | Raw Response: %q\n", jsonResponse.StatusCode, jsonResponse.Text, jsonResponse.Raw)
}

```

