package main

import (
	"flag"
	"fmt"
	"strings"
)

func t19_1() {
	strPtr := flag.String("str", "", "исходная строка")
	substrPtr := flag.String("substr", "", "подстрока для поиска")
	flag.Parse()

	if *strPtr == "" || *substrPtr == "" {
		fmt.Println("Необходимо указать оба флага -str и -substr")
		return
	}

	found := findSubstring(*strPtr, *substrPtr)

	fmt.Println(found)
}

func findSubstring(str, substr string) bool {
	strRunes := []rune(str)
	substrRunes := []rune(substr)

	for i := 0; i <= len(strRunes)-len(substrRunes); i++ {
		if strings.EqualFold(string(strRunes[i:i+len(substrRunes)]), substr) {
			return true
		}
	}

	return false
}

func main() {
	t19_1()
}
