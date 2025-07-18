package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var balance int64 = 0

func loadbal() {
	data, err := os.ReadFile("bal.txt")
	if err != nil {
		fmt.Println("your balance sheet is empty")
		balance = 1000
	}
	str := strings.TrimSpace(string(data))
	bal, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("error parsing the string")
	} else {
		balance = bal
	}
}

func savebal(newbal int64) {
	err := os.WriteFile("bal.txt", []byte(fmt.Sprint(newbal)), 0644)
	if err != nil {
		fmt.Println("sorry error in saving the balance")
	}
}

func deposit(amount int64) {
	balance += amount
	fmt.Printf("%d has been deposited ", amount)
	savebal(balance)
}
func withdraw(amount int64) {
	if amount > balance {
		fmt.Println("Insufficient balance.")
	} else {
		balance -= amount
		fmt.Printf("₹%d withdrawn successfully.\n", amount)
		savebal(balance)
	}
}

func viewBalance() {
	fmt.Printf("Your current balance is ₹%d\n", balance)
}

func main() {
	var opt, amount int64
	loadbal()
	for {
		fmt.Println("welcome to the console ")
		fmt.Println("please choose the service:")
		fmt.Println("1. deposit money")
		fmt.Println("2. withraw money")
		fmt.Println("3. view balance")
		fmt.Println("4. exit")

		fmt.Scan(&opt)

		switch opt {
		case 1:
			fmt.Print("Enter amount to deposit: ₹")
			fmt.Scan(&amount)
			deposit(amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ₹")
			fmt.Scan(&amount)
			withdraw(amount)

		case 3:
			viewBalance()
		case 4:
			fmt.Println("Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid option")
		}
	}
}
