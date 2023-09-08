package main

import (
	"fmt"
	"time"
)

type Holiday struct {
	Month    int
	Day      int
	Name     string
	IsPublic bool
}

var holidays = []Holiday{
	{1, 1, "New Year's Day", true},
	{7, 4, "Independence Day", true},
	{12, 25, "Christmas Day", true},
	// Add more holidays as needed
}

func isHoliday(year, month, day int) bool {
	for _, holiday := range holidays {
		if holiday.Month == month && holiday.Day == day {
			return true
		}
	}
	return false
}

func printCalendar(year int) {
	today := time.Now()
	currentMonth := today.Month()
	currentDay := today.Day()

	for month := 1; month <= 12; month++ {
		// Create a date for the first day of the month
		firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

		// Find the weekday of the first day (0: Sunday, 1: Monday, ..., 6: Saturday)
		weekday := firstDay.Weekday()

		// Print the header with the month and year
		header := fmt.Sprintf("   %s %d", firstDay.Month().String(), year)
		if currentMonth == time.Month(month) {
			header += " *"
		}
		fmt.Println(header)
		fmt.Println("Su,  Mo,  Tu,  We,  Th,  Fr,  Sa, ")

		// Print leading spaces to align the first day of the month correctly
		for i := 0; i < int(weekday); i++ {
			fmt.Print("   ")
		}

		// Get the last day of the month
		lastDay := firstDay.AddDate(0, 1, -1)

		// Loop through the days of the month
		for day := firstDay.Day(); day <= lastDay.Day(); day++ {
			// Check if the current day is today and mark it
			highlight := " "
			if currentMonth == time.Month(month) && currentDay == day {
				highlight = "*"
			}

			// Check if the current day is a holiday
			holidayMarker := " "
			if isHoliday(year, month, day) {
				holidayMarker = "H" // You can use any marker you prefer
			}

			// Print the day with leading zero if needed and apply highlights
			fmt.Printf("%s%2d%s ", highlight, day, holidayMarker)

			// If it's the last day of the week (Saturday), start a new line
			if weekday == 6 {
				fmt.Println()
				weekday = 0
			} else {
				weekday++
			}
		}

		// Print a newline at the end to separate months
		fmt.Println()
	}
}

func main() {
	currentTime := time.Now()
	year := currentTime.Year()

	printCalendar(year)

	fmt.Println("Press Enter to close.")
	fmt.Scanln() // Wait for the user to press Enter
}
