package parser

import (
	"regexp"

	"github.com/mozyy/crawler/engin"
	"github.com/mozyy/crawler/model"
)

var userRe = regexp.MustCompile(
	`<div class="des f-cl"[^>]*>(\S+)\s\|\s(\S+)\s\|\s(\S+)\s\|\s(\S+)\s\|\s(\S+)\s\|\s([^<]+)</div>`)

//<div class="des f-cl" data-v-3c42fade="">上海 | 58岁 | 中专 | 离异 | 165cm | 3000元以下</div>

// User is city list parser
func User(b []byte, name string) engin.Result {
	r := userRe.FindAllSubmatch(b, -1)
	var result engin.Result
	for _, v := range r {
		item := model.User{
			Name:      name,
			Address:   string(v[1]),
			Age:       string(v[2]),
			Education: string(v[3]),
			Marriage:  string(v[4]),
			Height:    string(v[5]),
			Income:    string(v[6]),
		}
		result.Items = append(result.Items, item)
	}
	return result
}
