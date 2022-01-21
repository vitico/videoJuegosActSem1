package main

import (
	"fmt"
	"time"
)

func main() {
	var ano = 0
	fmt.Println("Dime el ano en el que naciste, y yo te dire tu edad.")
	fmt.Print("> ")
	fmt.Scanln(&ano)
	var curAno = time.Now().Year()
	var edad = curAno - ano
	if edad > 200 {
		fmt.Println("Minimo seras vampiro, eres muy viejo! tienes")
	} else if edad > 140 {
		fmt.Print("Has roto record!! Tienes mas de 150 anos! aunque mas precisamente, son")
	} else {
		fmt.Print("Solo tienes")
	}
	fmt.Printf(" %d anos\n", edad)
}
