package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// TODO: completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		time.Sleep(500 * time.Millisecond)
		// TODO: simular trabajo

		ch <- x * x
	}()
	return ch
}
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()

			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// TODO: crea varios futuros y recolecta sus resultados: f1, f2, f3
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(5)
	f3 := asyncCuadrado(10)

	// TODO: Opción 1: esperar cada futuro secuencialmente
	fmt.Println("Opción 1:")
	fmt.Println("Resultado f1:", <-f1)
	fmt.Println("Resultado f2:", <-f2)
	fmt.Println("Resultado f3:", <-f3)
	//crear nuevos futuros para la opción 2
	f4 := asyncCuadrado(2)
	f5 := asyncCuadrado(5)
	f6 := asyncCuadrado(10)

	// TODO: Opción 2: fan-in (combinar múltiples canales)
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.
	fmt.Println("\nOpción 2 - Fan-in:")

	resultados := fanIn(f4, f5, f6)

	for resultado := range resultados {
		fmt.Println("Resultado:", resultado)
	}

	fmt.Println("Todos los futuros terminaron.")

}
