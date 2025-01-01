package argparse

import "errors"

type Iterator struct {
	values []string
	index  int
}

var Done = errors.New("Iterator complete")
var InsufficientItems = errors.New("Iterator had insufficient items")

func NewIterator(values []string) *Iterator {
	return &Iterator{
		values: values,
		index:  0,
	}
}

func (it *Iterator) HasNext() bool {
	return it.index < len(it.values)
}

func (it *Iterator) Next() (string, error) {
	if !it.HasNext() {
		return "", Done
	}

	v := it.values[it.index]
	it.index++

	return v, nil
}

func (it *Iterator) Take(count int) ([]string, error) {
	if it.index >= len(it.values) {
		return nil, Done
	}

	r := make([]string, count)

	for i := range count {
		v, err := it.Next()

		if err == Done {
			return nil, InsufficientItems
		} else if err != nil {
			return nil, err
		}

		r[i] = v
	}

	return r, nil
}

func (it *Iterator) TakeAll() ([]string, error) {
	return it.Take(len(it.values) - it.index - 1)
}
