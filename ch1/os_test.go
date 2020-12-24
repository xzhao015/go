package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOs(t *testing.T) {
	//copy
	r := strings.NewReader("some io.Reader stream to be read\n")
	b, err := io.Copy(os.Stdout, r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)
	//WriteString
	io.WriteString(os.Stdout, "Hello World")
}

func TestIoutil(t *testing.T) {
	//ReadAll return ([]byte, error)
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s", b)
	//ReadDir returns a list of directory entries sorted by filename []os.FileInfo, error
	files, err := ioutil.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	//ReadFile ([]byte, error)
	content, err := ioutil.ReadFile("./variable.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)
}

func TestIoutilWriteFile(t *testing.T) {
	//WriteFile
	//TempDir return (name string, err error)
	content, _ := ioutil.ReadFile("./variable.go")
	dir, err := ioutil.TempDir("C:/Users/ehxazai/go/src/github.com/xzhao015/go/ch1", "example")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(dir)
	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	t.Log(tmpfn)
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		t.Fatal(err)
	}
	t.Log(os.Stat(tmpfn))
}
func TestIoutilTmpFile(t *testing.T) {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
