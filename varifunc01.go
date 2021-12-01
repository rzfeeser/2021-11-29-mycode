/* Alta3 Research | RZFeeser
   Variadic Functions - Looking at functions in Go that accept
   an indefinite number of arguments. */

package main

import "fmt"

// variadic function, accepts n number of strings (people's names)
// display each string with a hello in front of it
func hello(names ...string) {
   fmt.Println(names, " ")

   for _, name := range names {
      fmt.Println("Why Hello,", name)
   }
}



// main
func main() {

	// call function with two names
	hello("Afzal", "Ananta")

	// call function with three names
	hello("Binaya", "Doug", "Elizabeth")

	// "unpack" a slice of "n" names and dump into our function
	students := []string{"Eric", "Evan", "Heidi", "Lavanya", "Sandeep"}
	hello(students...)  // the trailing ... let us "unpack" the slice and pass each value individually
}
