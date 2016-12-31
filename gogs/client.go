package gogs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	url    string       // address of service
	token  string       // access token
	client *http.Client // hhtp client
}

type genericResponse struct {
	Data json.RawMessage `json:"data"`
	Ok   bool            `json:"ok"`
}

const apiPrefix = "/api/v1"

func NewClient(url, token string) *Client {
	return &Client{
		url:    strings.TrimSuffix(url, "/"),
		token:  token,
		client: &http.Client{},
	}
}

func (c *Client) doRequest(method string, resource string, query map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.url+apiPrefix+resource, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+c.token)
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return c.client.Do(req)
}

func (c *Client) getResponse(method string, resource string, query map[string]string) ([]byte, error) {
	res, err := c.doRequest(method, resource, query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case 403:
		return nil, errors.New("Status 403")
	case 404:
		return nil, errors.New("Status 404")
	}
	return data, nil
}
