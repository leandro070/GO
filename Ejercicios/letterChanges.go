/*
 Have the function LetterChanges(str) take the str parameter being passed and modify it using the following algorithm.
 Replace every letter in the string with the letter following it in the alphabet (ie. c becomes d, z becomes a).
 Then capitalize every vowel in this new string (a, e, i, o, u) and finally return this modified string. 
 */

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func LetterChanges(str string) string {
	var buffer bytes.Buffer

	for i := 0; i < len(str); i++ {
		caracter := int(str[i])
		if (caracter > 64 && caracter < 91) || (caracter > 96 && caracter < 123) {
			caracter++
			if caracter == 97 || caracter == 101 || caracter == 105 || caracter == 111 || caracter == 117 {
				letra := strings.ToUpper(string(caracter))
				buffer.WriteString(letra)
			} else {
				buffer.WriteByte(byte(caracter))
			}
		} else {
			buffer.WriteByte(byte(caracter))
		}
	}

	cadena := buffer.String()

	return cadena

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(LetterChanges(readline()))

}
