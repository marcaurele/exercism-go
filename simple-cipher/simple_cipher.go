package cipher

import (
	"strings"
)

type shift struct {
	Distance int
}

type vigenere struct {
	Key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < 0 {
		distance += 26
	}
	return NewVigenere(string(rune('a' + distance)))
}

func NewVigenere(key string) Cipher {
	onlyA := true
	for _, l := range key {
		if l < 'a' || l > 'z' {
			return nil
		}
		if onlyA && l != 'a' {
			onlyA = false
		}
	}
	if onlyA {
		return nil
	}
	return vigenere{
		Key: key,
	}
}

func (v vigenere) Encode(input string) string {
	return v.shift(input, true)
}

func (v vigenere) Decode(input string) string {
	return v.shift(input, false)
}

func (v vigenere) shift(input string, encode bool) string {
	var result []rune
	keyLength := len(v.Key)
	key := []rune(v.Key)
	keyMove := 0
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			if encode {
				result = append(result, 'a'+((letter-'a'+(key[keyMove%keyLength]-'a')%26)%26))
			} else {
				result = append(result, 'a'+((26+letter-'a'-(key[keyMove%keyLength]-'a')%26)%26))
			}
			keyMove += 1
		}
	}
	return string(result)
}
