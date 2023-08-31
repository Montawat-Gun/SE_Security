package activity1

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

var shaOne = sha1.New()

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
		(*hashWord)[hashSha1(newWord)] = newWord
	}
	return words, *hashWord
}

func replaceSimbolWithNumber(word string, words []string, hashWord *map[string]string) {
	if len(words) <= 0 {
		return
	}

	oIndexs := []int{}
	lIndexs := []int{}
	iIndexs := []int{}
	for i := 0; i < len(word); i++ {
		if string(word[i]) == "o" {
			oIndexs = append(oIndexs, i)
		}
		if string(word[i]) == "l" {
			lIndexs = append(lIndexs, i)
		}
		if string(word[i]) == "i" {
			iIndexs = append(iIndexs, i)
		}
	}

	oCombinations := 1 << len(oIndexs)
	lCombinations := 1 << len(lIndexs)
	iCombinations := 1 << len(iIndexs)

	wordsIndex := 0
	for {
		currentWord := words[wordsIndex]

		replace(&words, hashWord, oCombinations, currentWord, oIndexs, "0")
		replace(&words, hashWord, lCombinations, currentWord, lIndexs, "1")
		replace(&words, hashWord, iCombinations, currentWord, iIndexs, "1")

		// for i := 1; i < oCombinations; i++ {
		// 	combinations := strings.Split(currentWord, "")
		// 	for j, oIndex := range oIndexs {
		// 		if ((i >> j) & 1) == 1 { // (7 >> 0000 0000) & 0000 0001 =
		// 			combinations[oIndex] = "0"
		// 		}
		// 	}
		// 	newWord := strings.Join(combinations, "")
		// 	if (*hashWord)[hasNewWord] == nil {
		// 		words = append(words, newWord)
		// 		(*hashWord)[hashSha1(newWord)] = newWord
		// 	}
		// }
		// for i := 1; i < lCombinations; i++ {
		// 	combinations := strings.Split(currentWord, "")
		// 	for j, lIndex := range lIndexs {
		// 		if ((i >> j) & 1) == 1 { // (7 >> 0000 0000) & 0000 0001 =
		// 			combinations[lIndex] = "1"
		// 		}
		// 	}
		// 	newWord := strings.Join(combinations, "")
		// 	if (*hashWord)[hasNewWord] == nil {
		// 		words = append(words, newWord)
		// 		(*hashWord)[hashSha1(newWord)] = newWord
		// 	}
		// }
		// for i := 1; i < iCombinations; i++ {
		// 	combinations := strings.Split(currentWord, "")
		// 	for j, iIndex := range iIndexs {
		// 		if ((i >> j) & 1) == 1 { // (7 >> 0000 0000) & 0000 0001 =
		// 			combinations[iIndex] = "1"
		// 		}
		// 	}
		// 	newWord := strings.Join(combinations, "")
		// 	hasNewWord := hashSha1(newWord)
		// 	if (*hashWord)[hasNewWord] == nil {
		// 		words = append(words, newWord)
		// 		(*hashWord)[hashSha1(newWord)] = newWord
		// 	}
		// }

		wordsIndex++
		if wordsIndex >= len(words) {
			break
		}
	}
}

func replace(words *[]string, hashWord *map[string]string, allCombinations int, currentWord string, indexs []int, replacement string) {
	for i := 1; i < allCombinations; i++ {
		combinations := strings.Split(currentWord, "")
		for j, iIndex := range indexs {
			if ((i >> j) & 1) == 1 { // (7 >> 0000 0000) & 0000 0001 =
				combinations[iIndex] = replacement
			}
		}
		newWord := strings.Join(combinations, "")
		hasNewWord := hashSha1(newWord)
		if _, ok := (*hashWord)[hasNewWord]; ok {
			*words = append(*words, newWord)
			(*hashWord)[hashSha1(newWord)] = newWord
		}
	}
}

func hashSha1(word string) string {
	shaOne.Reset()
	shaOne.Write([]byte(word))
	return hex.EncodeToString(shaOne.Sum(nil))
}
