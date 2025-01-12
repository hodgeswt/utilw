package rand

type LinearCongruentialGenerator struct {
	Multiplier int64
	Increment  int64
	Modulus    int64
	Seed       int64
	cache      map[int64]int64
	curr       int64
	last       int64
}

type LinearCongruentialGeneratorOptions struct {
	Multiplier int64 `json:"multiplier"`
	Increment  int64 `json:"increment"`
	Modulus    int64 `json:"modulus"`
	Seed       int64 `json:"seed"`
}

// Creates and initializes a new LCG
func NewLinearCongruentialGenerator(opts interface{}) (*LinearCongruentialGenerator, error) {
	var lcg = new(LinearCongruentialGenerator)

	lcg.cache = map[int64]int64{}

	err := lcg.Init(opts)
	if err != nil {
		return nil, err
	}

	return lcg, nil
}

// Initializes a Linear Congruential Generator. Required parameters are present
// in type [LinearCongruentialGeneratorOptions](LinearCongruentialGeneratorOptions)
func (it *LinearCongruentialGenerator) Init(opts interface{}) error {
	cast, ok := opts.(LinearCongruentialGeneratorOptions)

	if !ok {
		return ErrInvalidGeneratorOptions
	}

	it.Multiplier = cast.Multiplier
	it.Increment = cast.Increment
	it.Modulus = cast.Modulus
	it.Seed = cast.Seed

	return nil
}

// Gets the next random number from the generator
func (it *LinearCongruentialGenerator) Next() int64 {
	v, ok := it.cache[it.curr]
	if ok {
		it.curr++
		it.last = v
		return v
	}

	x := (it.Multiplier*it.last + it.Increment) % it.Modulus
	it.cache[it.curr] = x

	it.curr++
	it.last = x
	return x
}

// Gets the i-th result of .Next() from the current state
func (it *LinearCongruentialGenerator) At(i int64) int64 {
	count := int64(0)
	for {
		if count == i {
			return it.curr
		}

		it.Next()
		count++
	}
}

// Resets the output of Next() to be the same for first call with same seed,
// i.e.:
// ```go
// it = &NewExampleGenerator{}
// it.Init(interface{}{seed: 0})
// it.Next() // returns 1
// it.Next() // returns 2
// it.Next() // returns 3
// it.Reset()
// it.Next() // returns 1 again
// ```
func (it *LinearCongruentialGenerator) Reset() {
	it.curr = 0
}
