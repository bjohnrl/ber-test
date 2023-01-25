package encrypt

func Nimbus(str string) string {
	var encryptedStr string

	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
