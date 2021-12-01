/* Alta3 Research | RZFeeser
   Mimic a do-while (infinite True) loop with go */

package main

import "fmt"

func main() {

     z := 0
     for {
          z +=1
	  fmt.Println(z)
	  if z == 20 {      // the condition stops matching
		break        // break out of the loop
		}
	}


}
