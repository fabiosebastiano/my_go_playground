package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	//"os/exec"
	//"runtime"
)

// our struct which contains the complete
// array of all Users in the file
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

// a simple struct which contains all our
// social links
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	fmt.Println("- I'M TRYING TO OPEN AND READ AN XML FILE")
	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Successfully Opened users.xml")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &users)

	fmt.Println(users)

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	//fmt.Println("- YOU'RE RUNNIG ON: " + runtime.GOOS)
	//execute()
}

//esecuzioni di comandi @OS > runtime + os/exec
func execute() {

	out, err := exec.Command("ls", "-ltr").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed: ", string(out[:]))
}
