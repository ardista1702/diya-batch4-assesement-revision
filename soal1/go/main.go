package main

import (
	"fmt"
	"strings"
)

func getAlphabetIndex(char string) int {
	return int(char[0] - 'a')
}

func calculateTotalSteps(testCase string) int {
	alphabetList := []int{}
	for range 26 {
		alphabetList = append(alphabetList, 0)
	}
	splitetedTestCase := strings.Split(testCase, " ")
	s1 := splitetedTestCase[0]
	s2 := splitetedTestCase[1]

	totalSameChar := 0
	for _, r := range s1 {
		ch := string(r)
		index := getAlphabetIndex(strings.ToLower(ch))
		alphabetList[index]++
	}
	for _, r := range s2 {
		ch := string(r)
		index := getAlphabetIndex(strings.ToLower(ch))

		if alphabetList[index] > 0 {
			totalSameChar++
			alphabetList[index]--
		}
	}
	return len(s1) - totalSameChar + len(s2) - totalSameChar

}

func main() {
	testCases := []string{
		"KURSI BUI",
		"AYAM BAYAM",
		"HITAM PUTIH",
	}
	for _, tc := range testCases {
		fmt.Printf("total step %v\n", calculateTotalSteps(tc))
	}
}