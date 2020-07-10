//Darius Fiallo
//June 16, 2020
//Method Practice

package main

import "fmt"

//Type keyword is used to create a new type based on a data type
//example below

type s string //s is the name of this new type

//this method has a receiver of type s

func (text s) IsEmpty() bool {
	if len(text) > 0 {
		return false
	}
	return true
}

//Created a new data type because definition of receiver type must be in
//same package as the method so couldn't use string type.
//Below shows what happens

/*func (text string) IsEmpty() bool {
	if len(text) > 0 {
		return false
	}
	return true
}*/

func main() {
	text := s("Hi")

	fmt.Println("Type value: ", text)
	fmt.Println("Method value: ", text.IsEmpty())
}
