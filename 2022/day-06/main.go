package main

import (
	"fmt"
	"os"
	"strings"
)

// I think I can do this with a Set
// since Go doesn't have an impl for it...

type Set struct {
	items map[string]struct{}
}

func (s *Set) Add(el string) {
	s.items[el] = struct{}{}
}

func newSet() *Set {
	var set Set
	set.items = make(map[string]struct{})

	return &set
}

func main() {
	path := "./input.txt"

	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	buffer := strings.TrimSpace(string(file))

	const MARKER_LENGTH int = 4

	i := 0
	set := newSet()

	for {
    length := i +MARKER_LENGTH
		for _, char := range buffer[i : length] {
			set.Add(string(char))
		}

    if len(set.items) == MARKER_LENGTH {
      fmt.Println("The number of character before the first start-of-packet marker is", length)
			break
		}

    // Reset the set (lul)
    set = newSet()
		i++
	}
}
