package strand

func ToRNA(dna string) string {
	var rna string
	for _, ch := range dna {
		switch ch {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
