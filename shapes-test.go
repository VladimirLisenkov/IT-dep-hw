package shapes_test

import (
	"shapes"
	"testing"
)

func TestCircleArea(t *testing.T) {
	c := shapes.Circle{Radius: 5}
	area, err := c.Area()
	if err != nil {
		t.Error("произошла ошибка:", err)
	}

	if area != 78.53981633974483 {
		t.Errorf("получено %v", area)
	}

	c = shapes.Circle{Radius: -1}
	_, err = c.Area()
	if err == nil {
		t.Error("отрицательнsq радиус")
	}
}

func TestRectangleArea(t *testing.T) {
	r := shapes.Rectangle{Width: 5, Height: 10}
	area, err := r.Area()
	if err != nil {
		t.Error("произошла ошибка:", err)
	}

	if area != 50 {
		t.Errorf("получено %v", area)
	}

	r = shapes.Rectangle{Width: -1, Height: 10}
	_, err = r.Area()
	if err == nil {
		t.Error("ожидалась ошибка для отрицательных размеров")
	}
}
