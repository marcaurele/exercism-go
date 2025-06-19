package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	rot := []rune(plain)

	for i, r := range plain {
		if r >= 'a' && r <= 'z' {
			rot[i] = 'a' + (r+rune(shiftKey)-'a')%26
		} else if r >= 'A' && r <= 'Z' {
			rot[i] = 'A' + (r+rune(shiftKey)-'A')%26
		} else {
			rot[i] = r
		}
	}
	return string(rot)
}
