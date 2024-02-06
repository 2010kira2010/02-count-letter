package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	counts := make(map[rune]int)
	total := 0

	for _, char := range text {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			counts[unicode.ToLower(char)]++
			total++
		}
	}

	fmt.Println("Результат:")
	for char, count := range counts {
		percentage := float64(count) / float64(total) * 100
		fmt.Printf("Буква \"%c\" встречается в предложении %d раз(а). Частота встречаемости: %.2f%%\n", char, count, percentage)
	}
}
