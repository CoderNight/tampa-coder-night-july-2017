package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

//GetNumbers returns a map of numbers and their frequency based on a string of space delimitted numbers
func GetNumbers(s string) map[int]int {
	parsed := ParseNumbers(s)
	return CountNumbers(parsed)
}

//ParseNumbers returns an array of ints from a list of strings of numbers
func ParseNumbers(s string) []int {
	numStrings := strings.Split(s, " ")
	result := make([]int, len(numStrings))
	for i := range numStrings {
		converted, err := strconv.Atoi(numStrings[i])
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		result[i] = converted
	}
	return result
}

//CountNumbers returns frequencies of numbers in an array of numbers
func CountNumbers(numbers []int) map[int]int {
	result := make(map[int]int, 0)
	for _, v := range numbers {
		result[v]++
	}
	return result
}

//SubtractFrequency removes frequency from a map of number frequencies
func SubtractFrequency(foundNumbers []int, numFrequencies map[int]int) map[int]int {
	for _, v := range foundNumbers {
		numFrequencies[v]--
	}
	return numFrequencies
}

//OutputNumberFrequencyInOrder returns any positive frequencies in a map of integers
func OutputNumberFrequencyInOrder(frequencies map[int]int) []int {
	array := make([]int, 0)
	for k, f := range frequencies {
		for f > 0 {
			array = append(array, k)
			f--
		}
	}
	return array
}

//OrderMissingNumbers tests numbers found in an array against expected frequencies,
// outputing any missing numbers in order
func OrderMissingNumbers(foundNumbers []int, possibleNumbers map[int]int) []int {
	diff := SubtractFrequency(foundNumbers, possibleNumbers)
	output := OutputNumberFrequencyInOrder(diff)
	sort.Ints(output)
	return output
}

func main() {
	f, err := os.Open(os.Args[1])
	handleErr(err)
	s := bufio.NewScanner(f)
	handleErr(err)
	lines := []string{}
	for s.Scan() {
		// note! this will break on max lines. Since each number could be 5 chars long (100000)
		// with up to 200,000 numbers in a list, one line could be over a meg long
		lines = append(lines, s.Text())
	}
	A := ParseNumbers(lines[1])
	B := GetNumbers(lines[3])
	printOut(OrderMissingNumbers(A, B))
}

func printOut(numbers []int) {
	for _, v := range numbers {
		print(v, " ")
	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
