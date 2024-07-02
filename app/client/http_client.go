package client

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type HttpClient struct {
	client *resty.Client
}

func NewHttpClient() *HttpClient {
	client := resty.New()
	client.SetTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	return &HttpClient{
		client: client,
	}
}

func (h *HttpClient) Get(url string, headers map[string]string) (int, []byte, error) {
	resp, err := h.client.R().
		SetHeaders(headers).
		Get(url)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}

func (h *HttpClient) Post(url string, headers map[string]string, body []byte) (int, []byte, error) {
	resp, err := h.client.R().
		SetHeaders(headers).
		SetBody(body).
		Post(url)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}

func (h *HttpClient) BuildHeaders(customHeaders map[string]string) map[string]string {
	headers := make(map[string]string)

	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	for name, value := range customHeaders {
		headers[name] = value
	}

	return headers
}
