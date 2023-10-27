package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	t.Run("just giving char, no of iterations not provided",func(t *testing.T) {
		got:=Repeat("a",0)
	want:="aaaaa"

	if(got!=want){
		t.Errorf("wanted %q but got %q",want,got)
	}
	})
	t.Run("providing a char and number of iterations",func(t *testing.T) {
		got:=Repeat("a",7)
	want:="aaaaaaa"

	if(got!=want){
		t.Errorf("wanted %q but got %q",want,got)
	}
	})
	
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a",0)
	}
}		


func ExampleRepeat() {
	str := Repeat("y", 5)
	fmt.Println(str)
	// Output: yyyyy
}