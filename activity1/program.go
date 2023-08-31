package main

import (
	act1 "activity1/pkg"
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

func getWords() []string {
	resp, err := http.Get("https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10k-most-common.txt")
	checkErr(err)
	body, err := io.ReadAll(resp.Body)
	checkErr(err)
	return strings.Split(string(body), "\n")
}

func generateHash(words []string, result *map[string]string) {
	for _, word := range words {
		act1.GenerateHashWithRecursive(word, result)
	}
}

func main() {
	words := getWords()
	result := make(map[string]string)

	startTime := time.Now()
	fmt.Printf("start at %s \n", startTime.String())

	generateHash(words, &result)

	fmt.Printf("result : %v \n", len(result))
	fmt.Printf("end gen hash at %s \n", time.Now().String())

	records := [][]string{}
	for k, v := range result {
		records = append(records, []string{k, v})
	}

	f, err := os.Create("rainbow_table.csv")
	w := csv.NewWriter(f)
	w.WriteAll(records)
	checkErr(err)
	endTime := time.Now()
	fmt.Printf("complete at %s \n", endTime.String())
	fmt.Printf("time %s \n", endTime.Sub(startTime))
	fmt.Printf("result count %v \n", len(records))
}
