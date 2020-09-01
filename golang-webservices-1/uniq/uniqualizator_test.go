package main

import (
	"strings"
	"bytes"
	"testing"
	"bufio"
)

var testText = `1
2
3
3
3
4
5
6
`

var testOkResult = `1
2
3
4
5
6
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testText))
	out := new(bytes.Buffer)
	err := unique(in, out)
	if err != nil {
		t.Errorf("test failed: error occured")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test failed: results do not match \n %v %v", result, testOkResult)
	}
}

var testUnsorted = `1
2
1
`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testUnsorted))
	out := new(bytes.Buffer)
	err := unique(in, out)
	if err == nil {
		t.Errorf("test failed: error was not thrown")
	}
}
