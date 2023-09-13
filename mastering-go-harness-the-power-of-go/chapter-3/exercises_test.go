package exercises_test

import (
	exercises "mastering-go-exercises/chapter-3"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMapToSlice(t *testing.T) {
	t.Parallel()

	wantKey := []int{1, 2, 3}
	wantValue := []string{"a", "b", "c"}
	aMap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	gotKey, gotValue := exercises.MapToSlice(aMap)

	sort.Ints(gotKey)
	sort.Strings(gotValue)

	if !cmp.Equal(wantKey, gotKey) || !cmp.Equal(wantValue, gotValue) {
		t.Fatalf("wantKey %d, gotKey %d, wantValue %s, gotValue %s", wantKey, gotKey, wantValue, gotValue)
	}
}

func TestArgToStruct(t *testing.T) {
	t.Parallel()

	want := []exercises.ArgStructure{
		{Index: 1, Parameter: "1"},
		{Index: 2, Parameter: "2"},
		{Index: 3, Parameter: "3"},
	}

	got := exercises.ArgToStruct("1 2 3")

	if !cmp.Equal(want, *got) {
		t.Fatalf("Structures do not match")
	}
}
