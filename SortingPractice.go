package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*create an alias type
type mytype []string

// implement the Len method
func (s mytype) Len() int {
	return len(s)
}

// implement the Swap method
func (s mytype) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}*/

func bubbleSort(array []int) {
	a := arr(array)
	for !sort.IntsAreSorted(a) {

		sort.Sort(a)
	}
}

func (a arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a arr) Len() int {
	return len(a)
}

func (a arr) Less(i, j int) bool {
	return a[i] < a[j]
}

//Mergesort
func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	arr := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}
	}
	for i < len(left) {
		arr = append(arr, left[i])
		i++
	}
	for j < len(right) {
		arr = append(arr, right[j])
		j++
	}
	return arr
}

//Quicksort
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

//Selection Sort
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		var temp int = arr[i]
		var pos int = i
		for j := i; j < len(arr); j++ {
			if arr[j] <= temp {
				temp = arr[j]
				pos = j
			}
		}
		arr[pos] = arr[i]
		arr[i] = temp
	}
}

func main() {

	//SelectionSort
	arr := []int{4353, 246, 56, 452, 345, 23, 55, 6543, 4, 5, 24}
	fmt.Println(selectionSort(arr))

	//Mergesort
	arr := []int{4353, 246, 56, 452, 345, 23, 55, 6543, 4, 5, 24}
	fmt.Println(mergeSort(arr))

	//Quicksort
	var numbers []int = []int{11, 13, 8, 23, 5, 1, 39, 9}
	fmt.Println("\n--- Unsorted --- \n\n", numbers)
	quicksort(numbers)
	fmt.Println("\n--- Sorted ---\n\n", numbers, "\n")

	//Bubblesort
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Unsorted:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)

	//Implementing sort function
	// create a slice of strings
	fruits := []string{"pear", "passionfruit", "mango", "banana", "fig"}
	fmt.Println("Original slice:", fruits)

	// create a mytype variable
	myfruits := mytype(fruits)

	sort.Sort(myfruits)
	fmt.Println("Sorted by length:", myfruits)

	//Integer sort
	// define a slice
	numbers := []int{67, 18, 62, 60, 25, 64, 75, 5, 17, 55}
	fmt.Println("Original Numbers:", numbers)

	// use the sort.Ints function to sort the values in the slice
	sort.Ints(numbers)

	fmt.Println("Sorted Numbers:", numbers)

	//Sort Strings
	words := []string{"camel", "zebra", "horse", "dog", "elephant", "giraffe"}
	fmt.Println("Original slice:", words)

	// sort the values in the slice
	sort.Strings(words)
	fmt.Println("Sorted slice:", words)

	//Check sorted values
	// define a slice
	words := []string{"camel", "zebra", "horse", "dog", "elephant", "giraffe"}
	fmt.Println("Original slice:", words)
	fmt.Println("The original values are sorted:", sort.StringsAreSorted(words))

	// sort the values in the slice
	sort.Strings(words)
	fmt.Println("Sorted slice:", words)
	fmt.Println("The values are sorted:", sort.StringsAreSorted(words))
}
