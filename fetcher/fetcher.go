package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("err http code status: %d", resp.StatusCode)
	}

	// 强制gbk转换为utf8
	/*utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	byt, err := ioutil.ReadAll(utf8Reader)*/

	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 自动检测内容的编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	byt, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(byt, "")
	return e
}
