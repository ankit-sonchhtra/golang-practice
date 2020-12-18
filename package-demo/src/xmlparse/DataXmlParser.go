package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"structure"
)

func main() {

	xmlFile, err := os.Open("data.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	var data structure.Data

	xmlData, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(xmlData, &data)

	for i := 0; i < len(data.Persons); i++ {

		fmt.Println("FirstName:: ", data.Persons[i].FirstName)
		fmt.Println("LastName:: ", data.Persons[i].LastName)
		fmt.Println("Type:: ", data.Persons[i].Type)
		if data.Persons[i].Address.City != "" {
			fmt.Println("City:: ", data.Persons[i].Address.City)
		}

		if data.Persons[i].Address.State != "" {
			fmt.Println("State:: ", data.Persons[i].Address.State)
		}
		fmt.Println("\n")
	}
}
