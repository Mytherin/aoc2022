package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func getPixelOutput(cycle int, value int) string {
    screen_pos := cycle % 40
    if screen_pos >= value - 1 && screen_pos <= value + 1 {
        return "#"
    }
    return "."
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    output := ""
    cycle := 0
    value := 1
    for scanner.Scan() {
        line := scanner.Text()
        if line == "noop" {
            cycle++
            output += getPixelOutput(cycle, value)
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
        output += getPixelOutput(cycle, value)
        cycle++
        value += amount
        output += getPixelOutput(cycle, value)
    }
    width := 40
    for i := 0; i < len(output); i += width {
        fmt.Println(output[i:i+width])
    }


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
