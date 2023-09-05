package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a    float64
	b    float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 4.5, b: 4.5, want: 9},
		{a: 6, b: -6, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubstract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 1, want: 1},
		{a: 4, b: 4, want: 0},
		{a: 1, b: 2, want: -1},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 1, b: 1, want: 1}, {
			a: 2, b: 0.5, want: 1,
		},
		{a: 4, b: -9, want: -36},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if tc.want != got {
			t.Errorf("Divide(%f,%f):want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 3, b: 3, want: 9},
	}

	for _, tc := range testCases {
		got := calculator.Sqrt(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Sqrt(%f, %f), want: %f, got: %f", tc.a, tc.b, tc.want, got)
		}
	}
}
