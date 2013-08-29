package main

import (
	"fmt"
	"os"
	"testing"
)

const sandboxDir = "sandbox"
const sandboxDirPerms = 0755

/*
use the same struct filename/content for testing file writing and reading.
*/

type testFile struct {
	filename string
	content  []string
}

var fileContentTests = []testFile{
	{"3lines", []string{"one", "two", "three"}},
	{"empty", []string{}},
}

/*
use a sandbox for testing file creation/reading
*/
func sandboxUp() (err error) {
	err = os.Mkdir(sandboxDir, sandboxDirPerms)
	return
}

func sandboxDown() (err error) {
	err = os.RemoveAll(sandboxDir)
	return
}

func sandboxSetup() (err error) {
	err = sandboxDown()
	err = sandboxUp()
	return
}

func TestWriteLinesToFile(t *testing.T) {
	sandboxSetup()
	for _, tt := range fileContentTests {
		fileName := sandboxDir + "/" + tt.filename
		success := WriteLinesToFile(tt.content, fileName)
		if success != nil {
			t.Errorf("WriteLinesToFile(%s): error(%v)", fileName, success)
		}
	}
}

func TestReadLines(t *testing.T) {
	sandboxDestroy := true
	for _, tt := range fileContentTests {
		fileName := sandboxDir + "/" + tt.filename
		actual, _ := ReadLines(fileName)
		a := fmt.Sprintf("%v", actual)
		r := fmt.Sprintf("%v", tt.content)
		if a != r { // laugh all you want, I don't want to write a slice operator
			t.Errorf("ReadLines(%s): expected %s, actual %s", fileName, r, a)
			sandboxDestroy = false
		}
	}
	if sandboxDestroy {
		defer sandboxDown()
	}
}
