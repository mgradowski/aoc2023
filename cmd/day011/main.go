package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parse(line string) int {
	iFst := strings.IndexFunc(line, unicode.IsDigit)
	iLst := strings.LastIndexFunc(line, unicode.IsDigit)
	return 10*int(line[iFst]-'0') + int(line[iLst]-'0')
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		result += parse(line)
	}

	fmt.Println(result)
}
