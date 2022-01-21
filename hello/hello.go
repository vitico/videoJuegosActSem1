package main

import "fmt"

func hello() {
	fmt.Println("-- Pokemon Searcher --")
	fmt.Println("Introdusca el nombre del pokemon del que desea conseguir informacion:")
	fmt.Print("> ")
	var search string
	fmt.Scanln(search)
	fmt.Printf("to search %s\n", search)
}
