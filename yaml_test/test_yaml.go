package yaml_test

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

//ParserYaml is a function.
func ParserYaml(filepath string) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("File reading error", err)
	}
	log.Println(string(content))

	var t map[string]interface{}
	if err := yaml.Unmarshal(content, &t); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	num := t["kind"]
	fmt.Println(num)
}
