package main

import "fmt"

func main() {
	fmt.Println(`-- Simple Calculator --`)
	var opt int = -1
	for ok := true; ok; {
		fmt.Println("Que desea hacer?")
		fmt.Println("Las opciones son: ")
		fmt.Print(`1) ^
2) +
3) -
4) *
5) /
6) Exit
> `)
		fmt.Scanln(&opt)
		if opt > 6 || opt < 1 {
			fmt.Println("Opcion invalida. Por favor, seleccione una de las opciones indicadas.")
			continue
		}
		if opt == 6 {
			ok = false
			continue
		}
		var n1, n2 int
		fmt.Print("Cual es el 1er valor? ")
		fmt.Scan(&n1)
		fmt.Print("y el 2do? ")
		fmt.Scan(&n2)
		resultado := 0
		switch opt {
		case 1:
			resultado = n1 ^ n2
			break
		case 2:
			resultado = n1 + n2
			break
		case 3:
			resultado = n1 - n2
			break
		case 4:
			resultado = n1 * n2
			break
		case 5:
			resultado = n1 / n2
			break

		}
		fmt.Printf("El resultado es '%d'\n", resultado)
	}
	fmt.Println("Ha sido un placer calcular para usted. Adios!")

}
