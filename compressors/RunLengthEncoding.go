package compressors

import (
	"errors"
	"fmt"
)

type RunLengthEncoding struct{}

func (r RunLengthEncoding) Encode(in []byte) []byte {
	buffer := make([]byte, 0, 2_000_000)
	fmt.Printf("The size of buffer, %d\n", len(buffer))
	curr := in[0]
	freq := 0

	for _, v := range in {
		if curr == v {
			freq++
		} else {
			buffer = append(buffer, curr)
			buffer = append(buffer, byte(freq))
			curr = v
			freq = 1
		}
	}
	buffer = append(buffer, curr)
	buffer = append(buffer, byte(freq))
	return buffer
}

func (r RunLengthEncoding) Decode(in []byte) (out []byte) {
	buffer := make([]byte, 0, 2_000_000)
	peeker := peekTwo(in)
	for {
		x, y, err := peeker()
		if err != nil {
			break
		}
		for i := 0; i < int(y); i++ {
			buffer = append(buffer, x)
		}
	}
	return buffer
}

func peekTwo(in []byte) func() (x, y byte, err error) {
	i := 0

	return func() (x, y byte, err error) {
		if len(in)-1 <= i {
			return 0, 0, errors.New("Completed")
		}

		i += 2
		return in[i-2], in[i-1], nil
	}
}
