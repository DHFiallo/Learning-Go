package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		// use the ReadFile from ioutil package to read the entire file in memory
		data, err := ioutil.ReadFile("/Users/Adian/Documents/Visual Studio Code/Go/fileio/flatland01.pdf")

		// feedback message in case of error
		fmt.Println(err)

		// convert the file contents to a string and display it
		fmt.Print(string(data))
		//fmt.Print(strings.ToUpper(string(data)))
	*/

	/*//Program doesn't stop running when encountering an error, this makes it stop
	_, err = os.Create("/erroneous_path/file.txt")
	if err != nil {
		panic(err)
	}*/

	//Read only part of the file
	f, err := os.Open("C:/Users/Adian/Documents/Visual Studio Code/Go/fileio/flatland01.pdf")

	fmt.Println(err)

	//  create a slice of bytes
	b1 := make([]byte, 5)
	data, err := f.Read(b1)

	// feedback message in case of error; otherwise nil
	fmt.Println(err) // print error due to reading from the file. nil is no errors

	// display the slice
	fmt.Printf(string(b1[:data]))

	// close the file after completing the operations
	f.Close()
}
