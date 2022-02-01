package main

import (
	"encoding/json"
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

	
	//printReceipt(newReceipt(10000, 10, 100, 10, 100, 10)) //maxCardNumber, maxNumberOfItems, 
					//maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount
	printJsonReceipt(10000, 10, 100, 10, 100, 10)

}

func getUserInput(message1, message2 string) int {
	var input string
	fmt.Print(message1)
	fmt.Scanf("%s", &input)
	integer, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message2, integer)
	}
	return integer
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
	Number       int
	ItemId       int
	Quantity     int
	Price        float64
	NumberIsOdd  bool
	OneItemTotal float64
}

func newOneItemLine(number, maxItemId, oneItemMaxQuantity, oneItemMaxPrice int) oneItemLine {
	numberIsOdd := isOddNumber(number)
	quantity := itemQuantityGeneration(oneItemMaxQuantity)
	price := priceGeneration(oneItemMaxPrice)
	oneItemTotal := float64(quantity) * price
	return oneItemLine{Number: number, ItemId: itemIdGeneration(maxItemId), Quantity: quantity, Price: price, NumberIsOdd: numberIsOdd, OneItemTotal: oneItemTotal}
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
	CardNumber   int
	Discount     float64
	PromDiscount float64
	SliceOfLines []oneItemLine
}

func newReceipt(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount int) receipt {
	cardNumber := cardNumberGeneration(maxCardNumber)
	discount := discountGeneration(cardNumber)
	return receipt{CardNumber: cardNumber, Discount: discount, PromDiscount: promotion(promotionDiscount), SliceOfLines: getSliceOfLines(maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice)}
}

func printJsonReceipt(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount int) {
	rec := newReceipt(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount)
	jsonRec, err := json.MarshalIndent(rec, "", "  ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", jsonRec)
	}
}

func printReceipt(rec receipt) {
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|%2s | %11s| %s| %s| %s| %s|\n",
		"№", "ItemId", "Quantity", "Price", "Total Price", "Promotion Discount")
	fmt.Println("___________________________________________________________________")
	discount := rec.PromDiscount
	var total float64
	for _, line := range rec.SliceOfLines {
		if line.Number != 0 {
			if line.NumberIsOdd {
				fmt.Printf("|%2d | %11d| %8d| %5.2f| %11.2f| %18.2f|\n",
					line.Number, line.ItemId, line.Quantity, line.Price,
					line.OneItemTotal, line.OneItemTotal*discount)
				total += line.OneItemTotal - line.OneItemTotal*discount
				continue
			} else {
				fmt.Printf("|%2d | %11d| %8d| %5.2f| %11.2f| %19s\n",
					line.Number, line.ItemId, line.Quantity, line.Price,
					line.OneItemTotal, "|")
				total += line.OneItemTotal
			}
		}
	}
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|Total: %.2f %52s\n", total, "|")
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|Discount card: %d %47s\n", rec.CardNumber, "|")
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|Discount: %.0f%s %54s\n", rec.Discount*100, "%", "|")
	fmt.Println("___________________________________________________________________")
	fmt.Printf("|To be paid: %.2f; Saved: %.2f %32s\n",
		total-total*rec.Discount, total-(total-total*rec.Discount), "|")
	fmt.Println("___________________________________________________________________")
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
