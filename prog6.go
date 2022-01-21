package main

import (
	"fmt"
	"net/http"
)

var visits int = 0

func greet(w http.ResponseWriter, r *http.Request) {
	visits = visits + 1
	fmt.Fprintf(w, "Esta pagina ha sido visitada %d veces", visits)
}
func doNothing(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/favicon.ico", doNothing)

	http.ListenAndServe(":8080", nil)
}
