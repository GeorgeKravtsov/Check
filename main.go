package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	checkPrint(10000, 10, 100, 10, 100)
}

func cardNumberGeneration(maxNumber int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNumber)
}

func itemIdGeneration(maxId int) int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(maxId)
}

func itemQuantityGeneration(maxQuantity int) int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(maxQuantity)
}

func numberOfItemsInCheck(number int) int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(number)
}

func priceGeneration(maxOneItemPrice int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(maxOneItemPrice) * rand.Float64()
}

func checkPrint(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice int) {
	cardNumber := cardNumberGeneration(maxCardNumber)
	discount := discountGeneration(cardNumber)

	fmt.Println("_______________________________________________")
	fmt.Printf("|%2s | %10s | %s| %s| %s|\n", "№", "ItemId", "Quantity", "Price", "Total Price")
	fmt.Println("_______________________________________________")
	reverseCounter := numberOfItemsInCheck(maxNumberOfItems)
	itemCounter := 1
	var total float64
	var toBePaid float64
	for {
		id := itemIdGeneration(maxItemId)
		quantity := itemQuantityGeneration(oneItemMaxQuantity)
		oneItemPrice := priceGeneration(oneItemMaxPrice)
		oneItemTotal := float64(quantity) * oneItemPrice
		fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f|\n", itemCounter, id, quantity, oneItemPrice, oneItemTotal)
		total += oneItemTotal
		toBePaid = total - (total * discount)
		if reverseCounter <= 0 {
			break
		}

		reverseCounter--
		itemCounter++

	}
	fmt.Println("_______________________________________________")
	fmt.Printf("|Total: %.2f %32s\n", total, "|")
	fmt.Println("_______________________________________________")
	fmt.Printf("|Discount card: %d %27s\n", cardNumber, "|")
	fmt.Println("_______________________________________________")
	fmt.Printf("|Discount: %.0f%s %34s\n", discount*100, "%", "|")
	fmt.Println("_______________________________________________")
	fmt.Printf("|To be paid: %.2f; Saved: %.2f %14s\n", toBePaid, total-toBePaid, "|")
	fmt.Println("_______________________________________________")
}

func discountGeneration(cardNumber int) float64 {
	stringCardNumber := fmt.Sprintf("%d", cardNumber)
	numberOfSevens := strings.Count(stringCardNumber, "7")
	switch numberOfSevens {
	case 1:
		return 0.07
	case 2:
		return 0.17
	case 3:
		return 0.37
	case 4:
		return 0.7
	default:
		return 0.03
	}
}
