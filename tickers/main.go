package main

import (
	"fmt"
	"time"
)

func backgroundfunc() {
	ticker := time.NewTicker(time.Second * 1)

	for _ = range ticker.C {
		fmt.Println("-- TICK")
	}
}

func main() {
	fmt.Println("- STARTING GO TICKERS TUTORIAL -")

	go backgroundfunc()

	//Verrà eseguita prima di iniziare a vedere i tick...
	fmt.Println("- GO TICKERS TUTORIAL ENDED -")

	//POI OPERAZIONE INFINITA CHE PERMETTE DI VEDERE I TICK
	// anche for andava bene
	select {}
}
