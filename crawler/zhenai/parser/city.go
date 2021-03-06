package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser, /*func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(m[2]))
			}*/
		})
	}
	return result
}
