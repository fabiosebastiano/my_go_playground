package main

import (
	"fmt"
)

var risposta = AllaDomanda()

func AllaDomanda() int {
	return 42
}

/**
ESEMPIO DI USO DI FUNZIONE INIT
*/

//var name string

func init() {
	fmt.Println("CALLING INIT FUNC")
	//name = "fabio"
	risposta = 0
}
func main() {
	fmt.Println("+starting application: ", risposta)
}
