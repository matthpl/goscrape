package goscrape

import (
	"io"
	"io/ioutil"
)

type Matcher interface {
	FindMatch(sources Sources, match Matches) ([]string, error)
}

type ScrapeMatcher struct {
	client *HttpClient
}

func NewScrapeMatcher() (*ScrapeMatcher, error) {
	c, _ := NewHttpClient()
	return &ScrapeMatcher{
		client: c,
	}, nil
}

func processMatch(resp io.ReadCloser, rules Matches) ([]string, error) {
	var res []string
	defer resp.Close()
	body, _ := ioutil.ReadAll(resp)
	res = append(res, string(body))
	// TODO: actual process it
	return res, nil
}
func (sm ScrapeMatcher) FindMatch(sources Sources, match Matches) ([]string, error) {
	var res []string
	for _, u := range sources.Urls {
		resp, _ := sm.client.Fetch(u.Method, u.Url)
		str, _ := processMatch(resp, match)
		res = append(str, res...)
	}
	return res, nil
}
