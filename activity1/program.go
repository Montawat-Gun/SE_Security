package main

import (
	activity1 "activity1/pkg"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("start at %s \n", time.Now().String())
	resp, errRes := http.Get("https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10k-most-common.txt")
	checkErr(errRes)
	body, err := io.ReadAll(resp.Body)
	checkErr(err)
	words := strings.Split(string(body), "\n")

	result1 := make(map[string]string)
	result2 := make(map[string]string)
	result3 := make(map[string]string)
	result4 := make(map[string]string)
	result5 := make(map[string]string)

	generateHash(words[:2000], &result1)
	generateHash(words[2000:4000], &result2)
	generateHash(words[4000:6000], &result3)
	generateHash(words[6000:8000], &result4)
	generateHash(words[8000:], &result5)

	fmt.Printf("end gen hash at %s \n", time.Now().String())

	records := [][]string{}
	for k, v := range result1 {
		records = append(records, []string{k, v})
	}
	for k, v := range result2 {
		records = append(records, []string{k, v})
	}
	for k, v := range result3 {
		records = append(records, []string{k, v})
	}
	for k, v := range result4 {
		records = append(records, []string{k, v})
	}
	for k, v := range result5 {
		records = append(records, []string{k, v})
	}
	f, err := os.Create("tmp.csv")
	w := csv.NewWriter(f)
	w.WriteAll(records)
	checkErr(err)
	fmt.Printf("complete at %s \n", time.Now().String())

	tests := [][]string{}
	count := 0
	for _, word := range words {
		for _, r := range activity1.GenerateHash(word) {
			count++
			tests = append(tests, []string{fmt.Sprintf("%v", count), r})
		}
	}
	f, _ = os.Create("tmp1.csv")
	w = csv.NewWriter(f)
	w.WriteAll(tests)
}

func generateHash(words []string, result *map[string]string) {
	for _, word := range words {
		activity1.GenerateHashWord(word, result)
	}
}
