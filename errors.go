package ourjson

import (
	"fmt"
	"errors"
)

var (
	IndexOutOfRangeError = errors.New("index out of range")
	ValueNotNumberError = errors.New("value is not number")
)

type KeyNotFoundError struct {
	Key string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("key[%s] not existed", e.Key)
}

type ValueTransformTypeError struct {
	Type string
}

func (e ValueTransformTypeError) Error() string {
	return fmt.Sprintf("cannot transform into %s", e.Type)
}
