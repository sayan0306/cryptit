package main

import (
	"fmt"
	"github.com/sayan0306/criptit/decrypt"
	"github.com/sayan0306/criptit/encrypt"
)

func main() {
	str := encrypt.Nimbus("Sayan")
	fmt.Println(str)
	str2 := decrypt.Nimbus(str)
	fmt.Println(str2)
}
