package argparse

import "github.com/hodgeswt/utilw/pkg/logw"

type ArgumentParser struct {
	logger *logw.Logger
}

func NewArgumentParser(logger *logw.Logger) *ArgumentParser {
	var parser = new(ArgumentParser)
	parser.logger = logger

	if logger != nil {
		logger.LoadConfig(nil)
	}

	return parser
}

type InsufficientArguments struct {
	arguments []Argument
	message   string
}

func (it *InsufficientArguments) Error() string {
	return it.message
}

type InsufficientParameters struct {
	argument Argument
	message  string
}

func (it *InsufficientParameters) Error() string {
	return it.message
}

func (it *ArgumentParser) Parse(args []string, arguments []Argument, ignoreFirst bool) (map[string]Argument, error) {
	it.logger.Debugf("+argparse.Parse, args: %v, arguments: %v, ignoreFirst: %v", args, arguments, ignoreFirst)
	defer it.logger.Debug("-argparse.Parse")

	iter := NewIterator(args)

	first := true

	it.logger.Debug("starting iterator")
	for {
		it.logger.Debug("iteration")

		arg, err := iter.Next()

		if err == Done {
			it.logger.Debug("iterator done")
			break
		} else if err != nil {
			it.logger.Errorf("iterator error: %v", err)
			return nil, err
		}

		if ignoreFirst && first {
			it.logger.Debug("ignoring first")
			first = false
			continue
		}

		allParsed := true

		for _, argument := range arguments {
			if argument.Parsed() {
				continue
			}

			it.logger.Debugf("trying cli arg %s for argument %v", arg, argument)

			allParsed = false

			if argument.Parameters() > 0 {
				params, err := iter.Take(argument.Parameters())

				if err != nil {
					return nil, &InsufficientParameters{
						argument: argument,
						message:  "Insufficient parameters were provided",
					}
				}

				err = argument.Parse(arg, params...)

			} else {
				err = argument.Parse(arg)
			}

			if err != NoMatch && err != nil {
				it.logger.Debugf("unexpected error in utilw.argparse.Parse %v", err)
				return nil, err
			}
		}

		if allParsed {
			break
		}
	}

	if iter.HasNext() {
		remaining, err := iter.TakeAll()
		message := "unexpected cli arguments provided"

		if err != nil {
			it.logger.Debugf(message)
			return nil, &UnexpectedArguments{
				Values:  []string{},
				Message: message,
			}
		}

		it.logger.Errorf(message+": %v", remaining)
		return nil, &UnexpectedArguments{
			Values:  remaining,
			Message: message,
		}
	}

	allParsed := true
	out := map[string]Argument{}
	missing := []Argument{}

	for _, argument := range arguments {
		if !argument.Parsed() {
			missing = append(missing, argument)
			it.logger.Debugf("Argument not parsed: %v", argument)
			allParsed = false
		} else {
			out[argument.Name()] = argument
		}
	}

	if !allParsed {
		it.logger.Debugf("insuffcient cli arguments provided")
		return nil, &InsufficientArguments{
			arguments: missing,
			message:   "insufficient cli arguments provided",
		}
	}

	return out, nil
}
