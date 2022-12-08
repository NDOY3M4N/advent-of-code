package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scoreP2 uint

func main() {
	path := "./input.txt"
	content := readInput(path)

	for _, element := range content {
		playersPick := strings.Split(element, " ")
		pick1 := playersPick[0]
		pick2 := playersPick[1]

		switch pick1 {
		case "A":
			pick1 = "rock"
			break
		case "B":
			pick1 = "paper"
			break
		case "C":
			pick1 = "scissor"
			break
		}

		pick2 = decipherP2(pick2)

		winner := getWinnerCheat(pick1, pick2)
		handleScoring(winner)
	}

	fmt.Println("My score is", scoreP2)
}

func decipherP2(str string) string {
	result := ""

	switch str {
	case "X":
		result = "lose"
		break
	case "Y":
		result = "draw"
		break
	case "Z":
		result = "win"
		break
	}

	return result
}

func handleScoring(winner string) {
	switch winner {
	case "draw":
		scoreP2 += 3
		break
	case "P2 wins":
		scoreP2 += 6
		break
	case "P1 wins":
		scoreP2 += 0
	}
}

func cheat(instruction string) string {
	switch instruction {
	case "win":
		return "P2 wins"
	case "lose":
		return "P1 wins"
	case "draw":
		return "draw"
	default:
		return ""
	}
}

func getWinnerCheat(pick1, pick2 string) string {
	switch pick1 {
	case "rock":
		if pick2 == "win" {
			// player 2 should pick paper
			scoreP2 += 2

			return "P2 wins"
		} else if pick2 == "lose" {
			// player 2 should pick scissor
			scoreP2 += 3

			return "P1 wins"
		} else if pick2 == "draw" {
			// player 2 should pick rock
			scoreP2 += 1

			return "draw"
		}
	case "paper":
		if pick2 == "win" {
			scoreP2 += 3

			return "P2 wins"
		} else if pick2 == "lose" {
			// player 2 should pick scissor
			scoreP2 += 1

			return "P1 wins"
		} else if pick2 == "draw" {
			// player 2 should pick rock
			scoreP2 += 2

			return "draw"
		}
	case "scissor":
		if pick2 == "win" {
			scoreP2 += 1

			return "P2 wins"
		} else if pick2 == "lose" {
			scoreP2 += 2

			return "P1 wins"
		} else if pick2 == "draw" {
			scoreP2 += 3

			return "draw"
		}
	}

	return ""
}

func getWinner(pick1, pick2 string) string {
	if pick1 == pick2 {
		return "draw"
	}

	switch pick1 {
	case "rock":
		if pick2 == "paper" {
			return "P2 wins"
		}
	case "paper":
		if pick2 == "scissor" {
			return "P2 wins"
		}
	case "scissor":
		if pick2 == "rock" {
			return "P2 wins"
		}
	}

	return "P1 wins"
}

func readInput(path string) []string {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
