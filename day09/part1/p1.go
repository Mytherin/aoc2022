package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type point struct {
    x int
    y int
}

func setTailPos(tail_pos *point, new_pos point, visited map[point]bool) {
    tail_pos.x = new_pos.x;
    tail_pos.y = new_pos.y;
    visited[new_pos] = true
}

func abs(x int)int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}

func sign(x int)int {
    if x < 0 {
        return -1
    } else if x > 0 {
        return 1
    } else {
        return 0
    }
}

func moveRope(ch byte, head_pos *point, tail_pos *point, visited map[point]bool) {
    if ch == 'U' {
        head_pos.y -= 1;
    } else if ch == 'D' {
        head_pos.y += 1;
    } else if ch == 'R' {
        head_pos.x += 1;
    } else if ch == 'L' {
        head_pos.x -= 1;
    } else {
        panic("eek")
    }
    xdiff := head_pos.x - tail_pos.x
    ydiff := head_pos.y - tail_pos.y
    xdist := abs(xdiff)
    ydist := abs(ydiff)
    if xdist <= 1 && ydist <= 1 {
        return
    }
    new_pos := point { tail_pos.x, tail_pos.y }
    new_pos.x += sign(xdiff)
    new_pos.y += sign(ydiff)
    setTailPos(tail_pos, new_pos, visited)
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // L D R U
    visited := make(map[point]bool)
    head_pos := point{x: 0, y: 0}
    tail_pos := point{x: 0, y: 0}
    setTailPos(&tail_pos, point{x: 0, y: 0}, visited)
    for scanner.Scan() {
        line := scanner.Text()
        s := strings.Split(line, " ")
        amount, err := strconv.Atoi(s[1])
        if err != nil {
            log.Fatal(err)
        }
        for i := 0; i < amount; i++ {
            moveRope(s[0][0], &head_pos, &tail_pos, visited)
        }
    }
    fmt.Println(len(visited))


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
