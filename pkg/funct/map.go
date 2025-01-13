package funct

func Map[T any, U any](input []T, mapper func(T) (U, error)) ([]U, error) {
	out := make([]U, len(input))

	for i, x := range input {
		y, err := mapper(x)

		if err != nil {
			return nil, err
		}

		out[i] = y
	}

	return out, nil
}
