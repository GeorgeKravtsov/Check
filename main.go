package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	//	maxCardNumber := getUserInput("Enter maximum discount card number (10000 for example) ",
	//					"Maximum discount number is:")
	//	maxNumberOfItems := getUserInput("Enter maximum number of items (10 for example): ",
	//					"Maximum number of items is:")
	//	maxItemId := getUserInput("Enter maximum item ID (100 for example): ", "Maximum item ID is:")
	//	oneItemMaxQuantity := getUserInput("Enter maximum quantity of one item (10 for example): ",
	//					"Maximum one item quantity is:")
	//	oneItemMaxPrice := getUserInput("Enter maximum price of one item: (100 for example) ",
	//					"One item maximum price is:")
	//	promotionDiscount := getUserInput("Enter promotion discount (10 for exmple): ",
	//					"Promotion discount is:")

	printReceiptPrototype(10000, 10, 100, 10, 100, 10) //maxCardNumber, maxNumberOfItems,
	//maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount
	//	printReceipt(newReceipt(10000, 10, 100, 10, 100, 10))

}

func getUserInput(message1, message2 string) int {
	var input string
	fmt.Print(message1)
	fmt.Scanf("%s", &input)
	returnValue, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message2, returnValue)
	}
	return returnValue
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

func isOddNumber(number int) bool {
	if number%2 != 0 {
		return true
	}
	return false
}

type oneItemLine struct {
	number       int
	itemId       int
	quantity     int
	price        float64
	numberIsOdd  bool
	oneItemTotal float64
}

func newOneItemLine(number, maxItemId, oneItemMaxQuantity, oneItemMaxPrice int) oneItemLine {
	numberIsOdd := isOddNumber(number)
	quantity := itemQuantityGeneration(oneItemMaxQuantity)
	price := priceGeneration(oneItemMaxPrice)
	oneItemTotal := float64(quantity) * price
	return oneItemLine{number: number, itemId: itemIdGeneration(maxItemId), quantity: quantity, price: price, numberIsOdd: numberIsOdd, oneItemTotal: oneItemTotal}
}

func getSliceOfLines(maxItemNumber, maxItemId, oneItemMaxQuantity, oneItemMaxPrice int) []oneItemLine {
	numberOfItems := numberOfItemsInReceipt(maxItemNumber)
	sliceOfLines := make([]oneItemLine, numberOfItems)
	for number := 1; number <= numberOfItems; number++ {
		sliceOfLines = append(sliceOfLines, newOneItemLine(number, maxItemId, oneItemMaxQuantity, oneItemMaxPrice))
	}
	return sliceOfLines
}

type receipt struct {
	cardNumber   int
	discount     float64
	sliceOfLines []oneItemLine
	promDiscount float64
}

func newReceipt(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount int) receipt {
	cardNumber := cardNumberGeneration(maxCardNumber)
	discount := discountGeneration(cardNumber)
	return receipt{cardNumber: cardNumber, discount: discount, sliceOfLines: getSliceOfLines(maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice), promDiscount: promotion(promotionDiscount)}
}

func printReceipt(rec receipt) {
	for _, line := range rec.sliceOfLines {
		if line.number != 0 {
			fmt.Println(line)
		}
	}
	fmt.Println(rec.cardNumber)
	fmt.Println(rec.discount)
	fmt.Println(rec.promDiscount)
}

func printReceiptPrototype(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount int) {
	cardNumber := cardNumberGeneration(maxCardNumber)
	discount := discountGeneration(cardNumber)

	fmt.Println("________________________________________________________")
	fmt.Printf("|%2s | %10s | %s| %s| %s| %s|\n", "№", "ItemId", "Quantity", "Price", "Total Price", "Promotion Discount")
	fmt.Println("________________________________________________________")
	numberOfItems := numberOfItemsInReceipt(maxNumberOfItems)
	var total float64
	var toBePaid float64
	for number := 1; number <= numberOfItems; number++ {
		id := itemIdGeneration(maxItemId)
		quantity := itemQuantityGeneration(oneItemMaxQuantity)
		oneItemPrice := priceGeneration(oneItemMaxPrice)
		oneItemTotal := float64(quantity) * oneItemPrice
		switch number != 0 {
		case number == 1:
			oneItemTotal = oneItemPrice * float64(quantity)
			total += oneItemTotal
			toBePaid = total - (total * discount)
			fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f|\n",
				number, id, quantity, oneItemPrice, oneItemTotal)
		case number != 1 && number%2 != 0:
			total += oneItemTotal - oneItemTotal*promotion(promotionDiscount)
			fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f| %5.2f|\n",
				number, id, quantity, oneItemPrice, oneItemTotal,
				oneItemTotal*promotion(promotionDiscount))
			toBePaid = total - (total * discount)
			continue
		default:
			fmt.Printf("|%2d | %10d | %8d| %5.2f| %11.2f|\n", number, id, quantity,
				oneItemPrice, oneItemTotal)
			total += oneItemTotal
			toBePaid = total - (total * discount)
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
