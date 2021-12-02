/* Mutexes - protecting state */

package main

import (
    "fmt"
    // "sync"   // mutex support
    "time"
)

func main() {

    // Balance: shared variable
    balance := 0

    // A single mutex to be used by both goroutines
    // var m sync.Mutex

    // Function that accesses balance multiple times
    counter := func() {
        // m.Lock()
        for i := 0; i < 10000000; i++ {
            balance = balance + 10
        }
        // m.Unlock()
    }

    // Starting two goroutines
    go counter()
    go counter()

    // Waiting and then printing balance
    time.Sleep(time.Second)
    fmt.Printf("Balance: %v", balance) // should be 200000000
}
