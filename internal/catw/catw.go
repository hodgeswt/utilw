package catw

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hodgeswt/utilw/pkg/argparse"
	"github.com/hodgeswt/utilw/pkg/logw"
)

type InvalidArguments struct {
	Arguments []argparse.Argument
	Message   string
}

func (it *InvalidArguments) Error() string {
	return fmt.Sprintf("%s: %v", it.Message, it.Arguments)
}

func GetArguments() []argparse.Argument {
	x := new(FileArg)
	return []argparse.Argument{x}
}

func Run(args []string) error {
	logw.Debugf("+catw.Run, args: %v", args)
	defer logw.Debugf("-catw.Run")

	arguments := GetArguments()
	parsed, err := argparse.Parse(args, arguments, true)

	if err != nil {
		return err
	}

	allValid := true
	invalid := []argparse.Argument{}
	for _, argument := range parsed {
		if !argument.Valid() {
			logw.Debugf("%v", argument)
			invalid = append(invalid, argument)
			allValid = false
		}
	}

	if !allValid {
		return &InvalidArguments{
			Arguments: invalid,
			Message:   "invalid catw arguments",
		}
	}

	return run(parsed)
}

func run(args map[string]argparse.Argument) error {
	f, err := os.Open(args[FILEARG].Value()[0])

	if err != nil {
		return err
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
