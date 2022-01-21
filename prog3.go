package main

import (
	"fmt"
	"time"
)

func sleepSay(tiempo time.Duration, msg string) {
	time.Sleep(tiempo)
	fmt.Println(msg)
}
func main() {
	go sleepSay(time.Second, "Hello")
	go sleepSay(3*time.Second, "!")
	sleepSay(2*time.Second, "world")
	//imprime Hello y World, pero no a "!" por que el programa se cierra a los 2 segundos, sin esperar los 3 para el "!"
}
