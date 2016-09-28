package goscrape

import (
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Client interface {
	Fetch(method string, url string) (io.ReadCloser, error)
}

type HttpClient struct {
	client *http.Client
}

func (hc *HttpClient) Fetch(method string, url string) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func NewHttpClient() (*HttpClient, error) {
	opts := &cookiejar.Options{PublicSuffixList: publicsuffix.List}
	jar, err := cookiejar.New(opts)
	if err != nil {
		return nil, err
	}
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Jar:       jar,
		Timeout:   time.Second * 10,
		Transport: netTransport}

	ret := &HttpClient{
		client: client,
	}
	return ret, nil
}
