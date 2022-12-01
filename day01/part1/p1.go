package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    current_size := 0
    max_size := 0
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            if current_size > max_size {
                max_size = current_size
            }
            current_size = 0
        } else {
            number, err := strconv.Atoi(line)
            if err != nil {
                log.Fatal(err)
            }
            current_size += number
        }
    }
    fmt.Println("Max size: %d\n", max_size)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
