package main

import (
	"../utils"

	"fmt"
	"unicode"
	"strings"
)

const (
	ALPHABET_LENGTH int = 26
	ASCII_CODE_OF_A rune = 65
)
//===================================================================================================
func main() {
	showCipherText(cipherText(getPlainText(),getKey()))
}
//===================================================================================================
func getKey() string {
	fmt.Print("Please, enter a secret key: ")
	key := utils.GetString()
	if key != "" {
		return strings.ToUpper(key)
	}else{
		fmt.Println("Wrong key, you should enter string!")
	}

	return getKey()
}
//===================================================================================================
func getPlainText() string {
	fmt.Print("Please, enter a plain text: ")

	return utils.GetString()
}
//===================================================================================================
func cipherText(text string, key string) string{
	runesFromText := utils.GetRunesFromString(text)
	runesFromKey := utils.GetRunesFromString(key)

	var isLetter bool

	for i, k := 0, 0; i < len(runesFromText); i++ {
		runesFromText[i], isLetter = cipherRune(runesFromText[i], runesFromKey[k % len(runesFromKey)])
		if isLetter {
			k++
		}
	}

	return string(runesFromText)
}
//===================================================================================================
func cipherRune(r rune, key rune) (rune, bool) {
	key -= ASCII_CODE_OF_A
	if unicode.IsLetter(r) {
		r += rune(key)
		if !unicode.IsLetter(r) {
			r -= rune(ALPHABET_LENGTH)
		}
		return r, true
	}
	return r, false
}
//===================================================================================================
func showCipherText(text string){
	fmt.Println("ciphertext: ",text)
}