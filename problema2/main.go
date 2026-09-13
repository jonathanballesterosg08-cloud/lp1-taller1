package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular tareas que toman tiempo con time.Sleep y comparar
// ejecución secuencial vs concurrente midiendo la duración total.
// TODO: completa las secciones marcadas para observar la mejora.

func tarea(id int, dur time.Duration) {
	fmt.Printf("[tarea %d] iniciando, dur=%v\n", id, dur)
	time.Sleep(dur)
	fmt.Printf("[tarea %d] finalizada\n", id)
}

func secuencial(durs []time.Duration) time.Duration {
	inicio := time.Now()
	// TODO: ejecutar las tareas en orden, sin goroutines
	for i, d := range durs {
		tarea(i+1, d)
	}
	return time.Since(inicio)
}

func concurrente(durs []time.Duration) time.Duration {
	inicio := time.Now()
	var wg sync.WaitGroup
	// TODO: lanzar cada tarea en su propia goroutine y esperar con WaitGroup
	for i, d := range durs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			tarea(i+1, d)
		}()

	}
	wg.Wait()
	return time.Since(inicio)
}

func main() {
	// TODO: experimenta con diferentes duraciones
	durs := []time.Duration{
		700 * time.Millisecond,
		500 * time.Millisecond,
		1 * time.Second}

	d1 := secuencial(durs)
	fmt.Printf("Duración SEC: %.3f segundos/n", d1.Seconds())

	d2 := concurrente(durs)
	fmt.Printf("Duración CONC: %.3f segundos/n", d2.Seconds())

	fmt.Println("Nota: la ejecución concurrente debería ser ~max(durs). Cambia valores y observa.")
}
