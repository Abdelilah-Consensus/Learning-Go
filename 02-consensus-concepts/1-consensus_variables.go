package main

import "fmt"

// Defining variables at the general level (Global scope) | zero value

var blockHeight int
var currentTerm int
var isProposer string
var voteCount uint64

func main() {

// Printing variables on lines

fmt.Println("Block Height :", blockHeight)
fmt.Println("Current Term :", currentTerm)
fmt.Println("Is Proposer :", isProposer)
fmt.Println("Vote Count :", voteCount)

}
