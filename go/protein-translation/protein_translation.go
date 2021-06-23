package protein

import (
	"errors"
)

var ErrStop error = errors.New("STOP")
var ErrInvalidBase error = errors.New("Invalid codon")

var proteins = map[string][]string{
	"Methionine":    {"AUG"},
	"Phenylalanine": {"UUU", "UUC"},
	"Leucine":       {"UUA", "UUG"},
	"Serine":        {"UCU", "UCC", "UCA", "UCG"},
	"Tyrosine":      {"UAU", "UAC"},
	"Cysteine":      {"UGU", "UGC"},
	"Tryptophan":    {"UGG"},
	"STOP":          {"UAA", "UAG", "UGA"}}

func FromRNA(rna string) (out []string, e error) {
	e = nil
	for i := 0; i < len(rna)-2; i += 3 {
		poly, err := FromCodon(rna[i : i+3])
		if poly != "" {
			out = append(out, poly)
		}
		if err == ErrInvalidBase {
			e = err
		}
		if err == ErrStop {
			return
		}
	}
	return
}

func FromCodon(codon string) (string, error) {
	for p := range proteins {
		if contains(codon, proteins[p]) {
			if p == "STOP" {
				return "", ErrStop
			}
			return p, nil
		}
	}
	return "", ErrInvalidBase
}

func contains(codon string, sequence []string) bool {
	for _, v := range sequence {
		if codon == v {
			return true
		}
	}
	return false
}
