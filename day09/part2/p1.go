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

func adjustTailPos(head_pos *point, tail_pos *point, visited map[point]bool) {
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
    tail_pos.x = new_pos.x;
    tail_pos.y = new_pos.y;
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
    adjustTailPos(head_pos, tail_pos, visited)
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
    tail_count := 9
    tails := make([]point, 0, 10)
    for i := 0; i < tail_count; i++ {
        tails = append(tails, point{x: 0, y: 0})
    }
    visited[tails[tail_count - 1]] = true
    for scanner.Scan() {
        line := scanner.Text()
        s := strings.Split(line, " ")
        amount, err := strconv.Atoi(s[1])
        if err != nil {
            log.Fatal(err)
        }
        for i := 0; i < amount; i++ {
            moveRope(s[0][0], &head_pos, &tails[0], visited)
            for i := 1; i < tail_count; i++ {
                adjustTailPos(&tails[i - 1], &tails[i], visited)
            }
            visited[tails[tail_count - 1]] = true
        }
    }
    fmt.Println(len(visited))


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
