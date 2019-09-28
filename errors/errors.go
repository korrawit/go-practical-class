package main

import (
	"errors"
	"fmt"
)

var errA = errors.New("Err A")

type BusinessError struct {
	err error
}

func (b BusinessError) Error() string {
	return b.err.Error()
}

func main() {
	errBiz := BusinessError{
		err: errA,
	}
	fmt.Println(errBiz)
}
