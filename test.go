package main

import (
	"github.com/xzhao015/go/yaml_test"
)

func main() {
	//dateString := "20th Oct 2052"
	//fmt.Println(ch2.ReformatDate(dateString))
	//a := []int{1, 2, 3, 4, 5, 6}
	//output := ch2.RangeSum(a, 6, 1, 7)
	//fmt.Println("output is ", output)
	//output := ch2.MinDifference(a)
	//fmt.Println("output is ", output)
	//json.TestStandardJSON2("json/test.json")
	yaml_test.ParserYaml("yaml_test/test.yaml")
}
