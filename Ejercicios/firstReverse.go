/*
Have the function FirstReverse(str) take the str parameter being passed and return the string in reversed order.
For example: if the input string is "Hello World and Coders" then your program should return the string sredoC dna dlroW olleH.
*/

package main

import (
	"bytes"
	"fmt"
)

func FirstReverse(str string) string {
	// Aquí, bytes.Buffer es un buffer de bytes de tamaño variable.
	var buffer bytes.Buffer
	//Escribo el buffer
	for i := len(str) - 1; i >= 0; i-- {
		buffer.WriteByte(str[i])
	}
	return buffer.String()
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(FirstReverse(readline()))

}
