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
	First    int
	Second   int
	Letter   string
	Password string
}

func main() {
	passwords := loadPasswords()

	count := 0
	for _, in := range passwords {
		log.Println("Input: ", in)
		log.Println("Result: ", validPassword(in))
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
		first, second := atoi(s[0]), atoi(s[1])

		passwords = append(passwords, input{First: first, Second: second, Letter: letter, Password: pass})
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
	c1 := string(i.Password[i.First])
	c2 := string(i.Password[i.Second])

	containCount := 0
	if c1 == i.Letter {
		containCount++
	}
	if c2 == i.Letter {
		containCount++
	}

	log.Println("First: ", c1)
	log.Println("Second: ", c2)
	log.Println("Count: ", containCount)

	return containCount == 1
}
