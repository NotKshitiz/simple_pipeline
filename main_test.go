package main
import "testing"

func TestMain(t *testing.T){
	expected := "Hello, Jenkins Pipeline with Golang!"
	output:= getMessage()
	if output!=expected{
		t.Errorf("expected %s but got %s",expected,output)
	}
}

func getMessage() string {
	return "Hello, Jenkins Pipeline wit Golang!"
}