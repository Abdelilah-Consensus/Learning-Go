
// This is just the beginning, although the code isn't very well linked to the concepts, but I will learn until I can write real, meaningful code.


package main

import "fmt"

// Defining variables at the general level (Global scope) | zero value

var blockHeight uint64 
var peerCount int    // Network Monitoring 
var hasDoubleSign bool
var blockDensity uint64

func main() {

// Printing variables on lines

fmt.Println("Block Height :", blockHeight)
fmt.Println("Network Monitoring :", peerCount)
fmt.Println("Double Signing :", hasDoubleSign)
fmt.Println("Block Density :", blockDensity)

}
