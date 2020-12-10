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

type passport struct {
	byr int
	iyr int
	eyr int
	hgt
	hcl string
	ecl string
	pid int
	cid int
}

type hgt struct {
	cm int
	in int
}

func ReadInputBlocks(file *os.File) [][]string {
	var blocks [][]string
	//this double bracket seems to allow it to take a slice in a slice (TODO learn more)
	var lineCatcher []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()

		// if line if not empty add it to lineCatcher
		if len(l) != 0 {
			lineCatcher = append(lineCatcher, l)
		}
		if len(l) == 0 && len(lineCatcher) != 0 {
			//if current file line is empty & there is something in lineCatcher: Put caught lines into block
			blocks = append(blocks, lineCatcher)

			//empty line catcher after each block
			lineCatcher = []string{}
		}
	}

	return blocks
}

func mapBlock(blocks []string) passport {
	p := passport{}

	for _, b := range blocks {

		//split by white space to give single name value pairs divided by :
		s := strings.Fields(b)

		for _, v := range s {

			l := strings.Split(v, ":")
			switch l[0] {
			case "byr":
				n, _ := strconv.Atoi((l[1]))
				p.byr = n
			case "iyr":
				n, _ := strconv.Atoi((l[1]))
				p.iyr = n
			case "eyr":
				n, _ := strconv.Atoi((l[1]))
				p.eyr = n
			case "hgt":
				//extract unit from last 2 chars
				unit := l[1][len(l[1])-2:]
				//depending on unit convert to in and put into specific struct value
				if unit == "cm" {
					n, _ := strconv.Atoi((strings.TrimSuffix(l[1], "cm")))
					p.hgt.cm = n
				} else if unit == "in" {
					n, _ := strconv.Atoi((strings.TrimSuffix(l[1], "in")))
					p.hgt.in = n
				}
			case "hcl":
				p.hcl = l[1]
			case "ecl":
				p.ecl = string(l[1])
			case "pid":
				n, _ := strconv.Atoi((l[1]))
				p.pid = n
			case "cid":
				n, _ := strconv.Atoi((l[1]))
				p.cid = n

			}
		}
	}
	return p
}

func validateBlock(p passport) bool {
	// fmt.Println("byr:", p.byr, "iyr:", p.iyr, "cid:", p.cid, "ecl:", p.ecl, "eyr:", p.eyr, "hcl:", p.hcl, "pid:", p.pid, "hgt-cm:", p.hgt.cm, "hgt-in:", p.hgt.in)

	valid := true

	//byr
	if validateIntRange(p.byr, 1920, 2002) == false {
		// fmt.Println("invalid byr")
		valid = false
	}
	//iyr
	if validateIntRange(p.iyr, 2010, 2020) == false {
		// fmt.Println("invalid iyr")
		valid = false
	}
	//eyr
	if validateIntRange(p.eyr, 2020, 2030) == false {
		// fmt.Println("invalid eyr")
		valid = false
	}
	//pid
	if validateIntRange(p.pid, 000000001, 999999999) == false {
		// fmt.Println("invalid pid")
		valid = false
	}
	//hgt
	if p.hgt.cm != 0 {
		//hgt cm
		if validateIntRange(p.hgt.cm, 150, 193) == false {
			// fmt.Println("invalid hgt.cm")
			valid = false
		}
	} else if p.hgt.in != 0 {
		//hgt in
		if validateIntRange(p.hgt.in, 59, 76) == false {
			// fmt.Println("invalid hgt.in")
			valid = false
		}
	} else {
		//if no value in cm or in then fail as required field
		valid = false
	}
	//ecl
	switch p.ecl {
	case "amb", "oth", "brn", "grn", "gry", "hzl", "blu":
		{
			//valid - do nothing
		}
	default:
		{
			valid = false
			// fmt.Println("invalid ecl", p.ecl)
		}
	}
	//hcl
	if r, _ := regexp.MatchString("#[0-9a-f]{6}", p.hcl); r == false {
		// fmt.Println("invalid hcl", err) - add back err to above line to get err
		valid = false
	}
	return valid

}

//function to accept int range and check it is valid between them, returns true if valid, false if not
func validateIntRange(v int, min int, max int) bool {
	if v >= min && v <= max {
		return true
	}
	return false
}

func main() {
	//open file & error check
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//split into blocks
	blocks := ReadInputBlocks(file)
	// validate each block for validity
	count := 0
	for _, s := range blocks {

		// fmt.Println("start block-----------------------------")
		p := validateBlock(mapBlock(s))
		if p == true {
			count++
			// fmt.Println("This one was valid above---------------------------")

		}
	}
	fmt.Println(count)

}
