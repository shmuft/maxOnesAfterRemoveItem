package main

import "testing"

type addTest struct {
	array    []byte
	expected uint
}

var taskTests = []addTest{
	{[]byte{0, 0}, 0},
	{[]byte{0, 1}, 1},
	{[]byte{1, 0}, 1},
	{[]byte{1, 1}, 1},
	{[]byte{1, 1, 0, 1, 1}, 4},
	{[]byte{1, 1, 0, 1, 1, 0, 1, 1, 1}, 5},
	{[]byte{1, 1, 0, 1, 1, 0, 1, 1, 1, 0}, 5},
}

var additionalTests = []addTest{
	{[]byte{1, 1, 1, 1}, 3},
	{[]byte{0, 1, 1, 0, 1, 1, 0}, 4},
	{[]byte{0, 1, 1, 1}, 3},
}

func TestFromTask(t *testing.T) {
	for _, test := range taskTests {
		if output := maxOnesAfterRemoveItem(test.array); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestAdditionalTests(t *testing.T) {
	for _, test := range additionalTests {
		if output := maxOnesAfterRemoveItem(test.array); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
