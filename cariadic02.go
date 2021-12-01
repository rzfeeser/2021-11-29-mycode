/* Alta3 Research | RZFeeser 
   Variadic Functions - Go Lang function that accepts an indefinite number
   of args */

package main

import "fmt"

func hello(names ...string, age ...int) {
    
    fmt.Print(names, " ")
    
    total := 0
    
    for _, name := range names {
        total += 1
	fmt.Println("Hello", name)
    }
    
    fmt.Println(total)
}


func main() {

    //hello("larry", "steve", 22, 34)
    //hello("jane", "raj", "tom", 22, 43, 23)

    names := []string{"larry", "raj", "harry", "beth"}
    ages := []int{55, 63, 23, 22}
    hello(names..., ages...)
}
