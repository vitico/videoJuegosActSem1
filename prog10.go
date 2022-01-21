package main

import (
	"fmt"
	"sync"
	"time"
)

type ThreadDef struct {
	Seconds int
	Msg     string
}

var wg sync.WaitGroup // 1

func processThread(t ThreadDef) {
	defer wg.Done()
	time.Sleep(time.Duration(t.Seconds) * time.Duration(t.Seconds))
	fmt.Println(t.Msg)

}
func main() {

	var cant = 1
	fmt.Println("Cuantos thread quiere crear?")
	fmt.Scanln(&cant)

	var definiciones []ThreadDef
	for i := 0; i < cant; i++ {
		var t ThreadDef

		fmt.Printf("Cual sera el mensaje? ")
		fmt.Scanln(&t.Msg)

		fmt.Printf("y la duracion? ")
		fmt.Scanln(&t.Seconds)
		definiciones = append(definiciones, t)
	}

	fmt.Println("Comencare a correr los threads!!")
	wg.Add(cant)
	for i := 0; i < cant; i++ {
		go processThread(definiciones[i])
	}
	fmt.Println("Esperando threads!!")
	wg.Wait()
	fmt.Println("terminado!!")

}
