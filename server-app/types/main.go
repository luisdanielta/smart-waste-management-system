package types

import (
	"log"
)

type Error_t struct {
	Code    int // Code error log
	Message string
	Result  bool
	ERR     error
	Type    string
}

func (e *Error_t) clean() { e = nil }

func (e *Error_t) Check() {
	if e.ERR != nil {
		e.Code = 500
		e.Message = e.ERR.Error()
		e.Result = false
		log.Printf("code: %d - message: %s - result: %t", e.Code, e.Message, e.Result)
		e.clean()
		return
	}
}

/* test function check */
/*
func main() {
	var err Error_t
	err.ERR = errors.New("error")
	err.Check()
}
*/
