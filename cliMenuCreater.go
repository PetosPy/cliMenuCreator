package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	r := bufio.NewReader(os.Stdin)

	// First item
	firstItem := itemName(r, "Enter the first item's name:")

	// First item price
	firstItemPrice, err := itemPrice(r, "Enter the first item's price:")
	errHandler(err)

	// Second item
	secondItem := itemName(r, "Enter thesecond item's name:")

	// Second item price
	secondItemPrice, err := itemPrice(r, "Enter the second item's price:")

	errHandler(err)

	// Third item
	thirdItem := itemName(r, "Enter the third item's name:")

	// Third item price
	thirdItemPrice, err := itemPrice(r, "Enter the third item's price:")

	errHandler(err)

	fmt.Printf("|  %-14s  | $  %-10.2f |\n", firstItem, firstItemPrice)
	fmt.Printf("|  %-14s  | $  %-10.2f |\n", secondItem, secondItemPrice)
	fmt.Printf("|  %-14s  | $  %-10.2f |\n", thirdItem, thirdItemPrice)
}

// Support functions

func itemName(reader *bufio.Reader, message string) string {
	fmt.Println(message)
	firstItem, err := reader.ReadString('\n')

	errHandler(err)
	firstItem = strings.TrimSpace(firstItem)

	return firstItem
}

func itemPrice(r *bufio.Reader, message string) (float64, error) {
	var priceStr string

	fmt.Println(message)

	priceStr, err := r.ReadString('\n')

	errHandler(err)

	firstItemPrice, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

	return firstItemPrice, err
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}
