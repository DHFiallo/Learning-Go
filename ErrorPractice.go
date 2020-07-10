package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type account struct {
	number  string
	balance float64
}

type withdrawError struct {
	err     string
	value   float64
	balance float64
}

func (a *account) withdraw(value float64) (bool, error) {

	if a.balance <= 0 {
		return false, &withdrawError{"Withdrawal Error", value, a.balance}
	}
	if value <= 0 {
		return false, &withdrawError{"Withdrawal Error", value, a.balance}
	}
	if a.balance >= value {
		a.balance = a.balance - value
		return true, nil
	}
	return false, &withdrawError{"Withdrawal Error", value, a.balance}
	// use the errors package to display a new, custom error message
	//return false, errors.New("You cannot withdraw from this account.")
	//return false, fmt.Errorf("Withdrawal failed, because the requested amount %0.2f is higher than balance %0.2f.", value, a.balance)
}

func (e *withdrawError) Error() string {
	return fmt.Sprintf("%s: withdrawal failed because the requested amount %0.2f is higher than balance: %0.2f", e.err, e.value, e.balance)
}

func (e *withdrawError) balanceNegativeorZero() bool {
	return e.balance <= 0
}

func (e *withdrawError) AmountNegativeorZero() bool {
	return e.value <= 0
}

func (e *withdrawError) InsufficientFunds() bool {
	return e.balance-e.value < 0
}

func main() {
	//var s string = "10"
	var s string = "10x"

	// the ParseInt function returns the parsed integer or
	// the error if the conversion failed
	i, error := strconv.ParseInt(s, 10, 8)
	fmt.Println(i)
	fmt.Println(error)
	fmt.Println(reflect.TypeOf(error))

	if error == nil {
		fmt.Println(i)
	} else {
		fmt.Println("You cannot convert text into a number.")
		fmt.Println(error)
	}

	a := account{}
	a.number = "C21345345345355"
	//a.balance = 159
	a.balance = -100

	//	out, err := a.withdraw(200)
	// fmt.Println(out)

	_, err := a.withdraw(200)
	// output if error

	if err != nil {
		if err, ok := err.(*withdrawError); ok {
			if err.AmountNegativeorZero() {
				fmt.Println("Amount to be withdrawn is negative")
			}
			if err.balanceNegativeorZero() {
				fmt.Println("Balance is negative")
			}
			if err.InsufficientFunds() {
				fmt.Println("Insufficient funds")
			}
			return
		}
	}

	// output if no error
	fmt.Println("The withdrawal occurred succcessfully")
	fmt.Println("Your new balance is", a)

}
