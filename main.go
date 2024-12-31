package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string `xml:"profiles>name"`
	Age  int    `xml:"profiles>age"`
}

func main() {
	xmlData := `
        <Person>
			<profiles>
				<name>John Doe</name>
            	<age>30</age>
			</profiles>
        </Person>
    `

	var p Person
	err := xml.Unmarshal([]byte(xmlData), &p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
