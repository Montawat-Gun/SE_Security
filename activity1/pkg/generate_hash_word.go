package activity1

import (
	"strings"
)

func GenerateHashWord(word string, hashWord *map[string]string) {
	combinationWords, hashWords := combinationWord(word, hashWord)
	replaceSimbolWithNumber(strings.ToLower(word), combinationWords, &hashWords)
}

// ref https://codereview.stackexchange.com/questions/3222/finding-all-upper-lower-case-combinations-of-a-word
func combinationWord(word string, hashWord *map[string]string) ([]string, map[string]string) {
	word = strings.ToLower(word)
	combinationsCount := 1 << len(word) // equal 2 ^ len(word)
	words := []string{}
	for i := 0; i < combinationsCount; i++ {
		combinations := strings.Split(word, "")
		for j := 0; j < len(word); j++ {
			if ((i >> j) & 1) == 1 { // (7 >> 0000 0000) & 0000 0001 =
				combinations[j] = strings.ToUpper(string(word[j]))
			}
		}
		newWord := strings.Join(combinations, "")
		words = append(words, newWord)
		(*hashWord)[HashSha1(newWord)] = newWord
	}
	return words, *hashWord
}

func replaceSimbolWithNumber(word string, words []string, hashWord *map[string]string) {
	if len(words) <= 0 {
		return
	}

	indexs := []int{}
	for i := 0; i < len(word); i++ {
		if string(word[i]) == "o" || string(word[i]) == "i" || string(word[i]) == "l" {
			indexs = append(indexs, i)
		}
	}

	combinationsLength := 1 << len(indexs)

	wordsIndex := 0
	for {
		currentWord := words[wordsIndex]

		replace(&words, hashWord, combinationsLength, currentWord, indexs)

		wordsIndex++
		if wordsIndex >= len(words) {
			break
		}
	}
}

func replace(words *[]string, hashWord *map[string]string, allCombinations int, currentWord string, indexs []int) {
	for i := 0; i < allCombinations; i++ {
		combinations := strings.Split(currentWord, "")
		for j, index := range indexs {
			if ((i >> j) & 1) == 1 { // (7 >> 0000 0000) & 0000 0001 =
				if strings.ToLower(combinations[index]) == "o" {
					combinations[index] = "0"
				}
				if strings.ToLower(combinations[index]) == "l" {
					combinations[index] = "1"
				}
				if strings.ToLower(combinations[index]) == "i" {
					combinations[index] = "1"
				}
			}
		}
		newWord := strings.Join(combinations, "")
		hasNewWord := HashSha1(newWord)
		if _, isContain := (*hashWord)[hasNewWord]; !isContain {
			*words = append(*words, newWord)
			(*hashWord)[HashSha1(newWord)] = newWord
		}
	}
}
