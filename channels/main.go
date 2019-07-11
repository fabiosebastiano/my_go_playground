package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("+STARTING OF CHANNELS APP+")
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// equivalente a for senza condizione, ma è più chiaro
	for l := range c {
		go func(link string) { //funziona anonima o FUNCTIONAL LITERALS
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l) //chiamata alla funziona anonomia con parametro
	}

	fmt.Println("+END OF CHANNELS APP+")
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Something wrong with site: ", url)
		c <- url
		return
	}
	fmt.Println("Site ", url, " is OK")
	c <- url
}
