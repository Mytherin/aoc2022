package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

// Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles.
func getCycleScore(cycle int, value int) int {
    if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
        return cycle * value        
    }
    return 0
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    cycle := 0
    value := 1
    score := 0
    for scanner.Scan() {
        line := scanner.Text()
        if line == "noop" {
            cycle++
            score += getCycleScore(cycle, value)
            continue
        }
        s := strings.Split(line, " ")
        if s[0] != "addx" {
            panic("eek")
        }
        amount, err := strconv.Atoi(s[1])
        if err != nil {
            log.Fatal(err)
        }
        cycle++
        score += getCycleScore(cycle, value)
        cycle++
        score += getCycleScore(cycle, value)
        value += amount
    }
    fmt.Println(score)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
