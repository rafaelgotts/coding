package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    var inputString []string
    scanner := bufio.NewScanner(os.Stdin)

    linesToRead := GetLines(scanner)

    for i := 0; i < linesToRead ; i++ {
        inputString = append(inputString, GetInputLine(scanner))
    }

    for i := 0; i < len(inputString) ; i++ {
        fmt.Printf(
            "%s %s\n",
            getEvenChars(inputString[i]),
            getOddChars(inputString[i]),
        )
    }
}

func GetLines (scanner *bufio.Scanner) (lines int) {
    scanner.Scan()
    input, _ := strconv.Atoi(scanner.Text())

    return input
}

func GetInputLine(scanner *bufio.Scanner) (inputLine string) {
    scanner.Scan()

    return scanner.Text()
}

func getEvenChars(input string) (evenChars string) {
    input = input[:]
    for i := 0; i < len(input) ; i += 2 {
        evenChars += input[i:i+1]
    }

    return evenChars
}

func getOddChars(input string) (oddChars string) {
    input = input[:]
    for i := 1; i < len(input) ; i += 2 {
        oddChars += input[i:i+1]
    }

    return oddChars
}
