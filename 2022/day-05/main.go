package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Wanted to test my knowledge of Stack data struct
// So far it is working for the part 1 but for the part 2 ...
type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Delete the last item inserted into the stack
// and return the deleted item
func (s *Stack) Pop() string {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]

	s.items = s.items[:lastIndex]

	return lastItem
}

func (s *Stack) GetTop() string {
	lastIndex := len(s.items) - 1

	return s.items[lastIndex]
}

func (s *Stack) Pop9001(qty int) (itemsToMove []string) {
	l := len(s.items) - qty
	itemsToMove = s.items[l:]

  s.items = s.items[:l]
	return itemsToMove
}

func (s *Stack) Push9001(items []string) {
	for _, item := range items {
		s.Push(item)
	}
}

// TODO: automate this process, don't be lazy create a func to parse it.
// I will manually create the Stacks (crates)
var myStacks = map[int]*Stack{
	1: {items: []string{"B", "G", "S", "C"}},
	2: {items: []string{"T", "M", "W", "H", "J", "N", "V", "G"}},
	3: {items: []string{"M", "Q", "S"}},
	4: {items: []string{"B", "S", "L", "T", "W", "N", "M"}},
	5: {items: []string{"J", "Z", "F", "T", "V", "G", "W", "P"}},
	6: {items: []string{"C", "T", "B", "G", "Q", "H", "S"}},
	7: {items: []string{"T", "J", "P", "B", "W"}},
	8: {items: []string{"G", "D", "C", "Z", "F", "T", "Q", "M"}},
	9: {items: []string{"N", "S", "H", "B", "P", "F"}},
}

func main() {
	path := "./input.txt"
	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	parsedString := strings.Split(string(file), "\n\n")
	commands := strings.Split(parsedString[1], "\n")

	for _, command := range commands {
		// Somehow I have an empty line added to the slice
		if len(command) > 0 {
			qty, from, to := parseCommand(command)
			// Part 1
			// operateCrateMover9000(myStacks, qty, from, to)
			// Part 2
			operateCrateMover9001(myStacks, qty, from, to)
		}
	}

	fmt.Println("The crates that end up at the top are:", topCrates(myStacks))
}

func parseCommand(c string) (qty, from, to int) {
	qty, _ = strconv.Atoi(strings.Split(c, " ")[1])
	from, _ = strconv.Atoi(strings.Split(c, " ")[3])
	to, _ = strconv.Atoi(strings.Split(c, " ")[5])

	return
}

func operateCrateMover9000(s map[int]*Stack, qty, from, to int) {
	for qty > 0 {
		itemToMove := s[from].Pop()
		s[to].Push(itemToMove)

		qty--
	}
}

func operateCrateMover9001(s map[int]*Stack, qty, from, to int) {
	itemsToMove := s[from].Pop9001(qty)
	s[to].Push9001(itemsToMove)
}

func topCrates(s map[int]*Stack) (result string) {
	for i := 1; i <= len(myStacks); i++ {
		result += s[i].GetTop()
	}

	return
}
