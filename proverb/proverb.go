package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var output []string
	if len(rhyme) > 0 {
		for i := 1; i < len(rhyme); i++ {
			output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
		}
		output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return output
}
