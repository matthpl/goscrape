package goscrape

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Source struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}
type Sources struct {
	Urls []Source `json:"urls"`
}

type Matches struct {
	Patterns []string `json:"patterns"`
}

func (m Matches) ToString() string {
	str, _ := toJson(m)
	return str
}

func (s Sources) ToString() string {
	str, _ := toJson(s)
	return str
}

func toJson(s interface{}) (string, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func GetSourcesConfig(path string) (Sources, error) {
	var s Sources
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return s, err
	}
	json.Unmarshal(raw, &s)
	return s, nil
}

func GetMatchesConfig(path string) (Matches, error) {
	var s Matches
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return s, err
	}
	json.Unmarshal(raw, &s)
	return s, nil
}
