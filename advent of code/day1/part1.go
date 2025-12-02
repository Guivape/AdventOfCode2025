package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
	"strings"
)

var currentNumber int = 50
var result int = 0
var fullLap int = 100

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("failed to open file: %s", err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            }
            log.Fatalf("error reading file: %s", err)
        }

        line = strings.TrimSpace(line) 

        letter := string(line[0])            
        number, _ := strconv.Atoi(line[1:])
		
		normalizedNumber := number % fullLap
		if (letter == "L") {
			currentNumber -= normalizedNumber
		} else {
			currentNumber += normalizedNumber
		}

		currentNumber = currentNumber % fullLap

		if currentNumber == 0 {
    		result++
		}	

        fmt.Println(currentNumber)
    }
		fmt.Printf("Pw: %d\n", result)
}
