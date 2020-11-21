package main

import "fmt"

// NetPrice calculate final price after discount
func NetPrice(price int, discount int) float64 {
	var final float64
	final = (1 - float64(discount)/100) * float64(price)
	return final
}

func main() {
	var title string
	var copies int
	var author string
	var offer bool
	var discount int

	// Book represents information about a book
	type Book struct {
		Title           string
		Author          string
		Copies          int
		Series          string
		Sequence        int
		PriceCents      int
		DiscountPercent int
		Edition         int
	}
	offer = true
	discount = 10
	title = "The Happiness Project"
	copies = 99
	author = "John Arundel"

	// using the Book struct
	b := Book{
		Title:           "The Philosopher's Stone",
		Author:          "J K Rowling",
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
}
