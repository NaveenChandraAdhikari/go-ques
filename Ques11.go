package main
import {

	"fmt",
"encoding/json"
}
/*
Define a BankAccount struct; add methods:


Deposit(amount float64) error


Withdraw(amount float64) error (error if insufficient funds)


Balance() float64


Add JSON tags; marshal to JSON and print

 */

/*
first lers define a struct firsts
 */
type BankAccount struct {
	Owner string `json:"owner"`
	Balance float64 `json:"balance"`
}


/*\
add methods with error handling
 */
//deposit method with error handling
func(a *BankAccount) Deposit(amount float64) error {

	if amount <= 0 {
		return error.New("amount must be positive")
	}
	a.Balance += amount
	return nil

}

// withdrawl method with error for insufficienet funds

func (a *BankAccount) Withdraw(amount float64) error {

	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient balance")
	}
	a.Balance -=amount
	return nil
}
// Balance method (returns balance)
func (a *BankAccount) GetBalance() float64 {
	return a.Balance
}



func main() {
	// Create and use an account
	acc := BankAccount{Owner: "nav"}

	// Deposit money
	if err := acc.Deposit(200); err != nil {
		fmt.Println("Deposit error:", err)
	}

	// Attempt a withdrawal
	if err := acc.Withdraw(50); err != nil {
		fmt.Println("Withdraw error:", err)
	}

	// Print current balance
	fmt.Printf("Current balance for %s: %.2f\n", acc.Owner, acc.GetBalance())

	// Marshal to JSON and print
	jsonData, err := json.MarshalIndent(acc, "", "  ")
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}
	fmt.Println("Account JSON:\n" + string(jsonData))
}

