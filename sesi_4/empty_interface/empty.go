package main

func main() {
	// Empty Interface
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"

	randomValues = 20

	randomValues = true

	randomValues = []string{"Kresna", "Vespri"}

	// Type assertion
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	// Map & Slice with Empty Interface
	rs := []interface{}{1, "Kresna", true, 2, "Vespri", true}

	rm := map[string]interface{}{
		"Name":   "Kresna",
		"Status": true,
		"Age":    21,
	}

	_, _ = rs, rm
}
