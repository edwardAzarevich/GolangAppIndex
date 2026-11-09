package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader = bufio.NewReader(os.Stdin)

var UserBookMarks = map[int]string{}

func getPageAndInfo() (int, string) {
	var userKey int
	var userValue string
	fmt.Println("Please input page:")
	fmt.Scan(&userKey)
	fmt.Println("Please input info:")
	fmt.Scan(&userValue)
	return userKey, userValue
}

func getDeletePage() int {
	var userKey int
	fmt.Println("Please input page:")
	fmt.Scan(&userKey)
	return userKey
}

func addBookmark() {
	userKey, userValue := getPageAndInfo()
	UserBookMarks[userKey] = userValue
	fmt.Println("Added success!")
}

func deleteBookmark() {
	userKey := getDeletePage()
	delete(UserBookMarks, userKey)
	fmt.Println("Added success!")
}

func getMap(actions map[int]string) {
	for key, value := range actions {
		fmt.Printf("%d - %s,\n", key, value)
	}
}

func menuAction(actions map[int]string, userAction int) bool {
	switch actions[userAction] {
	case "view bookmarks":
		fmt.Println("view bookmarks")
		getMap(UserBookMarks)
	case "add bookmark":
		fmt.Println("add")
		addBookmark()
	case "delete bookmark":
		fmt.Println("delete")
		deleteBookmark()
	case "exit":
		fmt.Println("Goodbye!!!")
		return false
	default:
		fmt.Println("invalid value, please try again!")
	}
	return true
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
	getMap(actions)
	for {
		fmt.Scan(&userInput)
		shouldContinue := menuAction(actions, userInput)
		if !shouldContinue {
			break
		}
		getMap(actions)
	}
}
