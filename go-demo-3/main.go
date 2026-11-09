package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader = bufio.NewReader(os.Stdin)

func getAction(action map[int]string) {
	for key, value := range action {
		fmt.Printf("%d - %s,\n", key, value)
	}
}

func menuAction() {

}

func main() {
	var userInput int
	actions := map[int]string{
		1: "view bookmarks",
		2: "add bookmark",
		3: "delete bookmark",
		4: "exit",
	}
	fmt.Println("Welcome to application for bookmarks!")
	fmt.Println("Please choose your action:")
	getAction(actions)
	for {
		fmt.Scan(&userInput)
		switch actions[userInput] {
		case "view bookmarks":
			fmt.Println("view bookmarks")
		case "add bookmark":
			fmt.Println("add")
		case "delete bookmark":
			fmt.Println("delete")
		case "exit":
			fmt.Println("Goodbye!!!")
			return
		default:
			fmt.Println("invalid value, please try again!")
			continue
		}
		getAction(actions)
	}
}
