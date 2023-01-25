package main

import (
	"fmt"
	"github.com/bjohnrl/ber-test/decrypt"
	"github.com/bjohnrl/ber-test/encrypt"
)

func main() {
	givenString := "Hello All"
	encrptedString := encrypt.Nimbus(givenString)
	fmt.Printf("Encrypted String is %q\n", encrptedString)
	fmt.Printf("Decrypted string is %q\n", decrypt.Nimbus(encrptedString))

}
