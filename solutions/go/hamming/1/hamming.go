package hamming

import(
    "errors"
)

func Distance(strandA string, strandB string) (int, error) {
	if len(strandA) != len(strandB) {
		return 0, errors.New("Length doesn't match")
	}
	var counter int = 0
	for i := 0; i < len(strandA); i++ {
		if strandA[i] != strandB[i] {
			counter++
		}
	}
	return counter, nil
}
