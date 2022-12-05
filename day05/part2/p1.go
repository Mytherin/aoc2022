package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

// Stack is a type that represents a stack data structure.
type Stack []byte

// Push adds an element to the top of the stack.
func (s *Stack) Push(v byte) {
    *s = append(*s, v)
}

// Pop removes and returns the top element of the stack.
// If the stack is empty, it returns 0.
func (s *Stack) Pop() byte {
    if len(*s) == 0 {
        panic("Pop from empty stack")
    }

    v := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return v
}

func print_stacks(stacks []Stack) {
    for i := 0; i < len(stacks); i++ {
        text := fmt.Sprintf("%d: ", i)
        for j := 0; j < len(stacks[i]); j++ {
            text += fmt.Sprintf("[%c]", stacks[i][j])
        }
        fmt.Println(text)
    }
}

// Revert the elements in the array
func revert(arr Stack) Stack {
  // Create a new slice with the same length as the original array
  result := make(Stack, len(arr))

  // Iterate through the array in reverse order
  for i := len(arr) - 1; i >= 0; i-- {
    // Add the element at the current index to the result array
    result = append(result, arr[i])
  }

  return result
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    stack_count := 9
    stacks := make([]Stack, 0, stack_count)
    for i := 0; i < stack_count; i++ {
        new_stack := make(Stack, 0, 100)
        stacks = append(stacks, new_stack)
    }
    parse_initial_state := true

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            for i := 0; i < stack_count; i++ {
                stacks[i] = revert(stacks[i])
            }
            parse_initial_state = false
            continue
        }
        if parse_initial_state {
            for i := 0; i < stack_count; i++ {
                character := line[1 + i * 4]
                if character != ' ' && (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') {
                    // we have a crate here
                    stacks[i].Push(character)
                }
            }
        } else {
            // move 2 from 7 to 2
            line = strings.Replace(line, "move ", "", 1)
            line = strings.Replace(line, "from ", "", 1)
            line = strings.Replace(line, "to ", "", 1)
            s := strings.Split(line, " ")
            block_count, err := strconv.Atoi(s[0])
            if err != nil {
                log.Fatal(err)
            }
            source_stack, err := strconv.Atoi(s[1])
            if err != nil {
                log.Fatal(err)
            }
            target_stack, err := strconv.Atoi(s[2])
            if err != nil {
                log.Fatal(err)
            }
            moves := make([]byte, 0, block_count)
            for i := 0; i < block_count; i++ {
                if len(stacks[source_stack - 1]) == 0 {
                    break
                }
                rune := stacks[source_stack - 1].Pop()
                moves = append(moves, rune)
            }
            for i := len(moves) - 1; i >= 0; i-- {
                rune := moves[i]
                stacks[target_stack - 1].Push(rune)
            }
        }
    }
    result := ""
    for i := 0; i < stack_count; i++ {
        rune := stacks[i].Pop()
        result = fmt.Sprintf("%s%c", result, rune)
    }
    fmt.Println("Result: ", result)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
