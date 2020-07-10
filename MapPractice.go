//Darius Fiallo
//June 16th, 2020
//Map Practice

package main

import "fmt"

func main() {
	mapPractice := map[int]string{
		1: "Fries",
		2: "Cornbread",
		3: "Onion Rings",
		4: "Broccoli",
	}

	fmt.Println("Map value: ", mapPractice)

	//var m map[string]int
	//m["Hi"] = 1 Cannot add values to map, they're reference types
	//fmt.Println("Test map: ", m)

	//But if you want to add values, to get past this use the make function
	m := make(map[string]int)
	fmt.Println("Test map: ", m)
	m["Hi"] = 20
	m["How"] = 245
	m["test"] = 1
	m["Shelby"] = 2
	m["Darius"] = 3
	fmt.Println("Updated map: ", m)
	fmt.Println("Length of map: ", len(m)) //Shows number of keys

	//Retrieving a value based on it's key

	var num int = m["Hi"] //case sensitive, also, if grab value that isn't there, returns 0
	fmt.Println("Value of num: ", num)

	for key, value := range m {
		fmt.Println("Key: ", key, " Value: ", value)
	}

	//Checking if a key exists. Can replace 1st variable (a) with _ if
	//don't want to store and rather just check if it exists
	a, b := m["Hi"]
	fmt.Println("Value of a: ", a)
	fmt.Println("Value of b: ", b)

	//Deleting a key value pair

	delete(m, "Hi")
	fmt.Println("Updated map: ", m)

}
