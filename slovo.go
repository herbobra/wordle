package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

func word() string {
	file, err := os.Open("123.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	myString := string(b[:])
	stringa := strings.Split(myString, "\n")
	message := fmt.Sprint(stringa[rand.Intn(len(stringa))])

	fmt.Println(message)
	return message
}
