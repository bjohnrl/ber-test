package decrypt

func Nimbus(str string) string {
	var decryptStr string
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decryptStr += character
	}
	return decryptStr
}
