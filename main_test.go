package main

import (
	"io/ioutil"
	"testing"
)

func TestWriteConfiguration(t *testing.T) {
	writeConfiguration("/tmp/foobar.txt", "hello world")
	data, _ := ioutil.ReadFile("/tmp/foobar.txt")

	if string(data) != "hello world" {
		t.Errorf("File content should be 'hello world'. Got %s", data)
	}
}
