package main

import "fmt"

var blockTime int64

func main() {

voteRound()

}

func voteRound() {

round := 2  
votes := 6  
hasMajority:= votes/2 + 2  
  
fmt.Println("Round :", round)  
fmt.Println("Votes :", votes)  
fmt.Println("Has Majority :", hasMajority)

}
