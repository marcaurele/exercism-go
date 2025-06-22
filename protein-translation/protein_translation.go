package protein

import (
	"errors"
)

var ErrStop = errors.New("end of sequence")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	var output []string

	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])

		if err == ErrStop {
			return output, nil
		} else if err != nil {
			return nil, err
		}
		output = append(output, protein)
	}

	return output, nil
}

func FromCodon(codon string) (string, error) {
	var protein string

	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

	return protein, nil
}
