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
			for k, n3 := range nums {
				if i == j || i == k || j == k {
					continue
				}
				if sumsUp(n, n2, n3) {
					fmt.Println(fmt.Sprintf("Nums %d, %d and %d sums 2020 and multiply to %d", n, n2, n3, n*n2*n3))
				}
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

func sumsUp(a, b, c int) bool {
	return a+b+c == 2020
}
