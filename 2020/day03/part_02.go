package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type input struct {
	xStep int
	yStep int
}

func main() {
	m := loadMap("input")

	slopes := []input{
		input{xStep: 1, yStep: 1},
		input{xStep: 3, yStep: 1},
		input{xStep: 5, yStep: 1},
		input{xStep: 7, yStep: 1},
		input{xStep: 1, yStep: 2},
	}

	acc := 1

	for _, v := range slopes {
		trees := countTrees(m, v.xStep, v.yStep)
		acc = acc * trees

		log.Println(fmt.Sprintf("Slope (%d, %d) Tree count: %d", v.xStep, v.yStep, trees))
	}

	log.Println("Altogether: ", acc)
}

func countTrees(m []string, xStep, yStep int) int {
	xLen := len(m[0])
	x := 0
	y := 0

	trees := 0

	for {
		t := string(m[y][x])

		if t == "#" {
			trees++
		}

		x += xStep
		y += yStep

		if x >= xLen {
			x -= xLen
		}

		if y >= len(m) {
			break
		}
	}

	return trees
}

func loadMap(ext string) (m []string) {
	file, err := os.Open(fmt.Sprintf("./day_03.%s", ext))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return m
}
