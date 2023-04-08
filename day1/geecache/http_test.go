package geecache

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHTTPPool_Log(t *testing.T) {
	urlStr := "https://cong5.net/post/golang?name=张三&age=20&sex=1"
	parseUrl, _ := url.Parse(urlStr)
	escape := url.QueryEscape("group")

	fmt.Println(parseUrl)
	fmt.Println(escape)
}
