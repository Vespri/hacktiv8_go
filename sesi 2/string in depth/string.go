package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Looping Over String (byte-by-byte)
	city := "Palembang"

	for i := 0; i < len(city); i++ {
		fmt.Printf("Index : %d, byte : %d\n", i, city[i])
	}

	var byteCity []byte = []byte{80, 97, 108, 101, 109, 98, 97, 110, 103}

	fmt.Println(string(byteCity))

	// Looping Over String (rune-by-rune)
	str1 := "coding"

	str2 := "eâting"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	str3 := "drinking"

	str4 := "sleeping"

	fmt.Printf("str3 character length => %d\n", utf8.RuneCountInString(str3))
	fmt.Printf("str4 character length => %d\n", utf8.RuneCountInString(str4))

	str5 := "hâi"

	for i, s := range str5 {
		fmt.Printf("Index => %d, string => %s\n", i, string(s))
	}
}
