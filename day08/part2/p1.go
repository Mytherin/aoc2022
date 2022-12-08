package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func compute_scenic_score(trees [][]int, x int, y int, width int, height int) int {
    tree_height := trees[y][x]
    // look to the left
    left_score := 0
    right_score := 0
    up_score := 0
    down_score := 0
    for i := x - 1; i >= 0; i-- {
        left_score++
        if trees[y][i] >= tree_height {
            break
        }
    }

    // look to the right
    for i := x + 1; i < width; i++ {
        right_score++
        if trees[y][i] >= tree_height {
            break
        }
    }

    // look up
    for i := y - 1; i >= 0; i-- {
        up_score++
        if trees[i][x] >= tree_height {
            break
        }
    }

    // look down
    for i := y + 1; i < height; i++ {
        down_score++
        if trees[i][x] >= tree_height {
            break
        }
    }
    return left_score * right_score * up_score * down_score
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    trees := make([][]int, 0, 10)
    for scanner.Scan() {
        line := scanner.Text()

        tree_line := make([]int, 0, 10)
        for i := 0; i < len(line); i++ {
            number := int(line[i]) - int('0')
            tree_line = append(tree_line, number)
        }
        trees = append(trees, tree_line)
    }
    width := len(trees[0])
    height := len(trees)
    max_score := 0
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            score := compute_scenic_score(trees, x, y, width, height) 
            if score > max_score {
                max_score = score
            }
        }
    }
    fmt.Println(max_score)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
