package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// X for Rock, Y for Paper, and Z for Scissors
// The winner of the whole tournament is the player with the highest score.
// Your total score is the sum of your scores for each round.
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

type Move int

const (
    Rock Move    = 1
    Paper        = 2
    Scissors     = 3
)

func compute_score(us Move, them Move) int {
    switch us {
    case Rock: {
        switch them {
        case Rock:
            return 3 + 1 // draw
        case Paper:
            return 0 + 1 // loss
        case Scissors:
            return 6 + 1 // win
        }
    }
    case Paper: {
        switch them {
        case Rock:
            return 6 + 2 // win
        case Paper:
            return 3 + 2 // draw
        case Scissors:
            return 0 + 2 // loss
        }
    }
    case Scissors: {
        switch them {
        case Rock:
            return 0 + 3 // loss
        case Paper:
            return 6 + 3 // win
        case Scissors:
            return 3 + 3 // draw
        }
    }
    }
    panic("Unrecognized input")
}

func parse_move(s byte) Move {
    if s == byte('A') {
        return Rock
    } else if s == byte('B') {
        return Paper        
    } else if s == byte('C') {
        return Scissors
    }

    panic("Unrecognized move")
}

func parse_our_move(s byte, them Move) Move {
    if s == byte('X') {
        // end in a loss
        switch them {
        case Rock:
            return Scissors
        case Paper:
            return Rock
        case Scissors:
            return Paper
        }        
    } else if s == byte('Y') {
        // end in a draw
        return them
    } else if s == byte('Z') {
        // end in a win
        switch them {
        case Rock:
            return Paper
        case Paper:
            return Scissors
        case Scissors:
            return Rock
        }
    }

    panic("Unrecognized move")
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    score := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        them := parse_move(line[0])
        us := parse_our_move(line[2], them)
        score += compute_score(us, them)
    }
    fmt.Println("Score: %d\n", score)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
