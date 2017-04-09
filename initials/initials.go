package main

import (
	"../utils"
	"fmt"
	"strings"
)

const (
	ASCII_CODE_A rune = 65
	ASCII_CODE_Z rune = 90
)

func main() {
	showInitials(makeInitials(getName()))
}
//===================================================================================================
func makeInitials(name string) string{
	name = formatString(name)
	bytesName := utils.GetRunesFromString(name)
	initials := ""
	for i := 0; i < len(bytesName); i++ {
		if bytesName[i] >= ASCII_CODE_A && bytesName[i] <= ASCII_CODE_Z{
			initials += string(bytesName[i])
		}
	}

	return initials
}
//===================================================================================================
func formatString(name string) string{

	return strings.Title(strings.ToLower(name))
}
//===================================================================================================
func showInitials(initials string){
	fmt.Println(initials)
}
//===================================================================================================
func getName() string {
	fmt.Print("Please, enter your name: ")
	return utils.GetString()
}