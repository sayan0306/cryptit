package decrypt

func Nimbus(str string) string {
	var decryptedStr string
	for _, char := range str {
		ascii := int(char)
		ch := string(rune(ascii - 3))
		decryptedStr += ch
	}
	return decryptedStr
}
