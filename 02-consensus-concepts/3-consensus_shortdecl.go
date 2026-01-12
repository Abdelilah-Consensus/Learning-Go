package main

import "fmt"

// Definition of a variable: Block creation time

var blockTime int64

func main() {

// Create a new function inside main

voteRound()

}

// Calling the new function

func voteRound() {

round := 2  
votes := 6  
hasMajority := votes/2 + 2  
  
fmt.Println("Round :", round)  
fmt.Println("Votes :", votes)  
fmt.Println("Has Majority :", hasMajority)

}
