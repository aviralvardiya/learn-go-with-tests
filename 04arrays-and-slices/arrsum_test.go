package arrays

import (
	"reflect"
	"testing"
)

func TestArrsum(t *testing.T) {
	t.Run("test for collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := ArrSum(numbers)
		want := 15
		if got != want {
			t.Errorf("Got %d, want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// want:="bob" // be careful when usig DeepEqual, it isn't type safe

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v , want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum of given slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}
	})
	t.Run("sum of empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}
	})

}
