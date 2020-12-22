package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Avatar is a struct.
type Avatar struct {
	URLS  string `json:"url"`
	TypeS string `json:"type"`
}

//GithubS is a struct.
type GithubS struct {
	Handle    string `json:"handle"`
	Followers int    `json:"followers"`
}

//NameS is a struct.
type NameS struct {
	First    string `json:"first"`
	Last     string `json:"last"`
	FullName string `json:"fullName"`
}

//CompanyS is a struct.
type CompanyS struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

//CompanyT is a struct.
type CompanyT struct {
	Name string `json:"name1"`
	Size int    `json:"size1"`
}

//PersonS is a struct.
type PersonS struct {
	Name    NameS    `json:"name"`
	Github  GithubS  `json:"github"`
	Avatars []Avatar `json:"avatars"`
}

//GitPerson is a struct.
type GitPerson struct {
	Person   PersonS  `json:"person"`
	Company  CompanyS `json:"company"`
	Company1 CompanyT `json:"company1"`
}

//TestStandardJSON is a function.
func TestStandardJSON(jsonFilePath string) {
	//read json file
	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal("File reading error", err)
	}
	log.Println("read file success")
	newgitperson := &GitPerson{}
	err = json.Unmarshal(content, &newgitperson)

	fmt.Println(newgitperson.Person.Name.First)
	fmt.Println(newgitperson.Person.Avatars[0].URLS)
	fmt.Println("company.name:", newgitperson.Company.Name)
	fmt.Println("company1.name:", newgitperson.Company1.Name)
	for index, value := range newgitperson.Person.Avatars {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
		fmt.Println(value.URLS, value.TypeS)
		newgitperson.Person.Avatars[index].URLS = "http://1234567"
	}

	//modify json file
	newgitperson.Person.Name.First = "zhaoxia"
	newAvatar := Avatar{
		URLS:  "https://avatars1.githubusercontent.com/u/14009?v=3\u0026s=460",
		TypeS: "thumbnail"}
	newgitperson.Person.Avatars = append(newgitperson.Person.Avatars, newAvatar)
	newcompanyT := CompanyT{
		Name: "zhaoxia",
		Size: 3000}
	newgitperson.Company1 = newcompanyT
	data, err := json.MarshalIndent(newgitperson, "", "    ") // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	checkError(err)
	err = ioutil.WriteFile(jsonFilePath, data, 0777)
	checkError(err)

}
