package helper

import (
	"fmt"
	"strings"
)

func SearchBooks(books, input_arguments []string) {

	books_list := IsASubSlice(books, input_arguments)

	for i := 0; i < len(books_list); i++ {
		fmt.Println(books_list[i])
	}

}
func ListBooks(books []string) {
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i])
	}
}

func IsASubSlice(slice_one, slice_two []string) []string {
	var return_slice []string

	//making a string from a string array for strings.contains method
	str := strings.Join(slice_two, " ")

	//eliminating upper-lower case problem
	str = strings.ToUpper(str)

	if len(slice_one) < len(slice_two) {
		return return_slice
	}

	for i := 0; i < len(slice_one); i++ {
		if strings.Contains(strings.ToUpper(slice_one[i]), str) {
			return_slice = append(return_slice, slice_one[i])
		}
	}
	return return_slice
}
