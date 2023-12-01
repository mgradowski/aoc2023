package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var replacerFwd = strings.NewReplacer(
	"one", "1",
	"two", "2",
	"three", "3",
	"four", "4",
	"five", "5",
	"six", "6",
	"seven", "7",
	"eight", "8",
	"nine", "9")

var replacerRev = strings.NewReplacer(
	"eno", "1",
	"owt", "2",
	"eerht", "3",
	"ruof", "4",
	"evif", "5",
	"xis", "6",
	"neves", "7",
	"thgie", "8",
	"enin", "9")

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func parse(line string) int {
	fixedFwd := replacerFwd.Replace(line)
	fixedRev := replacerRev.Replace(reverse(line))
	iFst := strings.IndexFunc(fixedFwd, unicode.IsDigit)
	iLstRev := strings.IndexFunc(fixedRev, unicode.IsDigit)
	return 10*int(fixedFwd[iFst]-'0') + int(fixedRev[iLstRev]-'0')
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
