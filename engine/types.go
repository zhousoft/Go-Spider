package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//NilParser 返回空解析结果，占位用
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
