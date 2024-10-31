package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("sgfd.txt")
	if err != nil {
		log.Fatal(err)
	}

	sliceData := file.ReadString()

	emptySlice := sliceData.UniqueValue()

	sort.Strings(emptySlice)

	emptySlice.WriteString()
	defer file.Close()
}

func UniqueValue(sliceData, emptySlice []string) {
	for i := 0; i < len(sliceData); i++ {
		k := 0
		for g := 0; g < len(sliceData); g++ {
			if sliceData[i] == sliceData[g] {
				k++
			}
		}

		if k == 1 {
			emptySlice = append(emptySlice, sliceData[i])
		}
	}
	return
}

func ReadString(sliceData []string) {
	file, err := os.Open("shgfd.txt")
	if err != nil {
		err.Error()
	}
	scaner, err := os.ReadFile("shgfd.txt")
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sliceData = append(sliceData, string(scaner))
	}
	if err != nil {
		return
	}
	return
}

func WriteString(emptySlice []string) {
	file, err := os.Create("hw1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < len(emptySlice); i++ {
		emptySlice[i] = strings.ToUpper(emptySlice[i])
	}
	for i := 0; i < len(emptySlice); i++ {
		text := emptySlice[i]
		n := len(emptySlice[i])
		file.WriteString(text)
		file.WriteString("-")
		file.WriteString(string(n) + "\n")
	}

	fmt.Println("Done.")
	return
}
