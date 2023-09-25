package capture

/*
#include <stdio.h>
void printSomething() {
	printf("C\n");
}
*/
import "C"

import (
	"fmt"
	"strings"
	"testing"
)

func testCapture(t *testing.T) {
	out, err := Capture(func() {
		fmt.Println("Go")
		C.printSomething()
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "Go") {

	}

	//assert.NoError(t, err)
	//assert.Contains(t, string(out), "Go")
	//assert.NotContains(t, string(out), "C")
}

func testCaptureWithCGo(t *testing.T) {
	out, err := CaptureWithCGo(func() {
		fmt.Println("Go")
		C.printSomething()
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "Go") {

	}
	//assert.NoError(t, err)
	//assert.Contains(t, string(out), "Go")
	//assert.Contains(t, string(out), "C")
}
