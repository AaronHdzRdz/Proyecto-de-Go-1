package main

import (
	"fmt"
	"sync"
)

// factorial calcula n! de forma recursiva.
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// worker toma números del canal jobs, calcula el factorial y envía los pares al canal results.
func worker(id int, jobs <-chan int, results chan<- [2]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		fmt.Printf("✨ Worker %d procesando %d\n", id, n)
		results <- [2]int{n, factorial(n)}
	}
}

func main() {
	// Lista de números a procesar
	numbers := []int{5, 7, 10, 12, 15}

	jobs := make(chan int, len(numbers))
	results := make(chan [2]int, len(numbers))
	var wg sync.WaitGroup

	// Lanzamos N workers concurrentes
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Enviamos los jobs y cerramos el canal
	for _, n := range numbers {
		jobs <- n
	}
	close(jobs)

	// Esperamos a que todos terminen y cerramos results
	wg.Wait()
	close(results)

	// Mostramos los resultados
	fmt.Println("\n🎉 Resultados:")
	for res := range results {
		fmt.Printf("Factorial de %d = %d\n", res[0], res[1])
	}
}
