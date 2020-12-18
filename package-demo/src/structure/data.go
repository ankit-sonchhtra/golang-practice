package structure

import "encoding/xml"

type Data struct {
	XMLName xml.Name `xml:"data"`
	Persons []Person `xml:"person"`
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	Type      string   `xml:"type,attr"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	Address   Address  `xml:"address"`
}

type Address struct {
	XMLName xml.Name `xml:"address"`
	City    string   `xml:"city"`
	State   string   `xml:"state"`
}
