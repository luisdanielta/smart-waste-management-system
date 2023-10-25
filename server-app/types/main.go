package types

type Error_t struct {
	Code    int // Code error log
	Message string
	Result  bool
	Err     error
}
