package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const citylist = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(citylist)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 2

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser, //ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
