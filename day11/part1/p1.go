package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

// Monkey 3:
//   Starting items: 81, 68, 65, 81, 73, 77, 96
//   Operation: new = old + 2
//   Test: divisible by 17
//     If true: throw to monkey 4
//     If false: throw to monkey 6
type Operation int

const (
    Add Operation = 0
    Multiply      = 1
)

type Monkey struct {
    items []int
    left int
    op Operation
    right int
    divisible_by int
    if_true int
    if_false int
    inspections int
}

func parse_value(input string)int {
    if input == "old" {
        return -1
    } else {
        val, err := strconv.Atoi(input)
        if err != nil {
            panic(err)
        }
        return val
    }
}

func perform_round(monkeys []Monkey, monkey_index int) {
    monkey := monkeys[monkey_index]
    current_items := monkey.items
    monkeys[monkey_index].items = make([]int, 0, 10)
    for i := 0; i < len(current_items); i++ {
        monkeys[monkey_index].inspections += 1
        worry_level := current_items[i]
        left := monkey.left
        right := monkey.right
        if monkey.left < 0 {
            left = worry_level
        }
        if monkey.right < 0 {
            right = worry_level
        }
        if monkey.op == Add {
            worry_level = left + right
        } else if monkey.op == Multiply {
            worry_level = left * right
        } else {
            panic("Unknown operation")
        }
        worry_level /= 3
        target := monkey.if_false
        if worry_level % monkey.divisible_by == 0 {
            target = monkey.if_true
        }
        monkeys[target].items = append(monkeys[target].items, worry_level)
    }
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    monkeys := make([]Monkey, 0, 10)
    for scanner.Scan() {
        var result Monkey
        result.items = make([]int, 0, 10)
        result.inspections = 0
        line := scanner.Text()
        s := strings.Split(line, " ")
        if s[0] != "Monkey" {
            panic("expected monkey!")
        }
        monkey_index, err := strconv.Atoi(strings.Replace(s[1], ":", "", 1))
        if err != nil {
            panic(err)
        }
        if monkey_index != len(monkeys) {
            panic("wrong index!")
        }
        // parse items
        if !scanner.Scan() {
            panic("unexpected eof")
        }
        line = scanner.Text()
        s = strings.Split(line, ": ")
        if s[0] != "  Starting items" {
            panic("expected starting items! found " +  s[0])
        }
        s = strings.Split(s[1], ", ")
        for i := 0; i < len(s); i++ {
            n, err := strconv.Atoi(s[i])
            if err != nil {
                panic(err)
            }
            result.items = append(result.items, n)
        }
        // parse operation
        if !scanner.Scan() {
            panic("unexpected eof")
        }
        line = scanner.Text()
        s = strings.Split(line, ": ")
        if s[0] != "  Operation" {
            panic("expected operation!")
        }
        s = strings.Split(s[1], " = ")
        if s[0] != "new" {
            panic("expected new!")
        }
        string_split := ""
        if strings.Contains(s[1], "+") {
            string_split = " + "
            result.op = Add
        } else if strings.Contains(s[1], "*") {
            string_split = " * "
            result.op = Multiply
        } else {
            panic("unsupported operation")
        }
        s = strings.Split(s[1], string_split)
        result.left = parse_value(s[0])
        result.right = parse_value(s[1])
        // parse divisible
        if !scanner.Scan() {
            panic("unexpected eof")
        }
        line = scanner.Text()
        s = strings.Split(line, " ")
        n, err := strconv.Atoi(s[len(s) - 1])
        if err != nil {
            panic(err)
        }
        result.divisible_by = n
        // parse if true
        if !scanner.Scan() {
            panic("unexpected eof")
        }
        line = scanner.Text()
        s = strings.Split(line, " ")
        n, err = strconv.Atoi(s[len(s) - 1])
        if err != nil {
            panic(err)
        }
        result.if_true = n
        // parse if false
        if !scanner.Scan() {
            panic("unexpected eof")
        }
        line = scanner.Text()
        s = strings.Split(line, " ")
        n, err = strconv.Atoi(s[len(s) - 1])
        if err != nil {
            panic(err)
        }
        result.if_false = n

        // parse empty line
        finished := true
        if scanner.Scan() {
            finished = false
            line = scanner.Text()
            if line != "" {
                panic("expected end of line")
            }
        }
        monkeys = append(monkeys, result)
        if finished {
            break
        }
    }
    total_rounds := 20
    for round := 0; round < total_rounds; round++ {
        fmt.Println(round)
        for m := 0; m < len(monkeys); m++ {
            perform_round(monkeys, m)
        }
    }
    for m := 0; m < len(monkeys); m++ {
        fmt.Println(m, monkeys[m].inspections)
    }



    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
