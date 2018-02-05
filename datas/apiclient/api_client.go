package apiclient

import (
	"log"
	"net/http"
	"net/url"
)

type ApiClient struct {
	BaseUrl    string
	HTTPClient *http.Client
	Logger     *log.Logger
}

type Response struct {
	Response *http.Response
}

func NewApiClient(url string, logger *log.Logger) *ApiClient {
	return &ApiClient{
		BaseUrl:    url,
		HTTPClient: &http.Client{},
		Logger:     logger,
	}
}

func (client *ApiClient) GetRequest(querys map[string]string) *http.Response {

	// Request生成
	req, reqerr := http.NewRequest("GET", client.BaseUrl, nil)
	if reqerr != nil {
		log.Fatal(reqerr)
	}

	// get query 生成
	values := url.Values{}
	for k, v := range querys {
		values.Add(k, v)
	}
	req.URL.RawQuery = values.Encode()

	// HttpRequest
	res, doerr := client.HTTPClient.Do(req)
	if doerr != nil {
		log.Fatal(doerr)
	}

	return res
}
