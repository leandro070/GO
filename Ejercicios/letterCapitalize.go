/*
Have the function LetterCapitalize(str) take the str parameter being passed and capitalize the first letter of each word.
Words will be separated by only one space. 
*/


package main
import ("fmt"
  "strings"
  "bytes")

func LetterCapitalize(str string) string {
  var buffer bytes.Buffer
  palabras:=strings.Fields(str)
  for _,palabra:= range palabras{
    inicio:=strings.ToUpper(string(palabra[0]))
    resto:=string(palabra[1:])
    palabra=inicio+resto
    buffer.WriteString(palabra+" ")
  }

  return buffer.String()

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(LetterCapitalize(readline()))

}
