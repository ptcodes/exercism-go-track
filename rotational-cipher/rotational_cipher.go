package rotationalcipher

func RotationalCipher(input string, shiftKey int) string {
	encrypted := []rune{}
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+rune(shiftKey))%26
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+rune(shiftKey))%26
		}
		encrypted = append(encrypted, char)
	}
	return string(encrypted)
}
