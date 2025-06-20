package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	codes := map[string]string{
		"black":  "0",
		"brown":  "1",
		"red":    "2",
		"orange": "3",
		"yellow": "4",
		"green":  "5",
		"blue":   "6",
		"violet": "7",
		"grey":   "8",
		"white":  "9",
	}
	total := ""
	for i := 0; i < 2; i++ {
		total += codes[colors[i]]
	}
	result, _ := strconv.Atoi(total)
	return result
}
