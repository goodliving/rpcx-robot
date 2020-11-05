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

	suite := doc.SelectElement("robot")
	getTestInfo(suite)


	//fmt.Println("suite: ", suite)
}

func getTestInfo(e *etree.Element) {

	suites := e.SelectElements("suite")
	tests := e.SelectElements("test")

	if len(tests) != 0 {
		for _, test := range tests {
			for _, value := range test.Parent().Attr {
				if value.Key == "source" {
					fmt.Println("module: ", value.Value)
				}
			}

			for _, value := range test.Attr {
				if value.Key == "name" {
					fmt.Println("testName: ", value.Value)
				}
			}

			testStatus := test.SelectElement("status")
			for _, value := range testStatus.Attr {
				if value.Key == "status" {
					fmt.Println("testStatus: ", value.Value)

					if value.Value == "FAIL" {
						msg := testStatus.SelectElement("msg")
						fmt.Println("msg: ", msg)
					}
				}

				if value.Key == "critical" {
					fmt.Println("critical: ", value.Element().Text())
				}
			}
		}
	}

	if len(suites) == 0 {
		return
	}

	for _, v := range suites {
		getTestInfo(v)
	}
}