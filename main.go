package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	fmt.Println("_______________________________________________")
	fmt.Printf("|Discount card: %d %27s\n", cardNumberGeneration(maxCardNumber), "|")
	fmt.Println("_______________________________________________")
	fmt.Printf("|%2s | %10s | %s| %s| %s|\n", "№", "ItemId", "Quantity", "Price", "Total Price")
	fmt.Println("_______________________________________________")
	reverseCounter := numberOfItemsInCheck(maxNumberOfItems)
	itemCounter := 1
	var total float64
	for {
		id := itemIdGeneration(maxItemId)
		quantity := itemQuantityGeneration(oneItemMaxQuantity)
		oneItemPrice := priceGeneration(oneItemMaxPrice)
		oneItemTotal := float64(quantity) * oneItemPrice
		fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f|\n", itemCounter, id, quantity, oneItemPrice, oneItemTotal)
		total += oneItemTotal
		if reverseCounter <= 0 {
			break
		}

		reverseCounter--
		itemCounter++

	}
	fmt.Println("_______________________________________________")
	fmt.Printf("|Total: %.2f %32s\n", total, "|")
	fmt.Println("_______________________________________________")
}

func main() {
	checkPrint(10000, 10, 100, 10, 100)
}
