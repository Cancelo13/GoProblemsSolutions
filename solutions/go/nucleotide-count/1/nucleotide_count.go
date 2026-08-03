package nucleotidecount

import "errors"

type DNA string
type Histogram map[rune]int

func (Dna DNA) Counts() (Histogram, error) {
	h := Histogram{
        'A': 0,
        'C': 0,
        'G': 0,
        'T': 0,
    }
	for i := 0; i < len(Dna); i++ {
		if Dna[i] == 'A' || Dna[i] == 'C' || Dna[i] == 'G' || Dna[i] == 'T' {
			h[rune(Dna[i])]++
		} else {
			return nil, errors.New("INVALID")
		}
	}
	return h, nil
}
