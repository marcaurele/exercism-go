package cipher

import "strings"

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type caesar struct {
}
type shift struct {
	Distance int
}
type vigenere struct {
	Key string
}

func NewCaesar() Cipher {
	return caesar{}
}

func (c caesar) Encode(input string) string {
	var result []rune
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			result = append(result, 'a'+((letter-'a'+3)%26))
		}
	}
	return string(result)
}

func (c caesar) Decode(input string) string {
	var result []rune
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			result = append(result, 'a'+((letter-'a'+(26-3))%26))
		}
	}
	return string(result)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return shift{
		Distance: distance,
	}
}

func (c shift) Encode(input string) string {
	var result []rune
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			result = append(result, 'a'+((letter-'a'+26+rune(c.Distance))%26))
		}
	}
	return string(result)
}

func (c shift) Decode(input string) string {
	var result []rune
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			result = append(result, 'a'+((letter-'a'+(26-rune(c.Distance)))%26))
		}
	}
	return string(result)
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	diffThanA := false
	for _, l := range key {
		if l < 'a' || l > 'z' {
			return nil
		}
		if !diffThanA && l != 'a' {
			diffThanA = true
		}
	}
	if !diffThanA {
		return nil
	}
	return vigenere{
		Key: key,
	}
}

func (v vigenere) Encode(input string) string {
	var result []rune
	keyLength := len(v.Key)
	key := []rune(v.Key)
	keyMove := 0
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			result = append(result, 'a'+((letter-'a'+(key[keyMove%keyLength]-'a')%26)%26))
			keyMove += 1
		}
	}
	return string(result)
}

func (v vigenere) Decode(input string) string {
	var result []rune
	keyLength := len(v.Key)
	key := []rune(v.Key)
	keyMove := 0
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			result = append(result, 'a'+((26+letter-'a'-(key[keyMove%keyLength]-'a')%26)%26))
			keyMove += 1
		}
	}
	return string(result)
}
