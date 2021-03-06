package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type password struct {
	allowedRangeStart int
	allowedRangeEnd   int
	character         string
	passswordString   string
}

func main() {

	count := 0

	file, err := os.Open("passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//split fields by space
		fields := strings.Fields(scanner.Text())

		// fmt.Println(fields)
		//break out into numbers before and after -
		AllowedRange := strings.Split(fields[0], "-")
		//convert to int
		ARLInt, err := strconv.Atoi(AllowedRange[0])
		ARUInt, err := strconv.Atoi(AllowedRange[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		// check convert to int worked

		//add split data into structure condition
		condition := password{
			allowedRangeStart: ARLInt - 1,
			allowedRangeEnd:   ARUInt - 1,
			character:         strings.Trim(fields[1], ":"),
			passswordString:   fields[2],
		}

		result := evaluateCondition(condition)
		if result == true {
			count++
		}

	}
	//check scanner worked correctly
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

}

//evaluation function pass data type password
func evaluateCondition(condition password) bool {

	if string(condition.passswordString[condition.allowedRangeStart]) == condition.character && string(condition.passswordString[condition.allowedRangeEnd]) == condition.character {
		return false
	} else if string(condition.passswordString[condition.allowedRangeStart]) == condition.character || string(condition.passswordString[condition.allowedRangeEnd]) == condition.character {
		return true
	} else {
		return false
	}

}
