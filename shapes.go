package shapes

import (
	"errors"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New("некорректное значение радиуса")
	}
	return math.Pi * c.Radius * c.Radius, nil
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() (float64, error) {
	if r.Width < 0 || r.Height < 0 {
		return 0, errors.New("некорректные размеры прямоугольника")
	}
	return r.Width * r.Height, nil
}
