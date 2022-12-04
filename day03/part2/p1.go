package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.


func get_badge(groups []string) byte {
    first_line := groups[0]
    for c := 0; c < len(first_line); c++ {
        current_rune := first_line[c]
        found_count := 1
        for grp_idx := 1; grp_idx < len(groups); grp_idx++ {
            line := groups[grp_idx]
            for i := 0; i < len(line); i++ {
                if line[i] == current_rune {
                    found_count++
                    break
                }
            }
        }
        if found_count == len(groups) {
            return current_rune
        }
    }
    panic("Did not find badge")
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
    groups := make([]string, 0, 10)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        groups = append(groups, line)
        if len(groups) == 3 {
            char := get_badge(groups)
            score += get_item_score(char)
            groups = make([]string, 0, 10)
        }
    }
    fmt.Println("Score: %d\n", score)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
