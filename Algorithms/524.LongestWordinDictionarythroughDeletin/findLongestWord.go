package longestwordindictionarythroughdeletin

import "sort"

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		a, b := dictionary[i], dictionary[j]
		return len(a) > len(b) || len(a) == len(b) && a < b
	})
	for _, word := range dictionary {
		word_index := 0
		for s_index := range s {
			if s[s_index] == word[word_index] {
				word_index++
				if word_index == len(word) {
					return word
				}
			}
		}
	}
	return ""
}
