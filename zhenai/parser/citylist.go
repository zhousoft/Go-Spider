package parser

import (
	"Go-Spider/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//ParseCityList 城市列表解析
func ParseCityList(contents []byte) engine.ParseResult {
	//``中间的内容禁止转义
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 2
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
		limit--
		if limit < 0 {
			break
		}
	}
	return result
}
