package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("Airell", 23)

	var names = []string{"Kresna", "Vespri"}
	var printMsg = greet2nd("Hei", names)

	fmt.Println(printMsg)

	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	var area2nd, circumference2nd = calculate2nd(diameter)

	fmt.Println("Area :", area)
	fmt.Println("Circumference :", circumference)
	fmt.Println("Area Kedua :", area2nd)
	fmt.Println("Circumference Kedua :", circumference2nd)

	studentLists := print("Kresna", "Vespri", "Wijaya", "KVW", "VW")

	fmt.Printf("%v", studentLists)
	fmt.Println()

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)

	fmt.Println("Result :", result)

	profile("Kresna", "sate", "bakso", "model", "nasgor")
}

// Normal Funtion
func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}

// Return
func greet2nd(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// Returning multiple values
func calculate(d float64) (float64, float64) {
	// Menghitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	// Menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// Predefined return value
func calculate2nd(d float64) (area float64, circumference float64) {
	// Menghitung luas
	area = math.Pi * math.Pow(d/2, 2)

	// Menghitung keliling
	circumference = math.Pi * d

	return
}

// Variadic function #1
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// // Variadic function #2
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

// Variadic function #3
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
