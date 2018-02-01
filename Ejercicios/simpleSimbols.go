/*
Have the function SimpleSymbols(str) take the str parameter being passed and determine
if it is an acceptable sequence by either returning the string true or false.
The str parameter will be composed of + and = symbols with several letters between
them (ie. ++d+===+c++==a) and for the string to be true each letter must be surrounded by a + symbol.
So the string to the left would be false. The string will not be empty and will have at least one letter. 
*/

package main
import "fmt"

func SimpleSymbols(str string) string {

 	for i := 0; i < len(str); i++ {
		letra := int(str[i])
		anterior := i - 1
		posterior := i + 1
		if (letra > 96 && letra < 123) || (letra > 64 && letra < 91) {
			if anterior >= 0 && posterior <= len(str) {
				letraAnt := int(str[anterior])
				letraPost := int(str[posterior])
				if letraAnt != 43 || letraPost != 43 {
					return "false"
				}
			} else {
				return "false"
			}
		}
	}
	return "true"

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(SimpleSymbols(readline()))

}
