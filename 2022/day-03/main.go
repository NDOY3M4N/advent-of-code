package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  path := "./input.txt"
  var sumsOfPriority int32
  // var sumsOfPriority2 int32

  file, err := os.OpenFile(path, os.O_RDONLY, 0644)
  if err != nil {
    panic(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  if err := scanner.Err(); err != nil {
    panic(err)
  }

  var fileContent []string
  for scanner.Scan() {
    line := scanner.Text()

    itemList1, itemList2 := splitStringHalf(line)
    duplicateItem := getDuplicate(itemList1, itemList2)
    
    sumsOfPriority += getPriority(duplicateItem)

    fileContent = append(fileContent, line)
  }
  fmt.Println("The sum of the priorities for these item type is", sumsOfPriority)

  // Part 2
  groupFactor := 0
  sumsOfPriority = 0

  for groupFactor < len(fileContent) {
    group := fileContent[groupFactor:groupFactor + 3]
    for _, item := range group[0] {
      condition := strings.ContainsRune(group[1], item) && strings.ContainsRune(group[2], item)
      if condition {
        sumsOfPriority += getPriority(item)
        break
      }
    }

    groupFactor += 3
  }

  fmt.Println("The sum of the priorities for these item type is", sumsOfPriority)
}

func getPriority(char rune) rune {
  const lowercase_factor rune = 'a' - 1
  const uppercase_factor rune = 'A' - 27

  if char >= 'a' && char <= 'z' {
    return char - lowercase_factor
  }

  return char - uppercase_factor
}

func getDuplicate(itemList1, itemList2 string) rune {
  var result rune
  for _, item := range itemList1 {
    if strings.ContainsRune(itemList2, item) {
      result = item
    }
  }

  return result
}

func splitStringHalf(str string) (string, string) {
  middle := len(str) / 2
  return str[:middle], str[middle:]
}
