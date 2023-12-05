package arrays

import (
	"reflect"
	"testing"
	"hello/generics"
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

func TestArrsumGenerics(t *testing.T) {
	t.Run("test for collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumGenerics(numbers)
		want := 15
		if got != want {
			t.Errorf("Got %d, want %d, given %v", got, want, numbers)
		}
	})

}


func TestSumAllTailsGenerics(t *testing.T) {
	t.Run("sum of given slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}
	})
	t.Run("sum of empty slice", func(t *testing.T) {
		got := SumAllTailsGenerics([]int{}, []int{0, 9})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}
	})

}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		generics.AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		generics.AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

// func TestBadBank(t *testing.T) {
// 	transactions := []Transaction{
// 		{
// 			From: "Chris",
// 			To:   "Riya",
// 			Sum:  100,
// 		},
// 		{
// 			From: "Adil",
// 			To:   "Chris",
// 			Sum:  25,
// 		},
// 	}

// 	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
// 	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
// 	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
// }