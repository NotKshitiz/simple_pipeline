// main_test.go
package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello, Jenkins Pipeline with Golang!"
	output := getMessage()
	if output != expected {
		t.Errorf("expected %s but got %s", expected, output)
	}
}