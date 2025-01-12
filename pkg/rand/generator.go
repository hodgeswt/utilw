package rand

import "errors"

type Generator interface {
	Init(opts interface{}) error
	Next() int64
	At(i int64) int64
	Reset()
}

var (
	ErrInvalidGeneratorOptions = errors.New("ErrInvalidGeneratorOptions")
)
