package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
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
    sizes := make([]int, 0, 10)
    current_size := 0
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            sizes = append(sizes, current_size)
            current_size = 0
        } else {
            number, err := strconv.Atoi(line)
            if err != nil {
                log.Fatal(err)
            }
            current_size += number
        }
    }
    sort.Slice(sizes, func(a, b int) bool {
      return sizes[b] < sizes[a]
    })
    total_count := 0
    for i := 0; i < 3; i++ {
        total_count += sizes[i]
    }
    fmt.Println("Top 3 sizes", total_count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
