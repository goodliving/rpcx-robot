package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/beevik/etree"
	"log"
)

var num int

func main()  {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("output.xml"); err != nil {
		log.Fatal(err)
	}

	root := doc.SelectElement("robot")
	f := excelize.NewFile()
	index := f.NewSheet("robotframework自动化用例执行结果")
	getTestInfo(root, f, num + 1)

	f.SetActiveSheet(index)
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func getTestInfo(e *etree.Element, f *excelize.File, n int) {

	suites := e.SelectElements("suite")
	tests := e.SelectElements("test")

	if len(tests) != 0 {
		for _, test := range tests {
			next := n + 1
			fmt.Println("next: ", next)
			for _, value := range test.Parent().Attr {
				if value.Key == "source" {
					f.SetCellValue("robotframework自动化用例执行结果", fmt.Sprintf("A%d", next), value.Value)
				}
			}

			for _, value := range test.Attr {
				if value.Key == "name" {
					f.SetCellValue("robotframework自动化用例执行结果", fmt.Sprintf("B%d", next), value.Value)
					fmt.Println("testName: ", value.Value)
				}
			}

			testStatus := test.SelectElement("status")
			for _, value := range testStatus.Attr {
				if value.Key == "status" {
					f.SetCellValue("robotframework自动化用例执行结果", fmt.Sprintf("C%d", next), value.Value)
				}

				if value.Key == "critical" {
					f.SetCellValue("robotframework自动化用例执行结果", fmt.Sprintf("D%d", next), value.Element().Text())
				}
			}
		}
	}
	if len(suites) == 0 {
		return
	}

	for _, v := range suites {
		getTestInfo(v, f, n + 1)
	}
}