package main

import "fmt"
import "time"
import "math/rand"

type Player struct {
	Name string
	MMR  int
}

// Fungsi utama Heap Sort
func heapSort(playersB []Player) {
    players := playersB
	n := len(players)

	// Bangun max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(players, n, i)
	}

	// Ekstraksi elemen satu per satu dari heap
	for i := n - 1; i > 0; i-- {
		// Pindahkan elemen root (maksimum) ke akhir array
		players[0], players[i] = players[i], players[0]

		// Heapify pada heap yang tersisa
		heapify(players, i, 0)
	}
}

// Fungsi untuk memastikan struktur max heap
func heapify(players []Player, n, i int) {
	largest := i       // Anggap node saat ini sebagai terbesar
	left := 2*i + 1    // Indeks anak kiri
	right := 2*i + 2   // Indeks anak kanan

	// Jika anak kiri lebih besar dari root
	if left < n && players[left].MMR < players[largest].MMR {
		largest = left
	}

	// Jika anak kanan lebih besar dari root
	if right < n && players[right].MMR < players[largest].MMR {
		largest = right
	}

	// Jika root bukan yang terbesar
	if largest != i {
		// Tukar root dengan anak yang lebih besar
		players[i], players[largest] = players[largest], players[i]

		// Rekursif heapify subtree yang terkena dampak
		heapify(players, n, largest)
	}
}

// Sorting MMR dan menampilkan MMR
func bubbleSort(playersA []Player, playersB []Player) {
	startA := time.Now()
	n := len(playersA)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if playersA[j].MMR < playersA[j+1].MMR {
				// Swap
				playersA[j], playersA[j+1] = playersA[j+1], playersA[j]
			}
		}
	}
// 	fmt.Println("\nRanked Players:")
// 	for rank, player := range playersA {
// 		fmt.Printf("%d. %s: %d\n", rank+1, player.Name, player.MMR)
// 	}
	durationA := time.Since(startA)
	fmt.Println("Execution time : ", durationA)
	option(playersA, playersB)
}

// func reset(players []Player){
// 	// Seed untuk random generator
// 	rand.Seed(time.Now().UnixNano())

// 	// Array nama contoh
// 	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack"}

// 	// Slice untuk menyimpan dummy data
// 	players := make([]Player, num)

// 	for i := 0; i < num; i++ {
// 		// Pilih nama secara acak
// 		name := names[rand.Intn(len(names))] + fmt.Sprintf("_%d", i+1)

// 		// Generate nilai MMR secara acak (misalnya, antara 1000 hingga 2000)
// 		mmr := rand.Intn(1001) + 1000

// 		// Tambahkan ke array
// 		players[i] = Player{Name: name, MMR: mmr}
// 	}
// 	option(players)
// }

//Inisialisasi Data player
func main() {
	// Seed untuk random generator
	rand.Seed(time.Now().UnixNano())
    num := 50000
	// Array nama contoh
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack"}

	// Slice untuk menyimpan dummy data
	playersA := make([]Player, num)
	playersB := make([]Player, num)

	for i := 0; i < num; i++ {
		// Pilih nama secara acak
		name := names[rand.Intn(len(names))] + fmt.Sprintf("_%d", i+1)

		// Generate nilai MMR secara acak (misalnya, antara 1000 hingga 2000)
		mmr := rand.Intn(1001) + 1000

		// Tambahkan ke array
		playersA[i] = Player{Name: name, MMR: mmr}
		playersB[i] = Player{Name: name, MMR: mmr}
	}
	option(playersA, playersB)
}

//Menu
func option(playersA []Player, playersB []Player){
	var menu int
	fmt.Println("===========Menu==========")
	fmt.Println("1. Bubble sort")
	fmt.Println("2. Heap sort")
	fmt.Println("3. Simulasi Match")
	fmt.Println("4. Reset")
	fmt.Println("5. Show Data")
	fmt.Println("==========================")
	fmt.Scan(&menu)
	switch menu{
		case 1:
			bubbleSort(playersA, playersB)
			
		case 2: 
			heapSort(playersB)
			startB := time.Now()
// 			fmt.Println("\nRanked Players:")
// 			for rank, player := range playersB {
// 				fmt.Printf("%d. %s: %d\n", rank+1, player.Name, player.MMR)
// 			}
			durationB := time.Since(startB)
			fmt.Println("Execution time : ", durationB)
			option(playersA, playersB)
		case 3:
// 			simulate(players)
		case 4: 
// 			reset(players)
        case 5:
//             for rank, player := range players {
// 				fmt.Printf("%d. %s: %d\n", rank+1, player.Name, player.MMR)
// 			}
// 			option(players)
	}
}

//simulasi match
// func simulate(players []Player){
// 	var name string
// 	var winlose string
// 	fmt.Println("Masukkan nama player: ")
// 	fmt.Scan(&name)
// 	id := findPlayer(players, name)
// 	if id == -1 {
// 		fmt.Println("Player tidak ditemukan")
// 		option(players)
// 	}
// 	fmt.Println("menang/kalah")
// 	fmt.Scan(&winlose)
// 	if winlose == "menang" {
// 		players[id].MMR += 50
// 	} else if winlose == "kalah" {
// 		players[id].MMR -= 50
// 	}
// 	option(players)
// }

//function untuk mencari player dan mereturn index
// func findPlayer(players []Player, name string) int{
// 	for i, player := range players {
// 		if player.Name == name {
// 			return i
// 		}
// 	}
// 	return -1
// }
