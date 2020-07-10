//Darius Fiallo
//June 16, 2020
//Interfaces And Methods

package main

import "fmt"

/*type AccountOperations interface {
	// methods
	withdraw(value float64) bool
	deposit(value float64) bool
	displayInfo()
}*/

type AccountOperations interface {
	// Methods
	computeInterest() float64
}
type account struct {
	number  string
	balance float64
}

type SavingsAccount struct {
	number   string
	balance  float64
	interest float64
}

type CheckingAccount struct {
	number   string
	balance  float64
	interest float64
}

func (a *SavingsAccount) computeInterest() float64 {
	return 0.001
}

func (a *CheckingAccount) computeInterest() float64 {
	return 0.005
}

func describe(ao AccountOperations) {
	// we use %T to display the dynamic type of ao
	// and %v to display the dynamic value
	fmt.Printf("Interface type %T value %v\n", ao, ao)
}

func (a *account) withdraw(value float64) bool {
	if a.balance >= value {
		a.balance = a.balance - value
		return true
	}
	return false
}

func (a *account) deposit(value float64) bool {
	if value > 0 {
		a.balance = a.balance + value
		return true
	}
	return false
}

func CheckType(i interface{}) {
	switch i.(type) {
	case *SavingsAccount:
		fmt.Println("This is a saving account")
	case *CheckingAccount:
		fmt.Println("This is a checking account")
	default:
		fmt.Println("Unknown account")
	}
}

func (a *account) displayInfo() {
	fmt.Println(a.balance)
	fmt.Println(a.number)
}

func main() {
	var ao AccountOperations
	fmt.Println("initial value:", ao)

	// assign ao a pointer to an account value that we created.
	//We can only do this because the account type implements all the methods of AccountOperations Interface.
	//ao = &account{"C13443533535", 1500}

	//withdrawal amount
	/*ao.withdraw(150)
	ao.displayInfo()

	// deposit amount
	ao.deposit(1000)
	ao.displayInfo()*/

	a := SavingsAccount{}
	a.number = "S21345345345355"
	a.balance = 159

	var ao1 AccountOperations
	ao1 = &a
	fmt.Println("ao1 type:")
	describe(ao1)
	fmt.Println("Result for ao1")
	CheckType(ao1)

	b := CheckingAccount{}
	b.number = "C218678678345345355"
	b.balance = 2000

	var ao2 AccountOperations
	ao2 = &b
	fmt.Println("ao2 type:")
	describe(ao2)
	fmt.Println("Result for ao2")
	CheckType(ao2)
}
