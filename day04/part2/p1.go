package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func do_ranges_overlap(r1_lower int, r1_upper int, r2_lower int, r2_upper int) bool {
    if r1_upper < r2_lower {
        return false
    }
    if r2_upper < r1_lower {
        return false
    }
    return true
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    overlapping_ranges := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        s := strings.Split(line, ",")
        range_one := strings.Split(s[0], "-")
        range_two := strings.Split(s[1], "-")
        r1_lower, err := strconv.Atoi(range_one[0])
        if err != nil {
            log.Fatal(err)
        }
        r1_upper, err := strconv.Atoi(range_one[1])
        if err != nil {
            log.Fatal(err)
        }
        r2_lower, err := strconv.Atoi(range_two[0])
        if err != nil {
            log.Fatal(err)
        }
        r2_upper, err := strconv.Atoi(range_two[1])
        if err != nil {
            log.Fatal(err)
        }
        if do_ranges_overlap(r1_lower, r1_upper, r2_lower, r2_upper) {
            overlapping_ranges += 1
        }
    }
    fmt.Println("Overlapping Ranges: %d\n", overlapping_ranges)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
