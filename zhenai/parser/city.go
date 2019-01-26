package parser

import (
	"regexp"

	"github.com/mozyy/crawler/engin"
)

var cityRe = regexp.MustCompile(
	`<a href="(http://album.zhenai.com/u/\d+)" target="_blank">([^<]+)</a>`)

// City is city list parser
func City(b []byte) engin.Result {
	r := cityRe.FindAllSubmatch(b, -1)
	var result engin.Result
	for _, v := range r {
		name := string(v[2])
		result.Requests = append(result.Requests, engin.Request{
			URL: string(v[1]),
			Parser: func(b []byte) engin.Result {
				return User(b, name)
			},
		})
		result.Items = append(result.Items, name)
	}
	return result
}
