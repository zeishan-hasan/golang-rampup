// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// Go struct exercise

	type Book struct {
		title, author string
		year          int
	}

	// This is a struct literal and order matters
	book1 := Book{"Book no.1", "Author no.1", 2022}
	fmt.Println("book1:", book1)

	// Order does not matter
	book2 := Book{author: "Author2", title: "Title2", year: 2022}
	fmt.Println("book2:", book2)
	fmt.Println("title:", book2.title)

	store := book1.title
	fmt.Println("store:", store)

	// Copy struct
	book3 := book2
	fmt.Println("book3:", book3)
	book3.title = "Title3"
	fmt.Println("book3:", book3, "book2:", book2)
	// struct comparison
	fmt.Println("book1 == book2:", book1 == book2)
	fmt.Println("book1.year == book2.year", book1.year == book2.year)
	// fmt.Println("book1.year == book2.author", book1.year == book2.author) // Datatype should match for value comparison
}
