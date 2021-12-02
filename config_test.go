package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	b := []byte{}

	fmt.Println("need to write new config tests")
	ts := Token{}
	s, err := ts.Parse(b)

	if err != nil {
		t.Errorf("Expected string, got: %v %v", s, err)
	}
}
