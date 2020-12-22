package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	m := loadMap("input")

	xStep := 3
	yStep := 1

	trees := countTrees(m, xStep, yStep)

	log.Println("Tree count ", trees)
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

		if y == len(m) {
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
