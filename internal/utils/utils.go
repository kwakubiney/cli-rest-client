package utils


import (
	"bytes"
	"encoding/json"
	"net/http"
)


func MakeRequest(route string, port string, requestBody interface{}, method string) (*http.Response, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, route, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

