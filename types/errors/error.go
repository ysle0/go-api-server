package errors

import "fmt"

const (
	NotFoundUser = iota
)

var errMsgs = map[int64]string{
	NotFoundUser: "not found user",
}

func Errorf(code int64, args ...any) error {
	if msg, ok := errMsgs[code]; ok {
		return fmt.Errorf("%s : %v", msg, args)
	}
	return fmt.Errorf("errors code not found")
}
