package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {

	// Запись строки в кодировке Windows-1252
	encoder := charmap.Windows1252.NewEncoder()
	s, e := encoder.String("This is sample text with runes Š")
	if e != nil {
		panic(e)
	}
	io.WriteFile("123.txt", []byte(s), os.ModePerm)

	// Декодировка в UTF-8
	f, e := os.Open("123.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
