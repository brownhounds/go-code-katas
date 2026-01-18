package productexceptself

import "testing"

func assertSliceEq(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %d, want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("at index %d: got %d, want %d; full got=%v want=%v", i, got[i], want[i], got, want)
		}
	}
}

func TestProductExceptSelf_Basic(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := ProductExceptSelf(nums)
	want := []int{24, 12, 8, 6}
	assertSliceEq(t, got, want)
}

func TestProductExceptSelf_WithZero(t *testing.T) {
	nums := []int{1, 2, 0, 4}
	got := ProductExceptSelf(nums)
	want := []int{0, 0, 8, 0}
	assertSliceEq(t, got, want)
}

func TestProductExceptSelf_Negatives(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3}
	got := ProductExceptSelf(nums)
	want := []int{0, 0, 9, 0, 0}
	assertSliceEq(t, got, want)
}

func TestProductExceptSelf_SingleElement(t *testing.T) {
	nums := []int{5}
	got := ProductExceptSelf(nums)
	want := []int{1}
	assertSliceEq(t, got, want)
}

func TestProductExceptSelf_EmptySlice(t *testing.T) {
	var nums []int
	got := ProductExceptSelf(nums)
	var want []int
	assertSliceEq(t, got, want)
}
