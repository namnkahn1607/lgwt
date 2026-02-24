package struct_method_interface

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Retangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		if got := shape.Area(); got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 40.0},
		{name: "Circle", shape: Circle{10.0}, want: 62.83185307179586},
	}

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		if got := shape.Perimeter(); got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.want)
		})
	}
}
