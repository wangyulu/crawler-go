package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserRes
}

type ParserRes struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParserRes {
	return ParserRes{}
}
