package main

import (
	"github.com/moz0yy/crawler/engin"
	"github.com/moz0yy/crawler/zhenai/parser"
)

func main() {
	entry := engin.Request{
		URL:    "http://www.zhenai.com/zhenhun",
		Parser: parser.CityList,
	}
	engin.Run(entry)
}
