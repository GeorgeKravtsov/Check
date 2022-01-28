package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	//	maxCardNumber, maxNumberOfItems, maxItemId,
	//		oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount := getUserInput()
	printReceipt(10000, 10, 100, 10, 100, 10) //maxCardNumber, maxNumberOfItems, maxItemId,
	//oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount
}

func getUserInput() (int, int, int, int, int, int) {
	var input string
	fmt.Print("Enter maximum discount card number (10000 for example): ")
	fmt.Scanf("%s", &input)
	maxCardNumber, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Maximum discount number is:", maxCardNumber)
	}
	fmt.Print("Enter maximum number of items (10 for example): ")
	fmt.Scanf("%s", &input)
	maxNumberOfItems, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Maximum number of items is:", maxNumberOfItems)
	}
	fmt.Print("Enter maximum item ID (100 for example): ")
	fmt.Scanf("%s", &input)
	maxItemId, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Maximum item ID is:", maxItemId)
	}
	fmt.Print("Enter maximum quantity of one item (10 for example): ")
	fmt.Scanf("%s", &input)
	oneItemMaxQuantity, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Maximum one item quantity is:", oneItemMaxQuantity)
	}
	fmt.Print("Enter maximum price of one item: (100 for example) ")
	fmt.Scanf("%s", &input)
	oneItemMaxPrice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("One item maximum price is:", oneItemMaxPrice)
	}
	fmt.Print("Enter promotion discount (10 for exmple): ")
	fmt.Scanf("%s", &input)
	promotionDiscount, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promotion discount is:", promotionDiscount)
	}
	return maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount
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

func numberOfItemsInReceipt(number int) int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(number)
}

func priceGeneration(maxOneItemPrice int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(maxOneItemPrice) * rand.Float64()
}

func promotion(promotionDiscount int) float64 {
	return float64(promotionDiscount) / float64(100)
}



func printReceipt(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount int) {
	cardNumber := cardNumberGeneration(maxCardNumber)
	discount := discountGeneration(cardNumber)

	fmt.Println("________________________________________________________")
	fmt.Printf("|%2s | %10s | %s| %s| %s| %s|\n", "№", "ItemId", "Quantity", "Price", "Total Price", "Promotion Discount")
	fmt.Println("________________________________________________________")
	reverseCounter := numberOfItemsInReceipt(maxNumberOfItems)
	itemCounter := 1
	var total float64
	var toBePaid float64
	for i := reverseCounter; i > 0; i-- {
		id := itemIdGeneration(maxItemId)
		quantity := itemQuantityGeneration(oneItemMaxQuantity)
		oneItemPrice := priceGeneration(oneItemMaxPrice)
		oneItemTotal := float64(quantity) * oneItemPrice
		switch itemCounter != 0 {
		case itemCounter == 1:
			oneItemTotal = oneItemPrice * float64(quantity)
			total += oneItemTotal
			toBePaid = total - (total * discount)
			fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f|\n",
				itemCounter, id, quantity, oneItemPrice, oneItemTotal)
			itemCounter++
		case itemCounter != 1 && itemCounter%2 != 0:
			total += oneItemTotal - oneItemTotal*promotion(promotionDiscount)
			fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f| %5.2f|\n",
				itemCounter, id, quantity, oneItemPrice, oneItemTotal,
				oneItemTotal*promotion(promotionDiscount))
			toBePaid = total - (total * discount)
			itemCounter++
			continue
		default:
			fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f|\n", itemCounter, id, quantity,
				oneItemPrice, oneItemTotal)
			total += oneItemTotal
			toBePaid = total - (total * discount)
			itemCounter++
		}
	}
	fmt.Println("________________________________________________________")
	fmt.Printf("|Total: %.2f %32s\n", total, "|")
	fmt.Println("________________________________________________________")
	fmt.Printf("|Discount card: %d %27s\n", cardNumber, "|")
	fmt.Println("________________________________________________________")
	fmt.Printf("|Discount: %.0f%s %34s\n", discount*100, "%", "|")
	fmt.Println("________________________________________________________")
	fmt.Printf("|To be paid: %.2f; Saved: %.2f %14s\n", toBePaid, total-toBePaid, "|")
	fmt.Println("________________________________________________________")
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
