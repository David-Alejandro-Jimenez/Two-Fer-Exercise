package main

import "fmt"

func shareWith(name string) string {
	if name == "" {
		fmt.Println("One for you, one for me.")
		return "One for you, one for me."
	} else {
		fmt.Println("One for " + name + ", one for me.")
		return "One for " + name + ", one for me."
	}
}

func names() {
	var name1 = "Alice"
	var name2 = "Bob"
	var name3 = ""
	shareWith(name1)
	shareWith(name2)
	shareWith(name3)
}

func main() {
	names()
}