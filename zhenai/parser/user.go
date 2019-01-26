package parser

import (
	"regexp"

	"github.com/mozyy/crawler/engin"
	"github.com/mozyy/crawler/model"
)

var (
	userBaseRe = regexp.MustCompile(
		`<div class="des f-cl"[^>]*>([^<\s]+)\s\|\s([^<\s]+)\s\|\s([^<\s]+)\s\|\s([^<\s]+)\s\|\s([^<\s]+)\s\|\s([^<]+)</div>`)
	// userOtherBe = regexp.MustCompile()
)

// User is user list parser
func User(b []byte, name string) engin.Result {
	r := userBaseRe.FindAllSubmatch(b, -1)
	var result engin.Result
	if len(r) < 1 {
		return result
	}
	v := r[0]
	//长春 | 29岁 | 硕士 | 未婚 | 168cm | 8001-12000元
	result.Items = append(result.Items, model.User{
		Name:      name,
		Address:   string(v[1]),
		Age:       string(v[2]),
		Education: string(v[3]),
		Marriage:  string(v[4]),
		Height:    string(v[5]),
		Income:    string(v[6]),
	})
	return result
}
