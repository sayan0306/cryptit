package main

import (
	"fmt"
	"github.com/sayan0306/cryptit/decrypt"
	"github.com/sayan0306/cryptit/encrypt"
)

func main() {
	str := encrypt.Nimbus("Sayan")
	fmt.Println(str)
	str2 := decrypt.Nimbus(str)
	fmt.Println(str2)
}
