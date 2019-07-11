//youtube_monitor/websocket/websocket.go
package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fabiosebastiano/my_go_playground/youtube-stats/youtube"
	"github.com/gorilla/websocket"
)

//Read+Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Upgrade prende in input una richesta in arrivo e fa upgrade
//in una connessione socket
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	//permette ad host diversi dal nostro di connettersi al nostro websocket server
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	
	r.Header.Set("Connection", "upgrade")
	r.Header.Set("Upgrade", "websocket")
	r.Header.Set("Sec-Websocket-Version", "13")
	r.Header.Set("Sec-Websocket-Key", "LK2Qjb+GReKaqQrn4zEvsA==")
	
	//creiamo la connessione websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("!! errore durante upgrade!!! ", err)
		return ws, err
	}

	//se non ci sono errori, ritorno la nostra connessione websocket
	return ws, nil
}

//Writer è una funzione che, dato in input il puntatore al socket appena aperto
//usa tickers ogni 5 secondi per recuperare le info
func Writer(conn *websocket.Conn) {
	//for loop che dura per quanto sta aperta la connessione socket
	for {
		ticker := time.NewTicker(time.Second * 5)

		for t := range ticker.C {
			fmt.Printf("Aggiorno le Stats: %+v\n", t)

			//chiamo il servizio interno
			items, err := youtube.GetSubscribers()
			if err != nil {
				fmt.Println(err)
			}
			//se non ci sono errori, marshal dell'oggetto in jsonString
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			
			// and finally we write this JSON string to our WebSocket
			// connection and record any errors if there has been any
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}
		}

	}
}
