//Darius Fiallo
//June 25, 2020
//Date-Time Calculator

package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func calculatorOne() {

	fmt.Println("This is calculator 1. It can add or subtract lengths of time.")
	fmt.Println("The format for this will be DD-HH-MM-SS which is days, hours, minutes, and seconds.")
	var user int
	for {
		fmt.Println("Do you want to add or subtract? 1 for add, 2 for subtract.")
		fmt.Scanln(&user)
		if user == 1 || user == 2 {
			break
		} else {
			fmt.Println("Try again, but with a correct input.")
		}
	}

	fmt.Println("The first")
	time1 := getUserTime()

	fmt.Println("The second")
	time2 := getUserTime()

	if user == 1 {
		fmt.Println("Adding them both..")
		time3 := time.Date(0, 0, (time1.Day() + time2.Day()), (time1.Hour() + time2.Hour()), (time1.Minute() + time2.Minute()), (time1.Second() + time2.Second()), 0, time.Local)
		fmt.Println("Total:", time3.Day(), "days ", time3.Hour(), "hours ", time3.Minute(), "minutes ", time3.Second(), "seconds")
		fmt.Println("Days:", time3.Day())
		fmt.Println("Hours:", time3.Hour()+time3.Day()*24)
		fmt.Println("Minutes:", (time3.Hour()+time3.Day()*24)*60+time3.Minute())
		fmt.Println("Seconds:", ((time3.Hour()+time3.Day()*24)*60+time3.Minute())*60+time3.Second())
	} else {
		diff := time1.Sub(time2)
		fmt.Println("Subtracting them both..")
		fmt.Println("Total:", diff)
		fmt.Println("Days:", diff.Hours()/24)
		fmt.Println("Hours:", diff.Hours())
		fmt.Println("Minutes:", diff.Minutes())
		fmt.Println("Seconds:", diff.Seconds())
	}

	var choice string
	fmt.Println("Want to go again? Y/N")
	fmt.Scanln(&choice)
	if choice == "Y" {
		calculatorOne()
	}
}

func calculatorTwo() {

	fmt.Println("This is calculator 2. It can add or subtract time from a date")
	fmt.Println("The format for this will be MM-DD-YYYY-HH-mm-SS which is months, days, year, hours, minutes, and seconds in order.")
	var user int
	for {
		fmt.Println("Do you want to add or subtract? 1 for add, 2 for subtract.")
		fmt.Scanln(&user)
		if user == 1 || user == 2 {
			fmt.Println(user)
			break
		} else {
			fmt.Println("Try again, but with a correct date.")
		}
	}
	time1 := getUserDate()
	time2 := getUserTime()
	if user == 1 {
		fmt.Println("Adding them both..")
		time3 := time.Date(time1.Year(), time1.Month(), (time1.Day() + time2.Day()), (time1.Hour() + time2.Hour()), (time1.Minute() + time2.Minute()), (time1.Second() + time2.Second()), 0, time.Local)
		fmt.Println("Total:", time3)
	} else {
		fmt.Println("Subtracting them both..")
		date := time.Date(time1.Year(), time1.Month(), (time1.Day() - time2.Day()), (time1.Hour() - time2.Hour()), (time1.Minute() - time2.Minute()), (time1.Second() - time2.Second()), 0, time.Local)
		fmt.Println(date)
	}

	var choice string
	fmt.Println("Want to go again? Y/N")
	fmt.Scanln(&choice)
	if choice == "Y" {
		calculatorTwo()
	}

}

func getUserDate() time.Time {
	longDate := regexp.MustCompile("[0-9]{2}-[0-9]{2}-[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2}")

	var date1 string
	fmt.Println("What is the date? MM-DD-YYYY-HH-mm-SS")
	fmt.Scanln(&date1)
	for !longDate.MatchString(date1) {
		fmt.Println("Wrong. Try again. MM-DD-YYYY-HH-mm-SS")
		fmt.Scanln(&date1)
	}
	split1 := strings.Split(date1, "-")
	months, _ := strconv.Atoi(split1[0])   // Months
	days1, _ := strconv.Atoi(split1[1])    //Days
	year, _ := strconv.Atoi(split1[2])     //Year
	hours1, _ := strconv.Atoi(split1[3])   //Hours
	minutes1, _ := strconv.Atoi(split1[4]) //Minutes
	seconds1, _ := strconv.Atoi(split1[5]) //Seconds
	time1 := time.Date(year, time.Month(months), days1, hours1, minutes1, seconds1, 0, time.Local)
	return time1
}

func getUserTime() time.Time {
	shortTime := regexp.MustCompile("[0-9]{2}-[0-9]{2}-[0-9]{2}-[0-9]{2}")
	var date2 string
	fmt.Println("What is the time? DD-HH-MM-SS, this is days, hours, minutes, and seconds in order.")
	fmt.Scanln(&date2)
	for !shortTime.MatchString(date2) {
		fmt.Println("Wrong. Try again. DD-HH-MM-SS, this is days, hours, minutes, and seconds in order.")
		fmt.Scanln(&date2)
	}
	split2 := strings.Split(date2, "-")
	days2, _ := strconv.Atoi(split2[0])    //Days
	hours2, _ := strconv.Atoi(split2[1])   //Hours
	minutes2, _ := strconv.Atoi(split2[2]) //Minutes
	seconds2, _ := strconv.Atoi(split2[3]) //Seconds
	time2 := time.Date(0, 0, days2, hours2, minutes2, seconds2, 0, time.Local)
	return time2

}

func ageCalculator() {
	fmt.Println("Enter the first date:")
	userDate1 := getUserDate()
	fmt.Println("Enter the second date:")
	userDate2 := getUserDate()
	diff := userDate2.Sub(userDate1)
	diffYears := int(diff.Hours() / 24 / 365)
	diffPostYears := int(diff.Hours()) / 24 % 365
	fmt.Println()
	fmt.Println(diffYears, "years", diffPostYears/30, "months", diffPostYears%30, "days")
	fmt.Println(diff.Hours()/24/30, "months", math.Mod(diff.Hours()/24, 30), "days")
	fmt.Println(diff.Hours()/24/7, "weeks", math.Mod(diff.Hours()/24, 7), "days")
	fmt.Println(diff.Hours()/24, "days")
	fmt.Println(diff.Hours(), "hours")
	fmt.Println(diff.Minutes(), "minutes")
	fmt.Println(diff.Seconds(), "seconds")
	//Cast int to all these inputs
}

func main() {
	//Check for wrong input

	var user string
	fmt.Println("Do you want to use calculator 1, 2, or 3?")
	fmt.Scanln(&user)
	switch user {
	case "1":
		calculatorOne()
	case "2":
		calculatorTwo()
	case "3":
		ageCalculator()
	}
}
