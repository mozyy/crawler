package main

import (
	"github.com/mozyy/crawler/engin"
	"github.com/mozyy/crawler/zhenai/parser"
)

func main() {
	entry := engin.Request{
		URL:    "http://www.zhenai.com/zhenghun",
		Parser: parser.CityList,
	}
	engin.RunSync(entry)

}
