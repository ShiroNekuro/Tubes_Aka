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

// Sorting MMR dengan bubble sort
func bubbleSort(playersA []Player) {
	
	n := len(playersA)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if playersA[j].MMR < playersA[j+1].MMR {
				// Swap
				playersA[j], playersA[j+1] = playersA[j+1], playersA[j]
			}
		}
	}
}

//Inisialisasi Data player
func main() {
	// Seed untuk random generator
	rand.Seed(time.Now().UnixNano())
	
	// Jumlah data yang akan digenerate
    	num := 200000
    
	// Array nama contoh
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Udin"}

	// Slice untuk menyimpan dummy data
	playersA := make([]Player, num)
	playersB := make([]Player, num)

	for i := 0; i < num; i++ {
		// Pilih nama secara acak
		name := names[rand.Intn(len(names))] + fmt.Sprintf("_%d", i+1)

		// Generate nilai MMR secara acak
		mmr := rand.Intn(4000)

		// Tambahkan ke array
		playersA[i] = Player{Name: name, MMR: mmr}
		playersB[i] = Player{Name: name, MMR: mmr}
	}
	
	//menghitung execution time heap sort
	startB := time.Now()
	heapSort(playersB)
	durationB := time.Since(startB)
	fmt.Println("Heap sort execution time : ", durationB)
	
	//menghitung execution time bubble sort
	startA := time.Now()
	bubbleSort(playersA)
	durationA := time.Since(startA)
	fmt.Println("Bubble sort execution time : ", durationA)
}
