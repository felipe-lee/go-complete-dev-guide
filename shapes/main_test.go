package main

import (
	"fmt"
	"testing"
)

func TestTriangleArea(t *testing.T) {
	t.Parallel() // marks TestTriangleArea as capable of running in parallel with other tests

	areaTests := []struct {
		base         float64
		height       float64
		expectedArea float64
	}{
		{2, 3, 3},
		{1, 5, 2.5},
		{5, 5, 12.5},
	}

	for _, at := range areaTests {
		at := at // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(fmt.Sprintf("base:%f|height:%f", at.base, at.height), func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other

			ts := triangle{base: at.base, height: at.height}

			area := ts.getArea()

			if area != at.expectedArea {
				t.Errorf("Got %f, expected %f", area, at.expectedArea)
			}
		})
	}
}

func TestSquareArea(t *testing.T) {
	t.Parallel() // marks TestSquareArea as capable of running in parallel with other tests

	areaTests := []struct {
		sideLength   float64
		expectedArea float64
	}{
		{3, 9},
		{5, 25},
		{10, 100},
	}

	for _, at := range areaTests {
		at := at // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(fmt.Sprintf("sideLength:%f", at.sideLength), func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other

			ss := square{sideLength: at.sideLength}

			area := ss.getArea()

			if area != at.expectedArea {
				t.Errorf("Got %f, expected %f", area, at.expectedArea)
			}
		})
	}
}
