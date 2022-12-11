package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
  path := "./input.txt"

  file, err := os.OpenFile(path, os.O_RDONLY, 0644)
  if err != nil {
    panic(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  if err := scanner.Err(); err != nil {
    panic(err)
  }

  counter := 0

  for scanner.Scan() {
    line := scanner.Text()
    pairElf := strings.Split(line, ",")
    sectionElf1 := pairElf[0]
    sectionElf2 := pairElf[1]

    elf1Min, elf1Max := getSections(sectionElf1)
    elf2Min, elf2Max := getSections(sectionElf2)

    // TODO: Pretty sure I can refactor this
    // TODO: Yup, I'm pretty sure I can refactor this
    condPart1 := elf1Min >=elf2Min && elf1Max <= elf2Max || elf2Min >=elf1Min && elf2Max <= elf1Max
    condOverlap := elf1Max == elf2Min || elf2Min >= elf1Min && elf2Min <= elf1Max || elf1Min >= elf2Min && elf1Min <= elf2Max

    if condPart1 || condOverlap {
      counter += 1
    }
  }

  fmt.Println("The count is", counter)
}

func getSections (str string) (int, int) {
  sections := strings.Split(str, "-")
  min, _ := strconv.Atoi(sections[0])
  max, _ := strconv.Atoi(sections[1])

  return min, max
}
