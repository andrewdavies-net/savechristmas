package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	number := 0
	trees := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//print out for debug
		fmt.Println("Current number", number, " :", string(scanner.Text()[number]), (string(scanner.Text())))
		if string(scanner.Text()[number]) == "#" {
			trees++
		}
		// fmt.Println(string(scanner.Text()[number]))

		number = number + 3
		if number > 30 {
			number = number - 31
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

	}

	fmt.Println(trees)
}
