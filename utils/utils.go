package utils

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)
//=============================================
// Return rune array from sting
//=============================================
func GetRunesFromString(text string)[]rune {

	return []rune(text)
}
//=============================================
// Return text from user, print error if exist
//=============================================
func GetString() string{
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println(err)
	}

	return strings.TrimSpace(text)
}