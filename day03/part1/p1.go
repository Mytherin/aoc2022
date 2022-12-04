package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.


func get_double_item(line string) byte {
    total_length := len(line)
    bag_length := total_length / 2

    for i := 0; i < bag_length; i++ {
        current_rune := line[i]
        for j := 0; j < bag_length; j++ {
            if current_rune == line[bag_length + j] {
                return current_rune
            }
        }
    }
    panic("Did not find double item")
}

func get_item_score(r byte) int {
    ascii := int(r)
    if ascii >= 97 && ascii <= 122 {
        // a-z
        return ascii - 97 + 1;
    } else if ascii >= 65 && ascii <= 90  {
        // A-Z
        return ascii - 65 + 27
    }
    panic("Unrecognized input")
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    score := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        char := get_double_item(line)
        score += get_item_score(char)
    }
    fmt.Println("Score: %d\n", score)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
