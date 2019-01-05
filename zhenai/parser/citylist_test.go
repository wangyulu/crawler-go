package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	resultSize := 470
	parserRes := ParseCityList(content)
	if len(parserRes.Requests) != resultSize {
		t.Errorf("should: %d, actual: %d", resultSize, len(parserRes.Requests))
	}

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls {
		if parserRes.Requests[i].Url != url {
			t.Errorf("should url: %q, actual url:%q", parserRes.Requests[i].Url, url)
		}
	}
}
