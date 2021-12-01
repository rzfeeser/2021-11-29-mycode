/* Alta3 Research | RZFeeser
   User Input - Playing with fmt.Scanf for user input */

package main

import "fmt"

func main() {
	var day, year int // capture the users day and year they were born
	var name, month string // capture the user and the month they were born

	fmt.Println("Using the following pattern, enter your: <name> <year> <month> <day>")
	fmt.Scanf("%s %d %s %d", &name, &year, &month, &day) // capture the data and assign to our vars

	fmt.Printf("You entered: %s %d %s %d\n", name, year, month, day)  // display on stdout the data that was captured
}
