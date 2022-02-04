package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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

//	rec := newReceiptAuto(10000, 10, 100, 10, 100, 10) //maxCardNumber, maxNumberOfItems,maxItemId						//oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount
	//	printJsonReceipt(rec)
//	rec := readJsonToRec("rec.json")
	//	recJsonToFile(rec)
//	printReceiptAuto(rec)
fmt.Println(newReceipt(100, 10, 100, 10, 100, 10.0))
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

type oneItemLineAuto struct {
	Number       int
	ItemId       int
	Quantity     int
	Price        float64
	OneItemTotal float64
	NumberIsOdd  bool
	
}

func newOneItemLineAuto(number, maxItemId, oneItemMaxQuantity, oneItemMaxPrice int) oneItemLineAuto {
	numberIsOdd := isOddNumber(number)
	quantity := itemQuantityGeneration(oneItemMaxQuantity)
	price := priceGeneration(oneItemMaxPrice)
	oneItemTotal := float64(quantity) * price
	return oneItemLineAuto{Number: number, ItemId: itemIdGeneration(maxItemId), Quantity: quantity, Price: price, OneItemTotal: oneItemTotal, NumberIsOdd: numberIsOdd}
}

func getSliceOfLinesAuto(maxItemNumber, maxItemId, oneItemMaxQuantity, oneItemMaxPrice int) []oneItemLineAuto {
	numberOfItems := numberOfItemsInReceipt(maxItemNumber)
	sliceOfLines := make([]oneItemLineAuto, numberOfItems)
	for number := 1; number <= numberOfItems; number++ {
		sliceOfLines = append(sliceOfLines, newOneItemLineAuto(number, maxItemId, oneItemMaxQuantity, oneItemMaxPrice))
	}
	return sliceOfLines
}

type receiptAuto struct {
	CardNumber   int
	Discount     float64
	PromDiscount float64
	SliceOfLines []oneItemLineAuto
}

func (rec receiptAuto) total() float64 {
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

func (rec receiptAuto) toBePaid() float64 {
	return rec.total() - rec.total()*rec.Discount
}

func (rec receiptAuto) saved() float64 {
	return rec.total() - rec.toBePaid()
}

func newReceiptAuto(maxCardNumber, maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount int) receiptAuto {
	cardNumber := cardNumberGeneration(maxCardNumber)
	discount := discountGeneration(cardNumber)
	return receiptAuto{CardNumber: cardNumber, Discount: discount, PromDiscount: promotion(promotionDiscount), SliceOfLines: getSliceOfLinesAuto(maxNumberOfItems, maxItemId, oneItemMaxQuantity, oneItemMaxPrice)}
}

func recAutoJsonToFile(rec receiptAuto) {
	jsonRec, err := json.Marshal(rec)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", jsonRec)
	}
	file, e := os.Create("rec.json")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	} else {
		defer file.Close()
		file.Write(jsonRec)
		fmt.Println("File recorded")
	}
}

func readJsonToRecAuto(filename string) receiptAuto {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	data := receiptAuto{}
	json.Unmarshal([]byte(file), &data)
	return data
}

func printJsonReceiptAuto(rec receiptAuto) {
	jsonRec, err := json.MarshalIndent(rec, "", "  ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", jsonRec)
	}
}

func printReceiptAuto(rec receiptAuto) {
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

func newReceipt(cardNumber, numberOfItems, itemId, itemQuantity, promotionDiscount int, oneItemPrice float64 ) receipt {
	discount := discountGeneration(cardNumber)
	return receipt{CardNumber: cardNumber, Discount: discount, PromDiscount: promotion(promotionDiscount), SliceOfLines: getSliceOfLines(numberOfItems, itemId, itemQuantity, oneItemPrice)}
}
