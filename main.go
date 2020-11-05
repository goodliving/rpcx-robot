package main

import (
	"fmt"
	"github.com/beevik/etree"
	"log"
)

func main()  {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("output.xml"); err != nil {
		log.Fatal(err)
	}

	root := doc.SelectElement("robot")
	suite := root.SelectElement("suite")
	for _, value := range suite.Child {
		fmt.Println("value: ", value)
	}
}
