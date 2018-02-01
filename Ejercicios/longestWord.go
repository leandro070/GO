/*
Have the function LongestWord(sen) take the sen parameter being passed and return the largest word in the string.
If there are two or more words that are the same length, return the first word from the string with that length.
Ignore punctuation and assume sen will not be empty.
*/


package main
import ("fmt"
"strings")

func LongestWord(sen string) string {
cadenas := strings.Fields(sen) //separo la cadena en elementos de un array sacando los espacios
	cadenaLarga := ""

	for _, cadena := range cadenas {
		caracteresNoDeseados := "~`.,:;!@#$%^&*()_-+={[}]/\\'\""

		for i := 0; i < len(caracteresNoDeseados); i++ {
			//saco una runa de los caracteresNoDeseados
			runa := []rune(caracteresNoDeseados)[i]
			//compruebo si existe em la cadena el caracteres No Deseado anterior
			if strings.ContainsRune(cadena, runa) {
				//elimino caracter no deseado
				cadena = strings.Replace(cadena, string(runa), "", -1)
			}
		}
		//Compruebo el largo de las cadenas y guardo la mas larga
		if len(cadena) > len(cadenaLarga) {
			cadenaLarga = cadena
		}

	}
	return cadenaLarga


}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(LongestWord(readline()))

}
