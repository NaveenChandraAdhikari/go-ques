package main

/*
Modify BankAccount so methods use pointer receivers


Write a function CloneAccount(a *BankAccount) *BankAccount to deep copy it

 */

/*
What is a Deep Copy?
Makes a brand new copy of all the data, so the original and the clone don’t share any internal data: changing one does not affect the other.

In Go, a deep copy means creating a new struct with the same values, but as a different memory object.
 */
/*
What is a Shallow Copy?
Copies only the top-level structure.

If your struct contains references (like pointers, slices, maps), both the original and the copy may still point to the same underlying data—so changing one can affect the other.

For simple structs (only value types), shallow and deep copy look the same. For more complex/nested structs, the difference matters.
 */
/*
1. Basic: Shallow Copy Example
Here, a shallow copy of a struct is created by simple assignment:

go
type BankAccount struct {
    Owner   string
    Balance float64
}

func main() {
    acc1 := BankAccount{Owner: "Alice", Balance: 100}
    acc2 := acc1 // This is a shallow copy for value types
    acc2.Owner = "Bob" // Changing acc2 does NOT affect acc1

    fmt.Println(acc1) // {Alice 100}
    fmt.Println(acc2) // {Bob 100}
}
For structs with only value types (string, float64), shallow copy is usually safe.

2. Adding a Pointer: Showing Shallow Copy Pitfall
If your struct contains a slice or another pointer, shallow copy can be dangerous:

go
type BankAccount struct {
    Owner    string
    Balance  float64
    History  []string
}

func main() {
    acc1 := BankAccount{Owner: "Alice", Balance: 100, History: []string{"Init"}}
    acc2 := acc1 // shallow copy
    acc2.History = append(acc2.History, "Deposit 50")
    fmt.Println(acc1.History) // Both acc1 and acc2 now show: ["Init" "Deposit 50"]
}
Slices are pointers underneath; shallow copy means shared slice data.




When to Use Which?
Use a shallow copy for simple structs with only value types.

For structs with slices, maps, or pointers, always use a deep copy if you want independent copies.
*/


import (
"encoding/json"
"errors"
"fmt"
)

type BankAccount struct {
	Owner   string   `json:"owner"`
	Balance float64  `json:"balance"`
	Notes   []string `json:"notes"`
}

func (a *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	a.Balance += amount
	a.Notes = append(a.Notes, fmt.Sprintf("Deposited %.2f", amount))
	return nil
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	a.Notes = append(a.Notes, fmt.Sprintf("Withdrew %.2f", amount))
	return nil
}

func (a *BankAccount) BalanceAmount() float64 {
	return a.Balance
}

// deep copy (CloneAccount) using make + copy for slice ,,we can use for append but it reverses the order and veru inefficent way 
func CloneAccount(a *BankAccount) *BankAccount {
	if a == nil {
		return nil
	}
	notesCopy := make([]string, len(a.Notes))
	copy(notesCopy, a.Notes)
	return &BankAccount{
		Owner:   a.Owner,
		Balance: a.Balance,
		Notes:   notesCopy, // new slice 
	}
}

func main() {
	acc := &BankAccount{Owner: "Nav", Notes: []string{"Account opened"}}
	acc.Deposit(100)
	acc.Withdraw(30)

	// clonet the  account
	accClone := CloneAccount(acc)
	accClone.Owner = "Kav"
	accClone.Deposit(50)
	accClone.Notes = append(accClone.Notes, "Promo credit")

	// JSON print both
	orig, _ := json.MarshalIndent(acc, "", "  ")
	clone, _ := json.MarshalIndent(accClone, "", "  ")
	fmt.Println("Original Account JSON:\n", string(orig))
	fmt.Println("\nCloned Account JSON:\n", string(clone))
}




















