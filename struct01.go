/* RZFeeser | Alta3 Research
   Struct - Learning about creating structures, go's version of OOP */
package main

import "fmt"

// tracking information about a single student
type Student struct {
	Name string
	Location string
	ID int
}

func main() {

    alpha := Student{"Alice", "PA", 24601}
    alpha.ID = 17111

    fmt.Println(alpha)

    fmt.Println("The student", alpha.Name, "is in", alpha.Location)

}
