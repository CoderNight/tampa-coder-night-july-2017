package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"io"
	"log"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if !(stat.Mode() & os.ModeCharDevice == 0) {
		fmt.Println("no input")
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)

	sizeA, _ := strconv.Atoi(readLine(reader)[0])
	listA := readLine(reader)

	sizeB, _ := strconv.Atoi(readLine(reader)[0])
	listB := readLine(reader)

	if sizeA > sizeB {
		printMissing(listA, listB)
	} else {
		printMissing(listB, listA)
	}
}

func printMissing(larger []string, smaller []string) {
	missingSize := len(larger) - len(smaller)
	missing := make([]string, missingSize)

	seen := make(map[string]int)
	for _, val := range smaller {
		// keep track of what the smaller list *does* have
		seen[val]++
	}

	i := 0
	for _, val := range larger {
		// decrement what matches
		seen[val]--

		// if we decrement past 0, means we are missing the value
		if seen[val] < 0 {
			if i == 0 {
				missing[i] = val
				i++
			} else if i < missingSize {
				j := i

				// sort our collection of missing values as we add the new value
				for j > 0 && val < missing[j - 1] {
					// new value is less than existing for the position; swap
					existing := missing[j - 1]
					missing[j] = existing
					// decrement to check previous position for sort swapping
					if j > 0 {
						j--
					}
				}

				// only increment our missing values index if we haven't seen this value already
				if missing[j] != val {
					missing[j] = val
					i++
				}
			}
		}
	}

	fmt.Println("out", strings.Join(missing, " "))
}

func readLine(reader *bufio.Reader) []string {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return strings.Fields(line)
}
