package main

import "fmt"

type Book struct {
	Title           string
	Author          []string
	Copies          int
	Series          string
	Sequence        int
	PriceCents      int
	DiscountPercent int
	Edition         int
	Featured        bool
}

// NetPrice calculate final price after discount
func NetPrice(price int, discount int) float64 {
	var final float64
	final = (1 - float64(discount)/100) * float64(price)
	return final
}

func AddToCatalog(newBook Book, catalog []Book) []Book {
	catalog = append(catalog, newBook)
	return catalog
}

func FeaturedBooks(catalog []Book) []Book {
	var featuredCatalog []Book
	for _, b := range catalog {
		if b.Featured == true {
			featuredCatalog = AddToCatalog(b, featuredCatalog)
		}
	}
	return featuredCatalog
}

func main() {
	var title string
	var copies int
	var author string
	var offer bool
	var discount int

	// Book represents information about a book

	offer = true
	discount = 10
	title = "The Happiness Project"
	copies = 99
	author = "John Arundel"

	// using the Book struct
	b := Book{
		Title:           "The Philosopher's Stone",
		Author:          []string{"J K Rowling", "John Arundel"},
		Copies:          8,
		Series:          "Harry Potter",
		Sequence:        1,
		PriceCents:      100,
		DiscountPercent: 10,
	}

	fmt.Println(title)
	fmt.Println(copies)
	fmt.Println(author)
	fmt.Println(offer)
	fmt.Println(discount)
	fmt.Printf("%+v\n", b)
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Series:", b.Series)
	fmt.Println("Sequence:", b.Sequence)
	fmt.Println("Copies:", b.Copies)
	fmt.Println("Edition:", b.Edition)
	fmt.Println("Discount Percent:", b.DiscountPercent)
	fmt.Println("NetPrice:", NetPrice(b.PriceCents, b.DiscountPercent))

	// Slices
	stringSlice := []string{
		"This is a string",
		"This is another string",
	}
	fmt.Println("String slice", stringSlice)

	// book slice
	catalog := []Book{
		{Title: "Delightfully Uneventful Trip on the Orient Express"},
		{Title: "One Hundred Years of Good Company"},
	}
	fmt.Println(catalog)
	fmt.Println(catalog[0])
	fmt.Println(catalog[1])
	fmt.Println(len(catalog))

	// Index out of range
	//fmt.Println(catalog[2])
	catalog[0] = Book{Title: "Heart of Kindness"}
	fmt.Println(catalog[0].Title)
	catalog[0].Title = "The Grapes of Mild Irritation"
	fmt.Println(catalog[0].Title)
	c := Book{Title: "For the Love of Go", Author: []string{"John Arundel"}, Copies: 100, Featured: true}
	catalog = append(catalog, c)
	fmt.Println(catalog[2].Title)
	fmt.Println(catalog)
	// Iterating over a struct slice
	for i := range catalog {
		fmt.Println(i, ": ", catalog[i].Title)
	}

	d := Book{Title: "Talking with Tech Leads", Author: []string{"Pat Kua"}, Copies: 10}
	catalog = AddToCatalog(d, catalog)
	fmt.Println("New Catalog:", catalog)

	// Using the index returned by go
	for i, b := range catalog {
		fmt.Println(i, ": ", b.Title)
	}

	// Using _
	for _, b := range catalog {
		fmt.Println(b.Title, b.Author)
	}
	fmt.Println("Number of books in the catalog: ", len(catalog))

	// Using the FeaturedBooks function
	fmt.Println(FeaturedBooks(catalog))

}
