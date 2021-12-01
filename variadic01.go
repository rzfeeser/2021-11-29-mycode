/* Alta3 Research | RZFeeser 
   Variadic Functions - Go Lang function that accepts an indefinite number
   of args */

package main

import "fmt"

func hello(names ...string) {
    
    fmt.Print(names, " ")
    
    total := 0
    
    for _, name := range names {
        total += 1
	fmt.Println("Hello", name)
    }
    
    fmt.Println(total)
}


func main() {

    hello("larry", "steve")
    hello("jane", "raj", "tom")

    names := []string{"larry", "raj", "harry", "beth"}
    hello(names...)
}
