package parser

import (
	"Go-Spider/engine"
	"regexp"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//ParseCity 城市解析
func ParseCity(contents []byte) engine.ParseResult {
	//``中间的内容禁止转义
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//将name提取出来，因为在下面的闭包函数中需要传入name值
		//如果不将name提取出来,直接传入m[2]的话，会导致所有函数
		//共享一份m，结果时所有用户名都一样
		name := string(m[2])
		result.Items = append(result.Items, "User: "+name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, name)
				},
			})
	}
	return result
}
