package main

import (
	"bufio"
	"calculator"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("figures.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	var figures []interface{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "Circle" {
			circle := Circle{Radius: radius}
		} else if line[0] == "Rectangle" {
			rectangle := Rectangle{Width: width, Height: height}
		}
	}

	total, err := calculator.TotalArea(figures...)
	if err != nil {
		fmt.Println("Ошибка при расчете площади:", err)
		return
	}
	fmt.Println("Общая площадь:", total)
}
