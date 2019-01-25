package main

import (
	"github.com/mozyy/crawler/engin"
	"github.com/mozyy/crawler/zhenai/parser"
)

func main() {
	entry := engin.Request{
		URL:    "http://www.zhenai.com/zhenhun",
		Parser: parser.CityList,
	}
	// entry := engin.Request{
	// 	URL: "http://album.zhenai.com/u/1667190776",
	// 	Parser: func(b []byte) engin.Result {
	// 		return parser.User(b, "name")
	// 	},
	// }
	engin.RunSync(entry)
	// body, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/shanghai")
	// if err == nil {
	// 	fmt.Println(string(body))
	// }
}
