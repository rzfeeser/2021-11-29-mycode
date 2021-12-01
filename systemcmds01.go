/* Alta3 Research | RZFeeser
   Executing system commands  */

package main

import (
    "log"
    "os/exec"
    "fmt"
)

func main() {

    // prepares a "cmd" struct
    out, err := exec.Command("ls").Output()

    // err := cmd.Run() // start the command and wait to finish

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))

}

