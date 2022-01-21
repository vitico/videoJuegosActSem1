package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//encrypt
	f, err := os.Create("/tmp/dat1")
	check(err)
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	fmt.Fprintln(writer, "Hello World")
}
