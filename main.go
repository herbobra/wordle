package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	a := word()
	secretWord := a[:(len(a) - 1)] // Загаданное слово
	maxAttempts := 5               // Максимальное количество попыток

	fmt.Println("Добро пожаловать в игру Wordle!")
	fmt.Printf("Попробуйте отгадать загаданное слово из %d букв.\n", len(secretWord))

	attempts := 0
	success := false

	for attempts < maxAttempts && !success {
		fmt.Printf("Попытка %d: ", attempts+1)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if len(input) != len(secretWord) {
			fmt.Printf("Ваш ответ должен состоять из %d букв.\n", len(secretWord))
			continue
		}

		result := ""

		for i := 0; i < len(secretWord); i++ {
			if input[i] == secretWord[i] {
				result += string(input[i])
			} else if strings.ContainsRune(secretWord, rune(input[i])) {
				result += "+"
			} else {
				result += "-"
			}
		}

		fmt.Println(result)

		if result == secretWord {
			success = true
			fmt.Println("Вы отгадали слово!")
		} else {
			attempts++
		}
	}

	fmt.Println("Игра окончена.")
	if !success {
		fmt.Printf("Загаданное слово было: %s\n", secretWord)
	}
}
