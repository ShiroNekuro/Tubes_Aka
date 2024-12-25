package main

import "fmt"

type Player struct {
	Name string
	MMR  int
}

// Sorting MMR
func bubbleSort(players []Player) {
	n := len(players)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if players[j].MMR < players[j+1].MMR {
				// Swap
				players[j], players[j+1] = players[j+1], players[j]
			}
		}
	}
}

func main() {
	players := []Player{
		{Name: "Alice", MMR: 1200},
		{Name: "Bob", MMR: 1300},
		{Name: "Charlie", MMR: 1100},
		{Name: "Diana", MMR: 1250},
	}
  option(players)
}

func option(players [])
	// eksekusi sort
	/*bubbleSort(players)

	fmt.Println("\nRanked Players:")
	for rank, player := range players {
		fmt.Printf("%d. %s: %d\n", rank+1, player.Name, player.MMR)
	}
