package types

import "log"

type Error_t struct {
	Code    int // Code error log
	Message string
	Result  bool
	Err     error
}

func (e *Error_t) Check(err error) {
	if err != nil {
		e.Code = 500
		e.Message = "Internal Server Error"
		e.Result = false
		e.Err = err
		log.Println(err)
	}
}

/* test check error */
/* func main() {
	var err Error_t
	err.Check(errors.New("test"))
	fmt.Println(err)
} */