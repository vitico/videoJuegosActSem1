package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//random number generator XD
	var cant int
	fmt.Println("Cuantos numeros va a querer?")
	fmt.Scanln(&cant)
	for i := 0; i < cant; i++ {
		fmt.Printf("#%d = %d\n", i, rand.Int())
	}
}
