package main

import "fmt"

func addToShoppingList(item string, shoppingList *[]string) {
   *shoppingList = append(*shoppingList, item)
   //return *shoppingList
}


func main() {
var shoppingList []string
cookies, vaccine, lifePurpose := "oreos", "toilet paper", "Animal Crossing: New Horizons"
addToShoppingList(cookies, &shoppingList)
addToShoppingList(vaccine, &shoppingList)
addToShoppingList(lifePurpose, &shoppingList)
fmt.Println("Shopping list contains:", shoppingList)
fmt.Println(shoppingList[len(shoppingList)-1])
}
