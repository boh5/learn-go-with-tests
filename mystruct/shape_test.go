package mystruct

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("calculate perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f hasArea %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %.2f hasArea %.2f", tt.shape, got, tt.hasArea)
		}
	}
}
