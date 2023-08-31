package activity1

import (
	"strings"
)

func generateCombinations(input string, currentCombination string, index int, result *[]string) {
	if index == len(input) {
		*result = append(*result, currentCombination)
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

func GenerateHash(word string) []string {
	var result []string
	generateCombinations(word, "", 0, &result)
	return result
}
