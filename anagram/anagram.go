package anagram

import (
	"sort"
	"strings"
)

func stringSort(word string) string {
	word = strings.ToLower(word)
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func Detect(subject string, candidates []string) (result []string) {
	subjectSorted := stringSort(subject)

	for _, candidate := range candidates {
		if len(candidate) != len(subject) {
			continue
		}
		// words are not anagrams of themselves (case-insensitive)
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}
		candidateSorted := stringSort(candidate)
		if subjectSorted == candidateSorted {
			result = append(result, candidate)
		}
	}

	return
}
