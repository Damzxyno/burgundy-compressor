package compressors

import (
	"fmt"
)

type ZynoRunLengthEncoding struct{}

func (r ZynoRunLengthEncoding) Encode(in []byte) []byte {

	buffer := make([]byte, 0, 2_000_000)
	fmt.Printf("The size of buffer, %d\n", len(buffer))
	curr := in[0]
	freq := 0

	for _, v := range in {
		if curr == v {
			freq++
		} else {
			buffer = append(buffer, curr)
			if freq > 1 {
				fmt.Println("Ready", freq, freq)
				buffer = append(buffer, byte((freq-1)*-1))
			}
			curr = v
			freq = 1
		}
	}
	buffer = append(buffer, curr)
	if freq > 1 {
		buffer = append(buffer, byte((freq-1)*-1))
	}
	return buffer
}

func (r ZynoRunLengthEncoding) Decode(in []byte) (out []byte) {
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
