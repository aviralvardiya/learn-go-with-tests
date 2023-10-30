package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12, 8}
	got := Perimeter(rectangle)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f, want %.2f ", got, want)
	}

}

func TestArea(t *testing.T) {

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %.2f, want %.2f ", got, want)
	// 	}

	// }

	// t.Run("rectanle area", func(t *testing.T) {
	// 	rectangle := Rectangle{10, 11}
	// 	want := 110.0
	// 	checkArea(t, rectangle, want)
	// })
	// t.Run("circle area", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 20}, 200},
		{Circle{10}, 314.1592653589793},
		{Triangle{12,6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want
		if got != want {
			t.Errorf("for %#v=> got %g, want %g ",tt, got, want)
		}
	}

}
