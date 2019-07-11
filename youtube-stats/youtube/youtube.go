package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Response che modella la struttura del JSON
//che ci viene restituita dalle API
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

//Items Contiene l'id + le statistiche per un dato canale
type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

//Stats salviamo tutte le info importanti
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

//GetSubscribers recupera tutte le stats
func GetSubscribers() (Items, error) {
	var response Response

	//nuova GET con query params
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)

	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}

	//da qui definiamo i query param
	queryparam := req.URL.Query()
	queryparam.Add("key", os.Getenv("YOUTUBE_KEY"))
	queryparam.Add("id", os.Getenv("CHANNEL_ID"))
	queryparam.Add("part", "statistics")
	//prearo il nuovo URL
	req.URL.RawQuery = queryparam.Encode()

	//eseguiamo la request con tutti i parametri
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERRORE NELLA CHIAMATA: ", err)
		return Items{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	//se arriviamo qui è perchè abbiamo ottenuto un 200
	//leggo l'ogggetto arrivato
	body, _ := ioutil.ReadAll(resp.Body)
	//alla fine unmarshal del risultato dentro nostra struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Items{}, err
	}

	//mandiamo indietro solo il primo elemento
	return response.Items[0], nil
}
