package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

var romanMap = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
var intToRoman = [...]struct { value int; symbol string }{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"},
	{90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"},
	{5, "V"}, {4, "IV"}, {1, "I"},
}

var validRomanPattern = `^(M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`

func isValidRoman(s string) bool {
	matched, _ := regexp.MatchString(validRomanPattern, s)
	return matched
}

func RomanToInt(s string) (int, error) {
	if !isValidRoman(s) {
		return 0, fmt.Errorf("Número romano inválido")
	}

	total, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if val := romanMap[rune(s[i])]; val < prev {
			total -= val
		} else {
			total += val
			prev = val
		}
	}
	return total, nil
}

func IntToRoman(num int) string {
	var result strings.Builder
	for _, entry := range intToRoman {
		for num >= entry.value {
			result.WriteString(entry.symbol)
			num -= entry.value
		}
	}
	return result.String()
}

func main() {
	var input string
	fmt.Print("Ingrese un número o un número romano: ")
	fmt.Scanln(&input)

	if num, err := strconv.Atoi(input); err == nil {
		fmt.Println("En romano es:", IntToRoman(num))
	} else {
		if result, err := RomanToInt(input); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("En número es:", result)
		}
	}
}

