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
fmt.Println("Double Signing :", isProposer)
fmt.Println("Block Density :", blockDensity)

}
