package main

import (
	"fmt"
	helper "homework-1-ecepinaryildirim/pkg"
	"os"
)

func main() {

	books := []string{"Harry Potter and the Prisoner of Azkaban", "Harry Potter and the Goblet of Fire", "A Song of Ice and Fire", "Replay", "Everything's Eventual", "The Shining"}
	var input_arguments []string
	args := os.Args

	if len(args) == 1 {
		fmt.Printf("Commands that are available :\n search => for searching the book in the list \n list => for listing the books\n")
		return
	}

	if args[1] == "search" {
		//after the search word, all words are going to be searched inside books
		for i := 2; i < len(args); i++ {
			input_arguments = append(input_arguments, args[i])
		}
		helper.SearchBooks(books, input_arguments)

	} else if args[1] == "list" {
		helper.ListBooks(books)

	} else {
		//user has entered invalid input
		fmt.Println("Invalid input!")
		return
	}

}
