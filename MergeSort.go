
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

// 	mergesort(numbersToBeSorted)

// 	return time.Since(start)
// }

// func mergesort(vetorASerOrdenado []int) {
// 	if len(vetorASerOrdenado) <= 1 {
// 		return
// 	}

// 	meioDoArray := len(vetorASerOrdenado) / 2

// 	arrayDaEsquerda := make([]int, meioDoArray)
// 	arrayDaDireita := make([]int, len(vetorASerOrdenado)-meioDoArray)
// 	copy(arrayDaEsquerda, vetorASerOrdenado[:meioDoArray])
// 	copy(arrayDaDireita, vetorASerOrdenado[meioDoArray:])

// 	canalTerminou := make(chan struct{}, 1)

// 	go func() {
// 		mergesort(arrayDaEsquerda)
// 		canalTerminou <- struct{}{}
// 	}()

// 	mergesort(arrayDaDireita)
// 	<-canalTerminou

// 	merge(vetorASerOrdenado, arrayDaEsquerda, arrayDaDireita)
// }

// func merge(resultado, esquerda, direita []int) {
// 	var iEsquerda, iDireita, iResultado int

// 	for iEsquerda < len(esquerda) || iDireita < len(direita) {
// 		if iEsquerda < len(esquerda) && iDireita < len(direita) {
// 			if esquerda[iEsquerda] <= direita[iDireita] {
// 				resultado[iResultado] = esquerda[iEsquerda]
// 				iEsquerda++
// 			} else {
// 				resultado[iResultado] = direita[iDireita]
// 				iDireita++
// 			}
// 		} else if iEsquerda < len(esquerda) {
// 			resultado[iResultado] = esquerda[iEsquerda]
// 			iEsquerda++
// 		} else if iDireita < len(direita) {
// 			resultado[iResultado] = direita[iDireita]
// 			iDireita++
// 		}
// 		iResultado++
// 	}
// }
