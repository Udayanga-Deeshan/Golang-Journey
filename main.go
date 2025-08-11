package main

import (
	"bufio"
	"demo-app/greetings"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var (
	orderIDArray     []string
	phoneNumberArray []string
	sizeArray        []string
	qtyArray         []int
	amountArray      []float64
	statusArray      []int
)

const (
	XS  = 600
	S   = 800
	M   = 900
	L   = 1000
	XL  = 1100
	XXL = 1200

	PROCESSING int = 0
	DELIVERING int = 1
	DELIVERED  int = 2
)

var input = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Hello world")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	homePage()
}

func init() {
	orderIDArray = []string{}
}

func homePage() {
	fmt.Println("================== FASHION SHOP ========================")
	fmt.Println("\n[01] Place Order")
	fmt.Println("\n[02] Search Customer")
	fmt.Println("\n[03] Search Order")
	fmt.Println("\n[04] View Reports")
	fmt.Println("\n[05] Set Order Status")
	fmt.Println("\n[06] Delete Order")
	fmt.Print("\nEnter option : ")

	input.Scan()
	option := input.Text()

	clearConsole()

	switch option {
	case "1":
		placeOrder()

	case "2":
		searchCustomer()

	case "3":
		searchOrder()

	case "4":
		viewReports()

	case "5":
		setOrderStatus()

	case "6":
		deleteOrderStatus()

	default:
		fmt.Println("Invalid Option")
		homePage()

	}
}

func clearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func generateId() string {
	if len(orderIDArray) == 0 {
		return "ODR#0001"
	}

	lastOrderId := orderIDArray[len(orderIDArray)-1]

	parts := strings.Split(lastOrderId, "#")

	if len(parts) != 2 {
		return "ODR#00001"
	}

	idNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return "ODR#00001"
	}

	idNum++
	return fmt.Sprintf("ODR#%05d", idNum)

}

func validatePhoneNumber(phoneNumber string) bool {
	return len(phoneNumber) == 10 && phoneNumber[0] == '0'
}

func validateQty(qty int) bool {
	return qty > 0
}

func placeOrder() {
	fmt.Println("Welcome to the Place Order Section")
	id := generateId()
	fmt.Println("Order ID:", id)

	var phoneNumber string
phoneNumberInput:
	for {
		fmt.Print("\nEnter your Phone Number: ")
		_, err := fmt.Scan(&phoneNumber)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if validatePhoneNumber(phoneNumber) {
			fmt.Println("Your phone number is:", phoneNumber)
			break phoneNumberInput
		}

		fmt.Println("Phone number is not valid")
		fmt.Print("Do you want to enter phone number again (y/n)? ")

		var choice string
		fmt.Scan(&choice)
		choice = strings.ToLower(choice)

		if choice == "n" {
			clearConsole()
			homePage()
			return
		}
	}

	var shirtSize string
	fmt.Print("\nEnter the T-Shirt Size (XS, S, M, L, XL, XXL): ")
	fmt.Scan(&shirtSize)
	shirtSize = strings.ToUpper(shirtSize)
	fmt.Printf("You selected %s T-Shirt size\n", shirtSize)

	var qty int
quantityInput:
	for {
		fmt.Print("\nEnter the required Quantity: ")
		_, err := fmt.Scan(&qty)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		if validateQty(qty) {
			break quantityInput
		}
		fmt.Println("Quantity must be greater than 0")
	}

	amount := 0.0
	switch shirtSize {
	case "XS":
		amount = XS * float64(qty)
	case "S":
		amount = S * float64(qty)
	case "M":
		amount = M * float64(qty)
	case "L":
		amount = L * float64(qty)
	case "XL":
		amount = XL * float64(qty)
	case "XXL":
		amount = XXL * float64(qty)
	default:
		fmt.Println("Invalid size selected")
		return
	}

	fmt.Printf("\nAmount: $%.2f\n", amount)

	var confirm string
	fmt.Print("\nDo you want to place this order (y/n)? ")
	fmt.Scan(&confirm)
	confirm = strings.ToLower(confirm)

	if confirm == "y" {
		orderIDArray = append(orderIDArray, id)
		phoneNumberArray = append(phoneNumberArray, phoneNumber)
		sizeArray = append(sizeArray, shirtSize)
		qtyArray = append(qtyArray, qty)
		amountArray = append(amountArray, amount)
		statusArray = append(statusArray, 0)
		fmt.Println("\nOrder Placed Successfully!")
	}

	var another string
	fmt.Print("\nDo you want to place another order (y/n)? ")
	fmt.Scan(&another)
	another = strings.ToLower(another)

	clearConsole()
	if another == "n" {
		homePage()
	}
}

func searchCustomer() {
	fmt.Println("welcome to the search Customer Section")
}

func searchOrder() {

}

func viewReports() {

}

func setOrderStatus() {

}

func deleteOrderStatus() {

}
