package argparse

import (
	"errors"
	"fmt"
)

var NoMatch = errors.New("The provided data did not match this argument")

type Argument interface {
	Name() string
	Value() []string
	Parameters() int
	IsRequired() bool
	Parsed() bool
	String() string
	Valid() bool
	Parse(argument string, data ...string) error
}

type UnexpectedArguments struct {
	Values  []string
	Message string
}

func (it *UnexpectedArguments) Error() string {
	return fmt.Sprintf("%s: %v", it.Message, it.Values)
}
