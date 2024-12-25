package main

import "fmt"

type Player struct {
	Name string
	MMR  int
}

// Sorting MMR dan menampilkan MMR
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
	fmt.Println("\nRanked Players:")
	for rank, player := range players {
		fmt.Printf("%d. %s: %d\n", rank+1, player.Name, player.MMR)
	}
	option(players)
}

func reset(players []Player){
	players = []Player{
		{Name: "Alice", MMR: 1200},
		{Name: "Bob", MMR: 1300},
		{Name: "Charlie", MMR: 1100},
		{Name: "Diana", MMR: 1250},
	}
	option(players)
}

//Inisialisasi Data player
func main() {
	players := []Player{
		{Name: "Alice", MMR: 1200},
		{Name: "Bob", MMR: 1300},
		{Name: "Charlie", MMR: 1100},
		{Name: "Diana", MMR: 1250},
	}
	option(players)
}

//Menu
func option(players []Player){
	var menu int
	fmt.Println("===========Menu==========")
	fmt.Println("1. Tampilkan data")
	fmt.Println("2. Simulasi Match")
	fmt.Println("3. Reset")
	fmt.Println("==========================")
	fmt.Scan(&menu)
	switch menu{
		case 1:
			bubbleSort(players)
		case 2:
			simulate(players)
		case 3: 
			reset(players)
	}
}

//simulasi match
func simulate(players []Player){
	var name string
	var winlose string
	fmt.Println("Masukkan nama player: ")
	fmt.Scan(&name)
	id := findPlayer(players, name)
	if id == -1 {
		fmt.Println("Player tidak ditemukan")
		option(players)
	}
	fmt.Println("menang/kalah")
	fmt.Scan(&winlose)
	if winlose == "menang" {
		players[id].MMR += 50
	} else if winlose == "kalah" {
		players[id].MMR -= 50
	}
	option(players)
}

//function untuk mencari player dan mereturn index
func findPlayer(players []Player, name string) int{
	for i, player := range players {
		if player.Name == name {
			return i
		}
	}
	return -1
}
