package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits = []string{"apple", "banana", "manggo"}

	_ = fruits

	// make function
	var fruit = make([]string, 3)

	fmt.Printf("%#v", fruit)
	fmt.Println()

	// append function
	var object = make([]string, 3)

	_ = object

	object[0] = "computer"
	object[1] = "keyboard"
	object[2] = "hard disk"

	fmt.Printf("%#v", object)
	fmt.Println()

	// append function with ellipsis
	var objectFirst = []string{"spoon", "fork", "plate"}

	var objectSecond = []string{"pencil", "pen", "paper"}

	objectFirst = append(objectFirst, objectSecond...)

	fmt.Printf("%#v", objectFirst)
	fmt.Println()

	// copy function
	var fruitFirst = []string{"grape", "strawberry", "apple"}

	var fruitSecond = []string{"pineapple", "melon", "watermelon"}

	mix := copy(fruitFirst, fruitSecond)

	fmt.Println("fruitFirst :", fruitFirst)
	fmt.Println("fruitSecond :", fruitSecond)
	fmt.Println("Copied values :", mix)

	// Slicing
	var animalFirst = []string{"lion", "cat", "tiger", "dog", "wolf"}

	var animalSecond = animalFirst[1:4]
	fmt.Printf("%#v\n", animalSecond)

	var animalThird = animalFirst[0:]
	fmt.Printf("%#v\n", animalThird)

	var animalFourth = animalFirst[:3]
	fmt.Printf("%#v\n", animalFourth)

	var animalFifth = animalFirst[:]
	fmt.Printf("%#v\n", animalFifth)

	// Combining slicing and append
	var name = []string{"Kresna", "Vespri", "Wijaya"}

	name = append(name[:3], "KVW")

	fmt.Printf("%#v\n", name)

	// Backing array
	var toolFirst = []string{"VSCode", "Sublime", "Notepad++", "Firefox", "Chrome"}

	var toolSecond = toolFirst[2:4]

	toolSecond[0] = "postman"

	fmt.Println("First tools :", toolFirst)
	fmt.Println("Second tools :", toolSecond)

	// Cap function
	var drinkFirst = []string{"Ice Coffee", "Hot Coffee", "Ice Tea", "Hot Tea"}

	fmt.Println("drinkFirst cap :", cap(drinkFirst))
	fmt.Println("drinkFirst len :", len(drinkFirst))

	fmt.Println(strings.Repeat("#", 20))

	var drinkSecond = drinkFirst[0:3]

	fmt.Println("drinkSecond cap :", cap(drinkSecond))
	fmt.Println("drinkSecond len :", len(drinkSecond))

	fmt.Println(strings.Repeat("#", 20))

	var drinkThird = drinkFirst[1:]

	fmt.Println("drinkThird cap :", cap(drinkThird))
	fmt.Println("drinkThird len :", len(drinkThird))

	// Creating a new backing array
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars :", cars)
	fmt.Println("newCars :", newCars)
}
