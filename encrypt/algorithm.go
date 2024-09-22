package encrypt

func Nimbus(str string) string {
	var encryptedStr string
	for _, char := range str {
		ascii := int(char)
		ch := string(rune(ascii + 3))
		encryptedStr += ch
	}
	return encryptedStr
}
