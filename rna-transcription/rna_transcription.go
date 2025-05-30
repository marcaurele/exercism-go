package strand

func ToRNA(dna string) string {
	var result string
	for _, letter := range dna {
		switch letter {
		case 'G':
			result += "C"
		case 'C':
			result += "G"
		case 'T':
			result += "A"
		case 'A':
			result += "U"
		}
	}
	return result
}
