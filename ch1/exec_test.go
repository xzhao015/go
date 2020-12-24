package ch1

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestLookpath(t *testing.T) {
	//lookpath find executable named file in PATH
	path, err := exec.LookPath("go")
	if err != nil {
		t.Fatal("installing go is in your future")
	}
	t.Logf("go is available at %s\n", path)
}

func TestCommand(t *testing.T) {
	cmd := exec.Command("go", "version")
	//cmd.Stdin = strings.NewReader("version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(out.String())
}
