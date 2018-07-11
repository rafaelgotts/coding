package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)

    if isWeird(N) {
        fmt.Println("Weird")
    } else {
        fmt.Println("Not Weird")
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func isOdd(number int32) bool {
    mod := number % 2
    if mod == 0 {
        return false
    }

    return true
}

func isWeird(number int32) bool {
    if isOdd(number) || ( number >= 6 && number <= 20 ) {
        return true
    }

    return false
}