package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type input struct {
	byr int
	iyr int
	eyr int
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
	passports := loadPassports("invalid")

	count := 0
	for _, p := range passports {
		if isValidPassport(p) {
			count++
		}
	}

	log.Println("Valid passports: ", count)
}

func isValidPassport(in input) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if in.byr < 1920 {
		return false
	}
	if in.byr > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if in.iyr < 2010 {
		return false
	}
	if in.iyr > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if in.eyr < 2020 {
		return false
	}
	if in.eyr > 2030 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	c := 0
	if !strings.HasSuffix(in.hgt, "cm") {
		c++
	}
	if !strings.HasSuffix(in.hgt, "in") {
		if c == 1 {
			return false
		}
	}

	if strings.HasSuffix(in.hgt, "cm") {
		v, err := strconv.Atoi(strings.Replace(in.hgt, "cm", "", 1))
		if err != nil {
			return false
		}
		// If cm, the number must be at least 150 and at most 193.
		if v < 150 {
			return false
		}
		if v > 193 {
			return false
		}
	}
	if strings.HasSuffix(in.hgt, "in") {
		v, err := strconv.Atoi(strings.Replace(in.hgt, "in", "", 1))
		if err != nil {
			return false
		}
		// If in, the number must be at least 59 and at most 76.
		if v < 59 {
			return false
		}
		if v > 76 {
			return false
		}
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hex := regexp.MustCompile("#[a-f0-9]{6}")
	if !hex.MatchString(in.hcl) {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	possibles := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if _, ok := possibles[in.ecl]; !ok {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	validPid := regexp.MustCompile("[0-9]{9}")
	if !validPid.MatchString(in.pid) {
		return false
	}

	return true
}

func newInput() input {
	return input{byr: 0, iyr: 0, eyr: 0, hgt: "", hcl: "", ecl: "", pid: "", cid: ""}
}

func loadPassports(ext string) (p []input) {
	file, err := os.Open(fmt.Sprintf("./day_04.%s", ext))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reg := newInput()

	matchers := []parse{
		{name: "byr", matcher: regexp.MustCompile("byr:[\\w]+")},
		{name: "iyr", matcher: regexp.MustCompile("iyr:[\\w]+")},
		{name: "eyr", matcher: regexp.MustCompile("eyr:[\\w]+")},
		{name: "hgt", matcher: regexp.MustCompile("hgt:[\\w]+")},
		{name: "hcl", matcher: regexp.MustCompile("hcl:[#\\w]+")},
		{name: "ecl", matcher: regexp.MustCompile("ecl:[#\\w]+")},
		{name: "pid", matcher: regexp.MustCompile("pid:[#\\w]+")},
		{name: "cid", matcher: regexp.MustCompile("cid:[\\w]+")},
	}

	for scanner.Scan() {
		l := scanner.Text()

		if l == "" {
			p = append(p, reg)

			reg = newInput()
		} else {
			for _, p := range matchers {
				v := p.matcher.FindStringSubmatch(l)
				if len(v) == 1 {
					switch name := p.name; name {
					case "byr":
						v, err := strconv.Atoi(strings.Split(v[0], ":")[1])
						if err == nil {
							reg.byr = v
						}
					case "iyr":
						v, err := strconv.Atoi(strings.Split(v[0], ":")[1])
						if err == nil {
							reg.iyr = v
						}
					case "eyr":
						v, err := strconv.Atoi(strings.Split(v[0], ":")[1])
						if err == nil {
							reg.eyr = v
						}
					case "hgt":
						reg.hgt = strings.Replace(v[0], "hgt:", "", 1)
					case "hcl":
						reg.hcl = strings.Replace(v[0], "hcl:", "", 1)
					case "ecl":
						reg.ecl = strings.Replace(v[0], "ecl:", "", 1)
					case "pid":
						reg.pid = strings.Replace(v[0], "pid:", "", 1)
					case "cid":
						reg.cid = strings.Replace(v[0], "cid:", "", 1)
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
