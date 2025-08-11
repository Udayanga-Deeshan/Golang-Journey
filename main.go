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
	size             string
	qty              int
	amount           float64
	status           int
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

func placeOrder() {
	fmt.Println("welcome to the place order Section")
	id := generateId()
	fmt.Println("Order ID", id)

phoneNumberInput:
	for {
		fmt.Println("\n Enter your Phone Number")
		input.Scan()
		phoneNumber := input.Text()
		isValidPhoneNumber := validatePhoneNumber(phoneNumber)
		if isValidPhoneNumber {
			fmt.Println("Your phone Number is ", phoneNumber)
			break phoneNumberInput
		}
		fmt.Println("Phone numer is not Valid")
		fmt.Print("\nDo you want to enter phone number again (y/n): ")
		input.Scan()
		ch := input.Text()[0]
		if ch == 'Y' || ch == 'y' {
			continue
		} else if ch == 'N' || ch == 'n' {
			clearConsole()
			homePage()
		}

	}
	var shirtSize string
	fmt.Println("Enter the T Shirt Size")
	fmt.Scan(&shirtSize)
	fmt.Printf("you selected %v  T Shirt-Size\n", shirtSize)

	var qty int
	fmt.Println("Enter the required Quantity")
	fmt.Scan(&qty)
	fmt.Printf("You got %v items in that Size", qty)

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
