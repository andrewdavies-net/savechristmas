package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type movement struct {
	right int
	down  int
}

func main() {

	//make slice of 5 with data type movement
	moves := make([]movement, 5)

	//set values
	moves[0] = movement{
		right: 1,
		down:  1,
	}
	moves[1] = movement{
		right: 3,
		down:  1,
	}
	moves[2] = movement{
		right: 5,
		down:  1,
	}
	moves[3] = movement{
		right: 7,
		down:  1,
	}
	moves[4] = movement{
		right: 1,
		down:  2,
	}

	fmt.Println(moves)
	total := 1

	for _, treeRun := range moves {
		fmt.Println(treeRun)

		number := 0
		trees := 0

		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		lines := 0
		for scanner.Scan() {
			//check line count

			//mod
			if lines%treeRun.down == 0 {

				// fmt.Println("Current number", number, " :", string(scanner.Text()[number]), (string(scanner.Text())))

				if string(scanner.Text()[number]) == "#" {
					trees++
				}
				// fmt.Println(string(scanner.Text()[number]))

				number = number + treeRun.right
				if number > 30 {
					number = number - 31
				}

				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			}
			//keeping line number tally
			lines++

		}
		//end of section tally
		fmt.Println(trees)
		total = total * trees
	}
	fmt.Println(total)

}
