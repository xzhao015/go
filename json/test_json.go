package json

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/buger/jsonparser"
)

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//CheckJSON is a function.
func CheckJSON(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatal("file does not exist")
		return false
	}
	log.Println("file is exist")
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("File reading error", err)
		return false
	}
	log.Println("read file success")
	fullname, err := jsonparser.GetString(content, "person", "name", "fullName")
	if err != nil {
		log.Fatal("File reading error", err)
		return false
	}
	log.Println(fullname)
	var size int64
	if value, err := jsonparser.GetInt(content, "company", "size"); err == nil {
		size = value
		log.Println(size)
	} else {
		log.Fatal("cannot get company size", err)
		return false
	}
	jsonparser.Set(content, []byte("56"), "company", "size")
	return true
}
