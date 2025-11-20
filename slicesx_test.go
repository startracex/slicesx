package slicesx

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	raw := []int{1, 2, 3}
	out := Map(raw, func(v int, _ int) int { return v * 2 })
	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Map failed: expected %v, got %v", expected, out)
	}
}

func TestFilter(t *testing.T) {
	raw := []int{1, 2, 3, 4}
	out := Filter(raw, func(v int, _ int) bool { return v%2 == 0 })
	expected := []int{2, 4}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Filter failed: expected %v, got %v", expected, out)
	}
}

func TestForEach(t *testing.T) {
	raw := []int{1, 2, 3}
	sum := 0
	ForEach(raw, func(v int, _ int) {
		sum += v
	})
	if sum != 6 {
		t.Errorf("ForEach failed: expected 6, got %d", sum)
	}
}

func TestSome(t *testing.T) {
	raw := []int{1, 2, 3}
	ok := Some(raw, func(v int, _ int) bool { return v == 2 })
	if !ok {
		t.Errorf("Some failed: expected true, got false")
	}
}

func TestEvery(t *testing.T) {
	raw := []int{2, 4, 6}
	ok := Every(raw, func(v int, _ int) bool { return v%2 == 0 })
	if !ok {
		t.Errorf("Every failed: expected true, got false")
	}
}

func TestReduce(t *testing.T) {
	raw := []int{1, 2, 3}
	out := Reduce(raw, func(acc, cur int, _ int) int { return acc + cur }, 0)
	if out != 6 {
		t.Errorf("Reduce failed: expected 6, got %d", out)
	}
}

func TestReduceRight(t *testing.T) {
	raw := []string{"a", "b", "c"}
	out := ReduceRight(raw, func(acc, cur string, _ int) string {
		return acc + cur
	}, "")
	expected := "cba"

	if out != expected {
		t.Errorf("ReduceRight failed: expected %s, got %s", expected, out)
	}
}

func TestUnshift(t *testing.T) {
	raw := []int{3, 4}
	out := Unshift(raw, 1, 2)
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Unshift failed: expected %v, got %v", expected, out)
	}
}

func TestPush(t *testing.T) {
	raw := []int{1, 2}
	out := Push(raw, 3, 4)
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Push failed: expected %v, got %v", expected, out)
	}
}

func TestShift(t *testing.T) {
	raw := []int{1, 2, 3}
	out := Shift(raw)
	expected := []int{2, 3}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Shift failed: expected %v, got %v", expected, out)
	}

	// empty slice
	out2 := Shift([]int{})
	if len(out2) != 0 {
		t.Errorf("Shift failed on empty slice: expected empty, got %v", out2)
	}
}

func TestPop(t *testing.T) {
	raw := []int{1, 2, 3}
	out := Pop(raw)
	expected := []int{1, 2}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Pop failed: expected %v, got %v", expected, out)
	}

	// empty slice
	out2 := Pop([]int{})
	if len(out2) != 0 {
		t.Errorf("Pop failed on empty slice: expected empty, got %v", out2)
	}
}

func TestReverse(t *testing.T) {
	raw := []int{1, 2, 3}
	out := Reverse(raw)
	expected := []int{3, 2, 1}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Reverse failed: expected %v, got %v", expected, out)
	}
}

func TestFlat(t *testing.T) {
	raw := [][]int{
		{1, 2},
		{3},
		{4, 5},
	}
	out := Flat(raw)
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Flat failed: expected %v, got %v", expected, out)
	}
}

func TestFlatMap(t *testing.T) {
	raw := []int{1, 2, 3}
	out := FlatMap(raw, func(v int, _ int) []string {
		return []string{"x", fmt.Sprint(v)}
	})
	expected := []string{"x", "1", "x", "2", "x", "3"}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("FlatMap failed: expected %v, got %v", expected, out)
	}
}

func TestSplice(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	out := Splice(raw, 1, 2, 9, 9)
	expected := []int{1, 9, 9, 4, 5}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Splice failed: expected %v, got %v", expected, out)
	}
}

func TestIsSlice(t *testing.T) {
	if !IsSlice([]int{1}) {
		t.Errorf("IsSlice failed: expected true for slice")
	}
	if IsSlice(10) {
		t.Errorf("IsSlice failed: expected false for non-slice")
	}
	if IsSlice(nil) {
		t.Errorf("IsSlice failed: expected false for nil")
	}
}
