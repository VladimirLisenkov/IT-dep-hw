package calculator

import (
	"errors"
	"shapes"
)

func TotalArea(figures ...interface{}) (float64, error) {
	total := 0.0
	for _, figure := range figures {
		switch f := figure.(type) {
		case shapes.Circle:
			area, err := f.Area()
			if err != nil {
				return 0, err
			}
			total += area
		case shapes.Rectangle:
			area, err := f.Area()
			if err != nil {
				return 0, err
			}
			total += area
		default:
			return 0, errors.New("неизвестная фигура")
		}
	}
	return total, nil
}
