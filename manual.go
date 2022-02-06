package main

import (
	"fmt"
)

type oneItemLine struct {
	Number       int
	ItemId       int
	Quantity     int
	Price        float64
	OneItemTotal float64
	NumberIsOdd  bool
	
}

func newOneItemLine(number, itemId, oneItemQuantity int, oneItemPrice float64) oneItemLine {
	oneItemTotal := float64(oneItemQuantity) * oneItemPrice
	return oneItemLine{Number: number, ItemId: itemId, Quantity: oneItemQuantity, Price: oneItemPrice, OneItemTotal: oneItemTotal, NumberIsOdd: isOddNumber(number)}
}

func getSliceOfLines(itemNumber, itemId, oneItemQuantity int, oneItemPrice float64) []oneItemLine {
	sliceOfLines := make([]oneItemLine, itemNumber)
	for number := 1; number <= itemNumber; number++ {
		sliceOfLines = append(sliceOfLines, newOneItemLine(number, itemId, oneItemQuantity, oneItemPrice))
	}
	return sliceOfLines
}

type receipt struct {
	CardNumber   int
	Discount     float64
	PromDiscount float64
	SliceOfLines []oneItemLine
}

func (rec receipt) total() float64 {
	var total float64
	for _, line := range rec.SliceOfLines {
		if line.Number != 0 {
			if line.NumberIsOdd {
				total += line.OneItemTotal - line.OneItemTotal*rec.PromDiscount
				continue
			} else {
				total += line.OneItemTotal
			}
		}
	}
	return total
}

func (rec receipt) toBePaid() float64 {
	return rec.total() - rec.total()*rec.Discount
}

func (rec receipt) saved() float64 {
	return rec.total() - rec.toBePaid()
}

func newReceipt(cardNumber, numberOfItems, itemId, itemQuantity, promotionDiscount int, oneItemPrice float64) receipt {
	discount := discountGeneration(cardNumber)
	return receipt{CardNumber: cardNumber, Discount: discount, PromDiscount: promotion(promotionDiscount), SliceOfLines: getSliceOfLines(numberOfItems, itemId, itemQuantity, oneItemPrice)}
}

func printReceipt(rec receipt) {
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|%2s | %11s| %s| %s| %s| %s|\n",
		"№", "ItemId", "Quantity", "Price", "Total Price", "Promotion Discount")
	fmt.Println("___________________________________________________________________")
	for _, line := range rec.SliceOfLines {
		if line.Number != 0 {
			if line.NumberIsOdd {
				fmt.Printf("|%2d | %11d| %8d| %5.2f| %11.2f| %18.2f|\n",
					line.Number, line.ItemId, line.Quantity, line.Price,
					line.OneItemTotal, line.OneItemTotal*rec.PromDiscount)
				continue
			} else {
				fmt.Printf("|%2d | %11d| %8d| %5.2f| %11.2f| %19s\n",
					line.Number, line.ItemId, line.Quantity, line.Price,
					line.OneItemTotal, "|")
			}
		}
	}
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|Total: %.2f %52s\n", rec.total(), "|")
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|Discount card: %d %47s\n", rec.CardNumber, "|")
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|Discount: %.0f%s %54s\n", rec.Discount*100, "%", "|")
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|To be paid: %.2f; Saved: %.2f %32s\n",
		rec.toBePaid(), rec.saved(), "|")
	fmt.Println("___________________________________________________________________")
}
