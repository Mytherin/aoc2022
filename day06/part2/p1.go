package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    message_size := 14
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for i := message_size; i < len(line); i++ {
            start_pos := i - message_size
            found_equal := false
            for k := 0; k < message_size; k++ {
                for l := 0; l < message_size; l++ {
                    if k == l {
                        continue
                    }
                    if line[start_pos + k] == line[start_pos + l] {
                        found_equal = true
                        break
                    }
                }
            }
            if !found_equal {
                fmt.Println("Result: ", i)
                break
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
