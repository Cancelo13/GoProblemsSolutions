package proteintranslation

import "errors"

var ErrStop = errors.New("stop codon")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	var res []string = []string{}
    var valid bool = false
	for i := 0; i+2 < len(rna); i += 3 {
		var c string = rna[i : i+3]
		k, err := FromCodon(c)
		if err != nil {
			if err == ErrStop {
                valid = true
				break
			}
			return nil, err
		}
		res = append(res, k)
	}
	if len(rna)%3 != 0 && !valid {
		return nil, ErrInvalidBase
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}