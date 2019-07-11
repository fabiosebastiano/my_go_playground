package main

import (
	"fmt"
	"log"
	"net/http"
	//	"youtube-stats/websocket"
)

func init() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CIAO MONDO")
}

//espone tutte le stats del canale via websocket
func stats(w http.ResponseWriter, r *http.Request) {
	/*
		//per prima cosa chiamo Update function per modificare
		//la connessione da HTTP a WEBSOCKET
		ws, err := websocket.Upgrade(w, r)
		if err != nil {
			fmt.Fprintf(w, "%+v\n", err)
		}
		//ora posso chiamare funzione che polla ogni 5 secondi
		//e poi scrive nella connessione websocket
		go websocket.Writer(ws)
	*/
}

func main() {
	fmt.Println("YouTube Subscribe Monitor")

}
