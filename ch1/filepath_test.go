package ch1

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T) {
	//Abs return(string, error)
	absPath, err := filepath.Abs("file.go")
	if err != nil {
		t.Fatalf("get file abs path failed")
	} else {
		t.Log(absPath)
	}
	//Base return string
	t.Log(filepath.Base(absPath))
	//Clean return string
	t.Log(filepath.Clean(absPath))
	//Dir return string
	t.Log(filepath.Dir(absPath))
	//Ext
	t.Logf("No dots: %q\n", filepath.Ext("index"))
	t.Logf("One dot: %q\n", filepath.Ext("index.js"))
	t.Logf("Two dots: %q\n", filepath.Ext("main.test.js"))
	//Join
	t.Log("On Unix:")
	t.Log(filepath.Join("a", "b", "c"))
	t.Log(filepath.Join("a", "b/c"))
	t.Log(filepath.Join("a/b", "c"))
	t.Log(filepath.Join("a/b", "/c"))
	//Split
	paths := []string{
		"/home/arnie/amelia.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
	}
	t.Log("On Unix:")
	for _, p := range paths {
		dir, file := filepath.Split(p)
		t.Logf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
	}
	//SplitList
	t.Log("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
}

func TestWalk(t *testing.T) {
	//prepare test dir
	/*
		tree := "dir/to/walk/skip"
		tmpDir, err := ioutil.TempDir("", "")
		if err != nil {
			t.Fatal(err)
			err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
			if err != nil {
				os.RemoveAll(tmpDir)
				t.Fatal(err)
			}
		}
		defer os.RemoveAll(tmpDir)
		os.Chdir(tmpDir)*/

	subDirToSkip := "skip"
	err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}

}
