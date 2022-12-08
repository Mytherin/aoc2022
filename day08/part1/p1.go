package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func tree_is_visible(trees [][]int, x int, y int, width int, height int) bool {
    tree_height := trees[y][x]
    // look to the left
    found_higher := false
    for i := x - 1; i >= 0; i-- {
        if trees[y][i] >= tree_height {
            found_higher = true
            break
        }
    }
    if !found_higher {
        return true
    }

    // look to the right
    found_higher = false
    for i := x + 1; i < width; i++ {
        if trees[y][i] >= tree_height {
            found_higher = true
            break
        }
    }
    if !found_higher {
        return true
    }

    // look up
    found_higher = false
    for i := y - 1; i >= 0; i-- {
        if trees[i][x] >= tree_height {
            found_higher = true
            break
        }
    }
    if !found_higher {
        return true
    }

    // look down
    found_higher = false
    for i := y + 1; i < height; i++ {
        if trees[i][x] >= tree_height {
            found_higher = true
            break
        }
    }
    if !found_higher {
        return true
    }
    return false
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
    visible_trees := 0
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            if tree_is_visible(trees, x, y, width, height) {
                visible_trees++
            }
        }
    }
    fmt.Println(visible_trees)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
