package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type input struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

type parse struct {
	name    string
	matcher *regexp.Regexp
}

func main() {
	passports := loadPassports("input")

	count := 0
	for _, p := range passports {
		if isValidPassport(p) {
			count++
		}
	}

	log.Println("Valid passports: ", count)
}

func isValidPassport(in input) bool {
	return in.byr != "" &&
		in.ecl != "" &&
		in.eyr != "" &&
		in.hcl != "" &&
		in.hgt != "" &&
		in.iyr != "" &&
		in.pid != ""
}

func loadPassports(ext string) (p []input) {
	file, err := os.Open(fmt.Sprintf("./day_04.%s", ext))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reg := input{byr: "", iyr: "", eyr: "", hgt: "", hcl: "", ecl: "", pid: "", cid: ""}

	matchers := []parse{
		parse{name: "byr", matcher: regexp.MustCompile("byr:[\\w]+")},
		parse{name: "iyr", matcher: regexp.MustCompile("iyr:[\\w]+")},
		parse{name: "eyr", matcher: regexp.MustCompile("eyr:[\\w]+")},
		parse{name: "hgt", matcher: regexp.MustCompile("hgt:[\\w]+")},
		parse{name: "hcl", matcher: regexp.MustCompile("hcl:[#\\w]+")},
		parse{name: "ecl", matcher: regexp.MustCompile("ecl:[#\\w]+")},
		parse{name: "pid", matcher: regexp.MustCompile("pid:[#\\w]+")},
		parse{name: "cid", matcher: regexp.MustCompile("cid:[\\w]+")},
	}

	for scanner.Scan() {
		l := scanner.Text()

		if l == "" {
			p = append(p, reg)

			reg = input{byr: "", iyr: "", eyr: "", hgt: "", hcl: "", ecl: "", pid: "", cid: ""}
		} else {
			for _, p := range matchers {
				v := p.matcher.FindStringSubmatch(l)
				if len(v) == 1 {
					switch name := p.name; name {
					case "byr":
						reg.byr = v[0]
					case "iyr":
						reg.iyr = v[0]
					case "eyr":
						reg.eyr = v[0]
					case "hgt":
						reg.hgt = v[0]
					case "hcl":
						reg.hcl = v[0]
					case "ecl":
						reg.ecl = v[0]
					case "pid":
						reg.pid = v[0]
					case "cid":
						reg.cid = v[0]
					}
				}
			}
		}
	}

	p = append(p, reg)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return p
}
