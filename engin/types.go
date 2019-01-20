package engin

// Request is request struct
type Request struct {
	URL    string
	Parser func([]byte) Result
}

// Result is Request's result
type Result struct {
	Requests []Request
	Items    []interface{}
}

// NilParser nil parser
func NilParser(b []byte) Result {
	var r Result
	return r
}
