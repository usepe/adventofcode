package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	nums := loadNums()

	for i, n := range nums {
		for j, n2 := range nums {
			if i == j {
				continue
			}
			if sumsUp(n, n2) {
				fmt.Println(fmt.Sprintf("Nums %d and %d sums 2020 and multiply to %d", n, n2, n*n2))
			}
		}
	}
}

func loadNums() (nums []int) {
	file, err := os.Open("./day_01.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}

func sumsUp(a, b int) bool {
	return a+b == 2020
}
