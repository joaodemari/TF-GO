// por Fernando Dotti - PUCRS -
// aqui utilizamos como tarefa "cpu intensiva" calcular um valor
// primo com um determinado número de casas.
// em "casas" define-se tamanhos de 3, 6, 9, 12, 15, 18 casas
// em seguida, mede-se o tempo para gerar um valor primo com cada número de casas,
// com a função timeToGenPrime(...)
// a geração de um número primo é,
//    sorteia um valor com o numero de casas,
//    verifica se é primo
// repete até achar um primo
// diferentes execuções, iniciando com diferentes seeds, terão diferentes tempos.

// o trabalho "genPrime(tamanho)" representa uma computação intensiva,
// que é mais custosa quanto o tamanho do primo a ser gerado.

// suponha que voce tem que gerar um conjunto de N valores primos.
// calcule o speedup de uma solução paralela com P núcleos processadores.
// faca uma análise de speedup para os diferentes tamanhos de valores primos.
// Exemplos para
//   N: 50;  ou mais
//   P:  conforme seu hardware {2, 4, 6, 8, 10, 12}
//   tamanhos dos valores, coforme exemplificado no programa abaixo.

//Nomes: João Demari, Guilherme Kuhn, João Lima, Julio Cesar Filho

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var numberOfProcessors int = 12

func main() {
	casas := [6]int{999,
		999999,
		999999999,
		999999999999,
		999999999999999,
		9999999999999999}

	fmt.Println("Number of processors: ", numberOfProcessors)

	runtime.GOMAXPROCS(numberOfProcessors)
	var seed int64 = 328713721987319
	rand.Seed(seed)

	for _, tam := range casas {
		fmt.Println(tam, ' ', timeToGenPrime(tam))
	}
}

func genPrime(tam int) {
	isPrimeChan := make(chan int, numberOfProcessors)
	notPrimo := true
	for i := 0; i < numberOfProcessors; i++ {
		v := rand.Intn(tam)
		go isPrime(v, isPrimeChan)
	}
	primo := 0
	for notPrimo {
		vToBeConsulted := <-isPrimeChan
		if vToBeConsulted != 0 {
			notPrimo = false
			primo = vToBeConsulted
		} else {
			v := rand.Intn(tam)
			go isPrime(v, isPrimeChan)
		}
	}

	println(primo)
}

func timeToGenPrime(tam int) time.Duration {
	start := time.Now()
	genPrime(tam)
	return time.Since(start)
}

func isPrime(p int, isPrimeSync chan int) {
	if p%2 == 0 {
		isPrimeSync <- 0
	}
	for i := 3; i*i <= p; i += 2 {
		if p%i == 0 {
			isPrimeSync <- 0
		}
	}
	isPrimeSync <- p
}
