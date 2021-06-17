package strand

var m = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U'}

func ToRNA(dna string) string {
	var rna string
	for _, nucleotide := range dna {
		rna += string(m[nucleotide])
	}
	return rna
}
