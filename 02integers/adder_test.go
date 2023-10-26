package integers

import (
	"fmt"
	"testing"
)

func TestAdded(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("Expected %d but got %d", expected, sum)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func BenchmarkAdded(b *testing.B){
	for i:=0;i<b.N;i++{
		Add(2,2)
	}
}
