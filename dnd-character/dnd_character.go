package dndcharacter

import (
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return (score - 10 - score%2) / 2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	total := rand.Intn(6) + 1
	lowest := total
	for i := 0; i < 3; i++ {
		score := rand.Intn(6) + 1
		total += score
		if score < lowest {
			lowest = score
		}
	}
	return total - lowest
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
