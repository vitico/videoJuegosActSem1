package main

import (
	"fmt"
)

func main() {
	fmt.Println("Jugemos algo. Intenta adivinar el numero en el que estoy pensando!")
	fmt.Println("Tienes 3 intentos para adivinar. Vamos intentalo!")
	var num = 3
	var resp = 0
	for tries := 0; tries < 3; tries++ {
		fmt.Printf("Intento #%d: ", tries+1)
		fmt.Scanln(&resp)
		if resp == num {
			break
		}
		fmt.Printf("Te has equivocado. te quedan %d intentos\n", 3-tries-1)
	}
	fmt.Print("Usted ")
	if resp != num {
		fmt.Print("no ")
	}
	fmt.Print("ha ganado")

}
