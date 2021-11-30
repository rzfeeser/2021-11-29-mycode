/* Alta3 Resarch | RZFeeser
   Struct - intro to breaking down JSON */

package main

import "fmt"

type Location struct {
    Latitude string
    Longitude string
}

type Issdata struct {
	Timestamp int64
	IssPosition Location
	Message string
}

func main() {
    LatLon := Location{"21.8297", "-38.6633"}

    myiss := Issdata{4594234, LatLon, "success"}

    fmt.Println(myiss)

    fmt.Println(myiss.IssPosition)

    var otherplace Location   // otherplace.Latitude == ""

    fmt.Println(otherplace.Latitude)
}
