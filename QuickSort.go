// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"runtime"
// 	"time"
// )

// var numberOfProcessors = 12

// func GenerateRandomArray(n int) []int {

// 	var seed int64 = 328713721987319
// 	rand.Seed(seed)

// 	randomArray := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		randomArray[i] = rand.Intn(1000000)
// 	}

// 	return randomArray
// }

// func main() {

// 	casas := []int{100,
// 		1000,
// 		10000,
// 		100000,
// 		1000000,
// 		10000000}

// 	runtime.GOMAXPROCS(numberOfProcessors)

// 	println("Number of processors: ", numberOfProcessors)

// 	for _, tam := range casas {
// 		fmt.Println(tam, ' ', timeToMergeSort(tam))
// 	}
// }

// func timeToMergeSort(tam int) time.Duration {
// 	start := time.Now()

// 	numbersToBeSorted := GenerateRandomArray(tam)

// 	// print("Unsorted array: \n")
// 	// for i := 0; i < len(numbersToBeSorted); i++ {
// 	// 	print(numbersToBeSorted[i], " ")
// 	// }
// 	// print("\n")

// 	quickeSorte(numbersToBeSorted, 0, len(numbersToBeSorted)-1)

// 	// print("Sorted array: \n")
// 	// for i := 0; i < len(numbersToBeSorted); i++ {
// 	// 	print(numbersToBeSorted[i], " ")
// 	// }

// 	return time.Since(start)
// }

// func partition(arr []int, low, high int) int {
// 	pivot := arr[high]
// 	i := low
// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}
// 	arr[i], arr[high] = arr[high], arr[i]
// 	return i
// }

// func quickeSorte(vetorASerOrdenado []int, inicio, fim int) []int {

// 	if inicio < fim {
// 		var p int
// 		p = partition(vetorASerOrdenado, inicio, fim)
// 		canalTerminou := make(chan struct{}, 1)
// 		go func() {
// 			quickeSorte(vetorASerOrdenado, inicio, p-1)
// 			canalTerminou <- struct{}{}
// 		}()
// 		quickeSorte(vetorASerOrdenado, p+1, fim)
// 		<-canalTerminou
// 	}
// 	return vetorASerOrdenado
// }
