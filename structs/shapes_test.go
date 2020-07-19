package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{Height: 5, Width: 4}
		got := rectangle.Perimeter()
		want := 18.0
		if want != got {
			t.Errorf("%#v got %g want %g", rectangle, got, want)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{Radius: 5}
		got := circle.Perimeter()
		want := 31.41592653589793
		if want != got {
			t.Errorf("%#v got %g want %g", circle, got, want)
		}
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
