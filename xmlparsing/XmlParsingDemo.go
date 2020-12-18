package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	fmt.Println("Welcome to xml demo")

	xmlFile, err := os.Open("users.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("File opened successfully.")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var users Users

	xml.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User name: ", users.Users[0].Name)
		fmt.Println("User type: ", users.Users[0].Type)
		fmt.Println("Social-facebook: ", users.Users[0].Social.Facebook)
		fmt.Println("Social-twitter: ", users.Users[0].Social.Twitter)
		fmt.Println("Social-youtube: ", users.Users[0].Social.Youtube)
		fmt.Printf("\n")
	}

}
