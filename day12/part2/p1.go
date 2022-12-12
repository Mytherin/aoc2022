package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type GridNode struct {
    height int
    score int
}

type Node struct {
    x int
    y int
}

func get_grid_node(grid [][]GridNode, x int, y int) GridNode {
    return grid[y][x]
}

func should_add_node(c GridNode, n GridNode) bool {
    return n.score > c.score + 1 && n.height <= c.height + 1
}

func get_neighbors(grid [][]GridNode, x int, y int) []Node {
    result := make([]Node, 0, 10)
    c := get_grid_node(grid, x, y)

    if x > 0 {
        n := get_grid_node(grid, x - 1, y)
        if should_add_node(c, n) {
            result = append(result, Node { x: x - 1, y : y })
        }
    }
    if x + 1 < len(grid[0]) {
        n := get_grid_node(grid, x + 1, y)
        if should_add_node(c, n) {
            result = append(result, Node { x: x + 1, y : y })
        }
    }
    if y > 0 {
        n := get_grid_node(grid, x, y - 1)
        if should_add_node(c, n) {
            result = append(result, Node { x: x, y : y - 1 })
        }
    }
    if y + 1 < len(grid) {
        n := get_grid_node(grid, x, y + 1)
        if should_add_node(c, n) {
            result = append(result, Node { x: x, y : y + 1 })
        }
    }
    return result
}

func find_path(grid [][]GridNode, nodes []Node, end_x int, end_y int) int {
    for i := 0; i < len(nodes); i++ {
        x := nodes[i].x
        y := nodes[i].y
        score := get_grid_node(grid, x, y).score

        neighbors := get_neighbors(grid, x, y)
        for n_idx := 0; n_idx < len(neighbors); n_idx++ {
            n := neighbors[n_idx]

            grid[n.y][n.x].score = score + 1
            nodes = append(nodes, n)
        }
    }
    if grid[end_y][end_x].score == 99999 {
        return 99999
    }
    return grid[end_y][end_x].score
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid := make([][]GridNode, 0, 10)

    start_nodes := make([]Node, 0, 10)
    var end_x int
    var end_y int
    for scanner.Scan() {
        line := scanner.Text()
        result := make([]GridNode, 0, 10)
        for x := 0; x < len(line); x++ {
            var character byte
            if line[x] == 'S' {
                character = 'a'
            } else if line[x] == 'E' {
                end_x = x
                end_y = len(grid)
                character = 'z'
            } else {
                character = line[x]
            }
            if character == 'a' {
                var start Node
                start.x = x
                start.y = len(grid)   
                start_nodes = append(start_nodes, start)             
            }
            if character >= 'a' && character <= 'z' {
                var new_node GridNode
                new_node.height = int(character - 'a')
                new_node.score = 99999
                result = append(result, new_node)
            } else {
                panic("unrecognized letter")
            }
        }
        grid = append(grid, result)
    }
    shortest_path := 99999
    for s := 0; s < len(start_nodes); s++ {
        start := start_nodes[s]
        for y := 0; y < len(grid); y++ {
            for x := 0; x < len(grid[y]); x++ {
                grid[y][x].score = 99999
            }
        }
        grid[start.y][start.x].score = 0
        nodes := make([]Node, 0, 10)
        nodes = append(nodes, start)
        score := find_path(grid, nodes, end_x, end_y)
        if score < shortest_path {
            shortest_path = score
        }

    }
    fmt.Println(shortest_path)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
