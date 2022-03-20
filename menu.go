package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	for true {
		fmt.Println("==========================")
		fmt.Println("This is Menu Page")
		fmt.Println("1.list")
		fmt.Println("2.help")
		fmt.Println("3.quit")
		fmt.Println("==========================")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("This is List Page......")
		case 2:
			fmt.Println("This is help Page......")
		case 3:
			os.Exit(0)
		default:
			fmt.Println("wrong input, please re-enter")
		}
	}
}
