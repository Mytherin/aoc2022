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

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for i := 4; i < len(line); i++ {
            start_pos := i - 4
            found_equal := false
            for k := 0; k < 4; k++ {
                for l := 0; l < 4; l++ {
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
