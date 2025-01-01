package catw

import (
	"errors"
	"fmt"

	"github.com/hodgeswt/utilw/pkg/argparse"
)

var FILEARG = "filename"

type FileArg struct {
	value  string
	parsed bool
}

func (it *FileArg) Name() string {
	return FILEARG
}

func (it *FileArg) Value() []string {
	return []string{it.value}
}

func (it *FileArg) Parameters() int {
	return 1
}

func (it *FileArg) IsRequired() bool {
	return true
}

func (it *FileArg) Parsed() bool {
	return it.parsed
}

func (it *FileArg) Valid() bool {
	return it.parsed && it.value != ""
}

func (it *FileArg) Parse(arg string, data ...string) error {
	if !(arg == "-f" || arg == "--file") {
		return argparse.NoMatch
	}

	if len(data) != 1 {
		m := fmt.Sprintf("Invalid arguments provided for filename: %v", data)
		return errors.New(m)
	}

	it.value = data[0]
	it.parsed = true

	return nil
}

func (it *FileArg) String() string {
	m := "{Argument: %s, Required: %v, Parameters: %v"

	if it.Parsed() {
		m = m + ", Value: %v}"
		return fmt.Sprintf(m, it.Name(), it.IsRequired(), it.Parameters(), it.Value())
	} else {
		m = m + "}"
		return fmt.Sprintf(m, it.Name(), it.IsRequired(), it.Parameters())
	}
}
