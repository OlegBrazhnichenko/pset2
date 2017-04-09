package main

import (
	"../utils"
	"strconv"
	"fmt"
	"unicode"
)

const (
	ALPHABET_LENGTH int = 26
)

func main() {
	showCipherText(cipherText(getPlainText(),getKey()))
}
//===================================================================================================
func getKey() int{
	fmt.Print("Please, enter a secret key: ")
	key, err := strconv.Atoi(utils.GetString())
	if err == nil {
		return key % ALPHABET_LENGTH
	}else{
		fmt.Println("Wrong key, you should write a number!")
	}

	return getKey() % ALPHABET_LENGTH
}
//===================================================================================================
func getPlainText() string {
	fmt.Print("Please, enter a plain text: ")

	return utils.GetString()
}
//===================================================================================================
func cipherText(text string, key int) string{
	runesFromText := utils.GetRunesFromString(text)
	var cipheredText string
	for i := 0; i < len(runesFromText); i++ {
		cipheredText += string(cipherRune(runesFromText[i], key))
	}

	return cipheredText
}
//===================================================================================================
func cipherRune(r rune, key int) rune {
	if unicode.IsLetter(r){
		r += rune(key)
		if !unicode.IsLetter(r){
			r -= rune(ALPHABET_LENGTH)
		}
	}

	return r
}
//===================================================================================================
func showCipherText(text string){
	fmt.Println("ciphertext: ",text)
}