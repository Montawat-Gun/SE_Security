package activity1

import (
	"strings"
)

func generateCombinations(input string, currentCombination string, index int, result *map[string]string) {
	if index == len(input) {
		if _, hasValue := (*result)[currentCombination]; !hasValue {
			currentHashWord := HashSha1(currentCombination)
			(*result)[currentCombination] = currentHashWord
		}
		return
	}

	// Include lowercase letter
	generateCombinations(input, currentCombination+string(input[index]), index+1, result)

	// Include uppercase letter
	generateCombinations(input, currentCombination+strings.ToUpper(string(input[index])), index+1, result)

	// Replace 'i' with '1'
	if input[index] == 'i' {
		generateCombinations(input, currentCombination+"1", index+1, result)
	}

	// Replace 'l' with '1'
	if input[index] == 'l' {
		generateCombinations(input, currentCombination+"1", index+1, result)
	}

	// Replace 'o' with '0'
	if input[index] == 'o' {
		generateCombinations(input, currentCombination+"0", index+1, result)
	}
}

func GenerateHashWithRecursive(word string, result *map[string]string) {
	generateCombinations(word, "", 0, result)
}
