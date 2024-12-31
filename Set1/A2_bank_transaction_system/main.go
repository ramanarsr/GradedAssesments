package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID                 string
	Name               string
	Balance            float64
	TransactionHistory []string
}

var accounts []Account

const (
	DepositOption            = 1
	WithdrawOption           = 2
	ViewBalanceOption        = 3
	TransactionHistoryOption = 4
	ExitOption               = 5
)

func addAccount(id, name string, initialBalance float64) {
	accounts = append(accounts, Account{
		ID:                 id,
		Name:               name,
		Balance:            initialBalance,
		TransactionHistory: []string{"Account created with balance: $" + fmt.Sprintf("%.2f", initialBalance)},
	})
}

func deposit(accountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	for i, acc := range accounts {
		if acc.ID == accountID {
			accounts[i].Balance += amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, fmt.Sprintf("Deposited: $%.2f", amount))
			return nil
		}
	}
	return errors.New("account not found")
}

func withdraw(accountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}
	for i, acc := range accounts {
		if acc.ID == accountID {
			if acc.Balance < amount {
				return errors.New("insufficient balance")
			}
			accounts[i].Balance -= amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory, fmt.Sprintf("Withdrew: $%.2f", amount))
			return nil
		}
	}
	return errors.New("account not found")
}

func viewBalance(accountID string) (float64, error) {
	for _, acc := range accounts {
		if acc.ID == accountID {
			return acc.Balance, nil
		}
	}
	return 0, errors.New("account not found")
}

func displayTransactionHistory(accountID string) error {
	for _, acc := range accounts {
		if acc.ID == accountID {
			fmt.Println("Transaction History:")
			for _, txn := range acc.TransactionHistory {
				fmt.Println(txn)
			}
			return nil
		}
	}
	return errors.New("account not found")
}

func main() {
	addAccount("1", "Alice", 1000)
	addAccount("2", "Bob", 500)

	var choice int
	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		var accountID string
		var amount float64

		switch choice {
		case DepositOption:
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&accountID)
			fmt.Print("Enter amount to deposit: ")
			fmt.Scanln(&amount)
			if err := deposit(accountID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case WithdrawOption:
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&accountID)
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scanln(&amount)
			if err := withdraw(accountID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case ViewBalanceOption:
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&accountID)
			if balance, err := viewBalance(accountID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Current Balance: $%.2f\n", balance)
			}

		case TransactionHistoryOption:
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&accountID)
			if err := displayTransactionHistory(accountID); err != nil {
				fmt.Println("Error:", err)
			}

		case ExitOption:
			fmt.Println("Exiting system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
