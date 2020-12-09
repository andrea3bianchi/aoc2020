package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// don't want to parse this elegantly at all, hammertime
	dump, _ := ioutil.ReadFile("data")
	passports := strings.Split((string(dump)), "\n\n")

	var validPassports int
	for _, p := range passports {
		valid := true
		fields := map[string]string{
			"byr": "",
			"iyr": "",
			"eyr": "",
			"hgt": "",
			"hcl": "",
			"ecl": "",
			"pid": "",
			"cid": "",
		}
		p = strings.Replace(p, "\n", " ", -1) // getting rid of newlines to split properly below
		elements := strings.Split(p, " ")
		for _, e := range elements {
			kv := strings.Split(e, ":")
			fields[kv[0]] = kv[1]
		}
		for f := range fields {
			if fields[f] == "" && f != "cid" {
				// fmt.Printf("%d invalid because %s is not present\n%v\n",
				// 	i, f, fields)
				valid = false
				continue
			}
			switch f {
			case "byr":
				if !checkDate(fields[f], 1920, 2002) {
					valid = false
				}
			case "iyr":
				if !checkDate(fields[f], 2010, 2020) {
					valid = false
				}
			case "eyr":
				if !checkDate(fields[f], 2020, 2030) {
					valid = false
				}
			case "hgt":
				if len(fields[f]) < 2 {
					valid = false
					break
				}
				prefix := fields[f][len(fields[f])-2:]
				if prefix != "cm" && prefix != "in" {
					valid = false
					break
				}
				ns := fields[f][:len(fields[f])-2]
				switch prefix {
				case "cm":
					n, err := strconv.Atoi(ns)
					if err != nil {
						valid = false
						break
					}
					if n < 150 || n > 193 {
						valid = false
					}
				case "in":
					n, err := strconv.Atoi(ns)
					if err != nil {
						valid = false
						break
					}
					if n < 59 || n > 76 {
						valid = false
					}
				}
			case "hcl":
				if len(fields[f]) == 0 {
					valid = false
					break
				}
				if fields[f][0] != '#' {
					valid = false
					break
				}
				println(fields[f][1:])
				// it is at this point that I realized i should have been writing regex all along

				if !regexp.MustCompile(`^[a-f0-9_]{6}$`).MatchString(fields[f][1:]) {
					valid = false
					break
				}
			case "ecl":
				imhatingthis := map[string]bool{
					"amb": true,
					"blu": true,
					"brn": true,
					"gry": true,
					"grn": true,
					"hzl": true,
					"oth": true,
				}
				if !imhatingthis[fields[f]] {
					valid = false
				}
			case "pid":
				if !regexp.MustCompile(`^[0-9_]{9}$`).MatchString(fields[f]) {
					valid = false
				}
			}
		}
		if valid {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

func checkDate(s string, min, max int) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if i < min || i > max {
		return false
	}
	return true
}
