package parser

import (
	"regexp"

	"github.com/mozyy/crawler/engin"
)

var cityListRe = regexp.MustCompile(
	`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[\w]+)"[^>]*>([^<]+)</a>`)

// CityList is city list parser
func CityList(b []byte) engin.Result {
	r := cityListRe.FindAllSubmatch(b, -1)
	var result engin.Result
	for _, v := range r {
		result.Requests = append(result.Requests, engin.Request{
			URL:    string(v[1]),
			Parser: City,
		})
		result.Items = append(result.Items, string(v[2]))
	}
	return result
}
