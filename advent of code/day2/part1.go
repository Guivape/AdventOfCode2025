package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), ",")

	result := 0
	for _, r := range parts {
		bounds := strings.Split(r, "-")
		firstNumber, _ := strconv.Atoi(bounds[0])
		secondNumber, _ := strconv.Atoi(bounds[1])

		for i := firstNumber; i <= secondNumber; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 != 0 {
				continue
			}

			half := len(s) / 2
			if s[:half] == s[half:] {
				fmt.Println(i)
				result += i
			}
		}
	}

	fmt.Println(result)
}