/*


in the ques ::


With only Owner (string) and Balance (float64), both are simple value types.

Shallow and deep copy make no practical difference for value types—copying is always safe and independent.

By adding a slice field (Notes), the struct includes an internal reference type.

With slices, shallow copy and deep copy are different!

Shallow copy would cause both structs to share the same slice data.

Deep copy makes a new slice, ensuring real independence of the internal note history.

ALSO
 To Make Effects of Cloning Clear
When you change the Notes field in one account, with a deep copy, the other account’s notes do not change—proving true independence!

Without any reference types in the struct, it's impossible to visually demonstrate the benefit of a deep vs. shallow copy.


Struct fields	Need deep copy?	Example use
Only value types	No	Assignment is fine
Slices/maps/pointers	Yes (for safety)	Use make+copy


5. Conclusion Table
Method	Deep Copy?	Order Preserved?	Idiomatic and Efficient?
a := original (struct assignment)	SHALLOW*	Yes	Only if no slices/maps
Manual loop append([]string{i}, ...)	Deep of slice	No (reverses)	No (non-idiomatic)
make+copy	Deep of slice	Yes	Yes
append(dc.checkbook, original...)	Deep of slice	Yes	Also yes (if nil slice)
* shallow = slices/maps are NOT deep copied

6. Short Takeaway
Always use make+copy or append(..., ... ) for slices if you want a true copy.

Struct assignment is shallow: slices/maps/channels inside are still linked.

Your custom for loop works but reverses elements and is not idiomatic.


play with example:)

package main

import (
	"fmt"
)

func main() {

	type BankAccount struct {
		Owner     string
		Balance   float64
		checkbook []string
	}

	//original account
	original := BankAccount{
		Owner:     "NAV",
		Balance:   1000.50,
		checkbook: []string{"Deposit 500", "Withdraw 100"},
	}

	var dc BankAccount
	dc.Owner = original.Owner
	dc.Balance = original.Balance
	for _, i := range original.checkbook {
		dc.checkbook = append([]string{i}, dc.checkbook...)
	}
	//a := original
	//dc.checkbook[0] = "Mod"
	//fmt.Println(a)
	//fmt.Println(original)

	//fmt.Println("length", len(dc.checkbook))
	//fmt.Println(dc.Owner, dc.Balance, dc.checkbook)
	//n := copy(dc.checkbook, original.checkbook)
	//fmt.Println("Number of elements co:", n)
}

MORE EXAMPLE:
package main

import (
    "fmt"
)

// BankAccount struct with Notes field for demonstration
type BankAccount struct {
    Owner   string
    Balance float64
    Notes   []string
}

func main() {
    original := BankAccount{
        Owner:   "Alice",
        Balance: 1000,
        Notes:   []string{"Opened account", "Deposited 500"},
    }

    // --------------------------
    // SHALLOW COPY EXAMPLE
    // --------------------------
    shallowCopy := original // This copies the struct fields
    shallowCopy.Owner = "Bob"
    shallowCopy.Notes[0] = "Changed by shallowCopy!" // This will affect original too!

    fmt.Println("After shallow copy and modification:")
    fmt.Println("Original:", original)      // Notes[0] is changed!
    fmt.Println("Shallow:", shallowCopy)    // Same Notes content

    // --------------------------
    // DEEP COPY EXAMPLE
    // --------------------------
    // Re-initialize original for clear separation
    original = BankAccount{
        Owner:   "Alice",
        Balance: 1000,
        Notes:   []string{"Opened account", "Deposited 500"},
    }

    // Step 1: Make a new slice and copy contents
    notesCopy := make([]string, len(original.Notes))
    copy(notesCopy, original.Notes)

    // Step 2: Assign fields and use the new slice for deep copy
    deepCopy := original
    deepCopy.Owner = "Charlie"
    deepCopy.Notes = notesCopy // Prevents sharing with original

    // Now modify ONLY the deep copy's Notes
    deepCopy.Notes[0] = "Changed by deepCopy!"

    fmt.Println("\nAfter deep copy and modification:")
    fmt.Println("Original:", original)    // Notes[0] NOT changed
    fmt.Println("DeepCopy:", deepCopy)    // Only deepCopy's Notes[0] changed
}


*/
