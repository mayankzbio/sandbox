package concurrent_test

import (
	"testing"

	"github.com/mayankzbio/sandbox/sample/concurrent"
)

func TestAdder_Sum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"both positives", 2, 3, 5},
		{"one negative", -1, 5, 4},
		{"both negatives", -2, -3, -5},
		{"zero and positive", 0, 7, 7},
		{"zero and negative", 0, -4, -4},
	}

	adder := concurrent.Adder{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := adder.Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestArc_Sum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"empty first", "", "world", "world"},
		{"empty second", "hello", "", "hello"},
		{"both empty", "", "", ""},
		{"simple words", "hello", "world", "helloworld"},
		{"with space", "hello ", "world", "hello world"},
	}

	arc := concurrent.Arc{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := arc.Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
