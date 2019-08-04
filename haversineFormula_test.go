package main

import (
	"testing"
)

func TestHaversineFormula(t *testing.T) {

	//test for several coordinates
	result := harvesineFormula(51.5007, 0.1246, 40.6892, 74.0445)

	if result != 5574.840456848555 {
		t.Errorf("Expected result 5574.840456848555, but got %v", result)
	}

	result1 := harvesineFormula(0, 0, 0, 0)

	if result1 != 0 {
		t.Errorf("Expected result  0, but got %v", result)
	}

	result2 := harvesineFormula(-1, -1, -1, -1)

	if result2 != 0 {
		t.Errorf("Expected result  0, but got %v", result)
	}

}
