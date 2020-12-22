package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type input struct {
	Lower    int
	Upper    int
	Letter   string
	Password string
}

func main() {
	passwords := loadPasswords()

	count := 0
	for _, in := range passwords {
		if validPassword(in) {
			count++
		}
	}

	log.Println(fmt.Sprintf("Valid passwords %d", count))
}

func loadPasswords() (passwords []input) {
	file, err := os.Open("./day_02.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, ":")
		s, pass := strings.Split(s[0], " "), s[1]
		s, letter := strings.Split(s[0], "-"), s[1]
		lower, upper := atoi(s[0]), atoi(s[1])

		passwords = append(passwords, input{Lower: lower, Upper: upper, Letter: letter, Password: pass})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passwords
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func validPassword(i input) bool {
	qty := strings.Count(i.Password, i.Letter)
	return qty >= i.Lower && qty <= i.Upper
}
